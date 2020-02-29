package auth

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
)

// Data represents authentication token claim data
type Data struct {
	AccountID int64
}

// FromContext returns a claim if auth is passed
func FromContext(ctx context.Context) *DataClaim {
	var token string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md["authorization"]) > 0 {
		token = md["authorization"][0]
	}
	if token == "" {
		return nil
	}

	dc, err := TokenRead(token)
	if err != nil {
		log.Debug().Err(err).Msg("token read")
	}
	return dc
}

// Create generates a new token with a 24 hour timeout, returning the token as a string
func Create(ad *Data) (token string, err error) {
	dc := marshalDataClaim(ad)
	token, err = TokenCreate(dc)
	if err != nil {
		return
	}
	return
}

// Read takes a token string and turns it into Auth
func Read(token string) (ad *Data, err error) {
	dc, err := TokenRead(token)
	if err != nil {
		return
	}
	ad = unmarshalDataClaim(dc)
	return
}

func unmarshalDataClaim(dc *DataClaim) (ad *Data) {
	ad = &Data{}
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

func marshalDataClaim(ad *Data) (dc *DataClaim) {
	dc = &DataClaim{
		Data: make(map[string]interface{}),
	}
	if ad == nil {
		return
	}
	dc.Data["accountID"] = ad.AccountID
	return
}
