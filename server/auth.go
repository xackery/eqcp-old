package server

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

// AuthData represents authentication token claim data
type AuthData struct {
	AccountID int64
}

// AuthPermission is used internally to map access rights
type AuthPermission struct {
	dataClaim          *DataClaim
	accountID          int64
	status             int64
	label              string
	fields             []string
	isSelfOnly         bool
	isLoginNotRequired bool
	isAllFieldsOK      bool
	name               string
}

// AuthFromContext returns a claim if auth is passed
func (s *Server) AuthFromContext(ctx context.Context, endpoint string, scope string) (*AuthPermission, error) {

	ap := new(AuthPermission)
	ap.name = "unknown"

	if len(os.Getenv("EQCP")) > 0 {
		status, err := strconv.Atoi(os.Getenv("EQCP"))
		if err != nil {
			return nil, fmt.Errorf("invalid EQCP: %w", err)
		}
		if status < 1 {
			return nil, fmt.Errorf("invalid EQCP overide")
		}
		ap.status = int64(status)

		ap.label, ap.fields, ap.isSelfOnly, ap.isLoginNotRequired, ap.isAllFieldsOK = s.cfg.Permission(endpoint, scope, ap.status)
		if !ap.isLoginNotRequired {
			return ap, nil
		}
		if ap.label == "" {
			return nil, fmt.Errorf("permission denied")
		}
		return ap, nil
	}

	var token string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md["authorization"]) > 0 {
		token = md["authorization"][0]
	}

	if token == "" {
		//edge case is if ap.IsLoginNotRequired is false
		ap.label, ap.fields, ap.isSelfOnly, ap.isLoginNotRequired, ap.isAllFieldsOK = s.cfg.Permission(endpoint, scope, ap.status)
		if !ap.isLoginNotRequired {
			return nil, fmt.Errorf("must be logged in (no token)")
		}
		return ap, nil
	}

	ad, err := s.AuthRead(token)
	if err != nil {
		return nil, fmt.Errorf("authread: %w", err)
	}

	ap.accountID = ad.AccountID
	if ap.accountID < 1 {
		return nil, fmt.Errorf("token invalid (accountID not set)")
	}

	s.mutex.Lock()
	ap.status, ok = s.authStatus[ap.accountID]
	s.mutex.Unlock()
	if !ok {
		if err = s.db.GetContext(ctx, &ap.status, "SELECT status FROM account WHERE lsaccount_id = ? AND ls_id = 'local'", ap.accountID); err != nil {
			return nil, fmt.Errorf("account status: %w", err)
		}
		s.mutex.Lock()
		s.authStatus[ap.accountID] = ap.status
		s.mutex.Unlock()
	}

	ap.label, ap.fields, ap.isSelfOnly, ap.isLoginNotRequired, ap.isAllFieldsOK = s.cfg.Permission(endpoint, scope, ap.status)
	if ap.label == "" {
		return nil, fmt.Errorf("permission denied")
	}

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
	if dc == nil {
		return nil, fmt.Errorf("claim translation error")
	}
	return dc, nil
}
