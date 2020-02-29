package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var (
	signingKey = []byte("23jøˆ∆™3oeij23ojˆø™ˆ£∆")
)

// DataClaim represents token data
type DataClaim struct {
	Data map[string]interface{}
	jwt.StandardClaims
}

// TokenSign should be used once
func TokenSign(key []byte) {
	signingKey = key
}

// TokenCreate generates a new token with a 24 hour timeout, returning the token as a string
func TokenCreate(dc *DataClaim) (token string, err error) {
	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	dc.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expiresAt,
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, *dc)

	token, err = tokenClaim.SignedString(signingKey)
	if err != nil {
		err = errors.Wrap(err, "failed to generate signed token")
		return
	}
	return
}

// TokenRead takes a token string and turns it into a DataClaim
func TokenRead(token string) (dc *DataClaim, err error) {
	dc = &DataClaim{}
	parsedToken, err := jwt.ParseWithClaims(token, dc, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})
	if err != nil {
		err = errors.Wrap(err, "failed to parse token")
		return
	}

	dc, ok := parsedToken.Claims.(*DataClaim)
	if ok && parsedToken.Valid {
		return
	}
	err = fmt.Errorf("invalid token")
	return
}
