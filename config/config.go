package config

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqemuconfig"
)

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

// Config represents a configuration parse
type Config struct {
	mutex       sync.Mutex
	API         API
	Jwt         Jwt
	Grpc        Grpc
	Database    Database
	LoginServer LoginServer
	Permissions map[string]Permission
}

// Database related settings
type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Db       string
}

// API listening settings
type API struct {
	Host string
}

// Grpc listening settings
type Grpc struct {
	Host string
}

// LoginServer listening settings
type LoginServer struct {
	WebAPIHost string
	APIToken   string
	IsEnabled  bool `toml:"enabled"`
}

// Jwt key settings
type Jwt struct {
	PrivateKeyPath string
	PublicKeyPath  string
}

// Permission for various endpoints
type Permission struct {
	Status    int64
	Endpoints map[string]map[string]PermissionEntry
}

// PermissionEntry is for each endpoint
type PermissionEntry struct {
	IsSelfOnly         bool `toml:"selfOnly"`
	IsLoginNotRequired bool `toml:"loginNotRequired"`
	Fields             []string
}

// NewConfig creates a new configuration
func NewConfig(ctx context.Context) (*Config, error) {
	var f *os.File
	cfg := Config{}
	path := "eqcp.conf"

	emucfg, err := eqemuconfig.GetConfig()
	if err != nil {
		log.Debug().Err(err).Msg("eqemu")
		emucfg = new(eqemuconfig.Config)
	}

	isNewConfig := false
	fi, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, errors.Wrap(err, "config info")
		}
		f, err = os.Create(path)
		if err != nil {
			return nil, errors.Wrap(err, "create eqcp.conf")
		}
		fi, err = os.Stat(path)
		if err != nil {
			return nil, errors.Wrap(err, "new config info")
		}
		isNewConfig = true
	}
	if !isNewConfig {
		f, err = os.Open(path)
		if err != nil {
			return nil, errors.Wrap(err, "open config")
		}
	}

	defer f.Close()
	if fi.IsDir() {
		return nil, fmt.Errorf("eqcp.conf is a directory, should be a file")
	}

	if isNewConfig {
		_, err = f.WriteString(defaultConfig)
		if err != nil {
			return nil, errors.Wrap(err, "write new eqcp.conf")
		}
		fmt.Println("a new eqcp.conf file was created. Please edit this file, then run again.")
		if runtime.GOOS == "windows" {
			option := ""
			fmt.Println("press a key then enter to exit.")
			fmt.Scan(&option)
		}
		os.Exit(0)
	}

	_, err = toml.DecodeReader(f, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "decode eqcp.conf")
	}

	if cfg.Jwt.PrivateKeyPath == "" {
		return nil, fmt.Errorf("jwt.private_key_path not set in eqcp.conf")
	}
	if cfg.Jwt.PublicKeyPath == "" {
		return nil, fmt.Errorf("jwt.public_key_path not set in eqcp.conf")
	}

	if cfg.Database.Host == "" {
		cfg.Database.Host = emucfg.Database.Host
	}
	if cfg.Database.Db == "" {
		cfg.Database.Db = emucfg.Database.Db
	}
	if cfg.Database.Password == "" {
		cfg.Database.Password = emucfg.Database.Password
	}
	if cfg.Database.Port == "" {
		cfg.Database.Port = emucfg.Database.Port
	}
	if cfg.Database.Username == "" {
		cfg.Database.Username = emucfg.Database.Username
	}
	return &cfg, nil
}

// Permission returns permissions of a provided field
func (c *Config) Permission(endpoint string, scope string, status int64) (label string, fields []string, isSelfOnly bool, isLoginNotRequired bool, isAllFieldsOK bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	curStatus := int64(-1)
	for acct, entry := range c.Permissions {

		if entry.Status > status {
			continue
		}
		if entry.Status < curStatus {
			continue
		}
		endpoint, ok := entry.Endpoints[strings.Title(endpoint)]
		if !ok {
			continue
		}
		perm, ok := endpoint[strings.Title(scope)]
		if !ok {
			continue
		}
		fields = perm.Fields
		if len(fields) == 1 && fields[0] == "*" {
			isAllFieldsOK = true
		}
		isSelfOnly = perm.IsSelfOnly
		isLoginNotRequired = perm.IsLoginNotRequired
		label = acct
		return
	}

	return
}
