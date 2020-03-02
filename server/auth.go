package server

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

// AuthData represents authentication token claim data
type AuthData struct {
	AccountID int64
}

// AuthPermission is used internally to map access rights
type AuthPermission struct {
	dataClaim *DataClaim
	accountID int64
	status    int64
	commands  []string
}

func (ap *AuthPermission) hasCommand(command string) bool {
	command = strings.ToLower(command)
	for _, cmd := range ap.commands {
		if command == cmd {
			return true
		}
	}
	return false
}

// AuthFromContext returns a claim if auth is passed
func (s *Server) AuthFromContext(ctx context.Context) (*AuthPermission, error) {
	var token string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md["authorization"]) > 0 {
		token = md["authorization"][0]
	}
	if token == "" {
		return nil, fmt.Errorf("requires token")
	}

	dc, err := s.TokenRead(token)
	if err != nil {
		return nil, errors.Wrap(err, "read")
	}

	var rows *sqlx.Rows
	var ap *AuthPermission
	val, ok := s.loginPermissions.Load(dc.Data["AccountID"])
	if ok {
		ap, ok = val.(*AuthPermission)
		if !ok {
			return nil, fmt.Errorf("permission invalid")
		}
		return ap, nil
	}

	ap = new(AuthPermission)
	ap.accountID = dc.Data["AccountID"].(int64)
	if ap.accountID < 1 {
		return nil, fmt.Errorf("token invalid (accountID not set)")
	}

	if err = s.db.GetContext(ctx, ap.status, "SELECT status FROM account WHERE id = ?", ap.accountID); err != nil {
		return nil, fmt.Errorf("failed to get account status")
	}

	rows, err = s.db.QueryxContext(ctx, "SELECT command, access FROM comand_settings WHERE access <= ?", ap.status)
	if err != nil {
		return nil, fmt.Errorf("failed to get command_settings")
	}
	for rows.Next() {
		command := ""
		if err = rows.Scan(command); err != nil {
			return nil, fmt.Errorf("scan command_sttings")
		}
		ap.commands = append(ap.commands, strings.ToLower(command))
	}

	s.loginPermissions.Store(ap.accountID, ap)
	return ap, nil
}

// AuthCreate generates a new token with a 24 hour timeout, returning the token as a string
func (s *Server) AuthCreate(ad *AuthData) (token string, err error) {
	dc := s.marshalDataClaim(ad)
	token, err = s.TokenCreate(dc)
	if err != nil {
		return
	}
	return
}

// AuthRead takes a token string and turns it into Auth
func (s *Server) AuthRead(token string) (ad *AuthData, err error) {
	dc, err := s.TokenRead(token)
	if err != nil {
		return
	}
	ad = s.unmarshalDataClaim(dc)
	return
}

func (s *Server) unmarshalDataClaim(dc *DataClaim) (ad *AuthData) {
	ad = &AuthData{}
	if dc == nil {
		return
	}

	v, ok := dc.Data["accountID"]
	if !ok {
		return
	}
	ad.AccountID = int64(v.(float64))
	return
}

func (s *Server) marshalDataClaim(ad *AuthData) (dc *DataClaim) {
	dc = &DataClaim{
		Data: make(map[string]interface{}),
	}
	if ad == nil {
		return
	}
	dc.Data["accountID"] = ad.AccountID
	return
}

// DataClaim represents token data
type DataClaim struct {
	Data map[string]interface{}
	jwt.StandardClaims
}

// TokenCreate generates a new token with a 24 hour timeout, returning the token as a string
func (s *Server) TokenCreate(dc *DataClaim) (token string, err error) {
	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	dc.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expiresAt,
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodRS256, *dc)

	token, err = tokenClaim.SignedString(s.privateKey)
	if err != nil {
		err = errors.Wrap(err, "failed to generate signed token")
		return
	}
	return
}

// TokenRead takes a token string and turns it into a DataClaim
func (s *Server) TokenRead(token string) (*DataClaim, error) {

	dc := &DataClaim{}

	parsedToken, err := jwt.ParseWithClaims(token, dc, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse token")
	}

	dc, ok := parsedToken.Claims.(*DataClaim)
	if !ok && !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token provided")
	}
	return dc, nil
}
