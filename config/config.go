package config

import (
	"context"
	"fmt"
	"os"
	"runtime"
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
	API         API
	Jwt         Jwt
	Grpc        Grpc
	Database    Database
	LoginServer LoginServer
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
	Enabled    bool
}

// Jwt key settings
type Jwt struct {
	PrivateKeyPath string
	PublicKeyPath  string
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
