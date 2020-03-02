package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/xackery/eqcp/pb"
)

// LoginServerList lists servers
func (s *Server) LoginServerList(ctx context.Context, req *pb.LoginServerListRequest) (*pb.LoginServerListResponse, error) {
	if !s.isLoginServerUp {
		return nil, fmt.Errorf("loginserver api not availabile")
	}
	apiResp, err := s.resty.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(s.cfg.LoginServer.APIToken).
		Get(fmt.Sprintf("http://%s/v1/servers/list", s.cfg.LoginServer.WebAPIHost))
	if err != nil {
		return nil, errors.Wrap(err, "loginserver api")
	}

	var messagePayload struct {
		Message string `json:"message"`
	}

	if apiResp.StatusCode() != 200 {

		if err = json.Unmarshal(apiResp.Body(), &messagePayload); err != nil {
			return nil, errors.Wrap(err, "decode response")
		}
		return nil, fmt.Errorf("loginserver api: %s", messagePayload.Message)
	}

	var payload []struct {
		Localip         string `json:"local_ip"`
		Playersonline   int64  `json:"players_online"`
		Remoteip        string `json:"remote_ip"`
		Serverlistid    int64  `json:"server_list_id"`
		Serverlongname  string `json:"server_long_name"`
		Servershortname string `json:"server_short_name"`
		Serverstatus    int64  `json:"server_status"`
		Zonesbooted     int64  `json:"zones_booted"`
	}

	resp := &pb.LoginServerListResponse{}
	if err = json.Unmarshal(apiResp.Body(), &payload); err != nil {
		if err = json.Unmarshal(apiResp.Body(), &messagePayload); err != nil {
			return nil, errors.Wrap(
				err, "decode response")
		}
		if messagePayload.Message != "There were no results found" {
			return nil, fmt.Errorf("%s", messagePayload.Message)
		}
		return resp, nil
	}

	for _, srv := range payload {
		server := &pb.Server{
			Serverlongname:  srv.Serverlongname,
			Servershortname: srv.Servershortname,
			Serverlistid:    srv.Serverlistid,
			Serverstatus:    srv.Serverstatus,
			Zonesbooted:     srv.Zonesbooted,
			Localip:         srv.Localip,
			Remoteip:        srv.Remoteip,
			Playersonline:   srv.Playersonline,
		}
		resp.Servers = append(resp.Servers, server)
	}

	return resp, nil
}

// LoginServerLogin does a login request
func (s *Server) LoginServerLogin(ctx context.Context, req *pb.LoginServerLoginRequest) (*pb.LoginServerLoginResponse, error) {
	if !s.isLoginServerUp {
		return nil, fmt.Errorf("loginserver api not availabile")
	}
	client := resty.New()
	apiResp, err := client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(s.cfg.LoginServer.APIToken).
		SetBody(req).
		Post(fmt.Sprintf("http://%s/v1/account/credentials/validate/local", s.cfg.LoginServer.WebAPIHost))
	if err != nil {
		return nil, errors.Wrap(err, "loginserver api")
	}

	if apiResp.StatusCode() != 200 {
		var messagePayload struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}

		if err = json.Unmarshal(apiResp.Body(), &messagePayload); err != nil {
			return nil, errors.Wrap(err, "decode response")
		}
		msg := messagePayload.Message
		if msg == "" {
			msg = messagePayload.Error
		}
		return nil, fmt.Errorf("loginserver api: %s", msg)
	}

	var payload struct {
		Data struct {
			AccountID int `json:"account_id"`
		} `json:"data"`
		Message string `json:"message"`
	}

	if err = json.Unmarshal(apiResp.Body(), &payload); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}

	if payload.Data.AccountID < 1 {
		return nil, fmt.Errorf("failed to login")
	}

	token, err := s.AuthCreate(&AuthData{AccountID: int64(payload.Data.AccountID)})
	if err != nil {
		return nil, errors.Wrap(err, "create token")
	}

	return &pb.LoginServerLoginResponse{Token: token}, nil
}

// LoginServerLogout does a logout request
func (s *Server) LoginServerLogout(ctx context.Context, req *pb.LoginServerLogoutRequest) (*pb.LoginServerLogoutResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

// LoginServerCreate makes a new account
func (s *Server) LoginServerCreate(ctx context.Context, req *pb.LoginServerCreateRequest) (*pb.LoginServerCreateResponse, error) {
	if !s.isLoginServerUp {
		return nil, fmt.Errorf("loginserver api not availabile")
	}
	client := resty.New()
	if req.Username == "" {
		return nil, fmt.Errorf("username must be set")
	}
	if req.Password == "" {
		return nil, fmt.Errorf("password must be set")
	}
	if req.Email == "" {
		return nil, fmt.Errorf("email must be set")
	}
	apiResp, err := client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(s.cfg.LoginServer.APIToken).
		SetBody(req).
		Post(fmt.Sprintf("http://%s/v1/account/create", s.cfg.LoginServer.WebAPIHost))
	if err != nil {
		return nil, errors.Wrap(err, "loginserver api")
	}

	if apiResp.StatusCode() != 200 {
		var messagePayload struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}

		if err = json.Unmarshal(apiResp.Body(), &messagePayload); err != nil {
			return nil, errors.Wrap(err, "decode response")
		}
		msg := messagePayload.Message
		if msg == "" {
			msg = messagePayload.Error
		}
		return nil, fmt.Errorf("loginserver api: %s", msg)
	}

	var payload struct {
		Data struct {
			AccountID int `json:"account_id"`
		} `json:"data"`
		Message string `json:"message"`
	}

	if err = json.Unmarshal(apiResp.Body(), &payload); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}

	if payload.Data.AccountID < 1 {
		return nil, fmt.Errorf("failed to login")
	}

	resp := &pb.LoginServerCreateResponse{}
	return resp, nil
}
