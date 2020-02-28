package server

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/pb"
)

// NpcSearch implements SCRUD endpoints
func (s *Server) NpcSearch(ctx context.Context, req *pb.NpcSearchRequest) (*pb.NpcSearchResponse, error) {
	return &pb.NpcSearchResponse{
		Limit: 1,
	}, nil
}

// NpcCreate implements SCRUD endpoints
func (s *Server) NpcCreate(ctx context.Context, req *pb.NpcCreateRequest) (*pb.NpcCreateResponse, error) {
	return &pb.NpcCreateResponse{
		Id: 1234,
	}, nil
}

// NpcRead implements SCRUD endpoints
func (s *Server) NpcRead(ctx context.Context, req *pb.NpcReadRequest) (*pb.NpcReadResponse, error) {
	log.Info().Msg("read request")
	return &pb.NpcReadResponse{
		Npc: &pb.Npc{
			ID:   1234,
			Name: "foo!",
		},
	}, nil
}

// NpcUpdate implements SCRUD endpoints
func (s *Server) NpcUpdate(ctx context.Context, req *pb.NpcUpdateRequest) (*pb.NpcUpdateResponse, error) {
	return &pb.NpcUpdateResponse{}, nil
}

// NpcDelete implements SCRUD endpoints
func (s *Server) NpcDelete(ctx context.Context, req *pb.NpcDeleteRequest) (*pb.NpcDeleteResponse, error) {
	return &pb.NpcDeleteResponse{}, nil
}

// NpcPatch implements SCRUD endpoints
func (s *Server) NpcPatch(ctx context.Context, req *pb.NpcPatchRequest) (*pb.NpcPatchResponse, error) {
	return &pb.NpcPatchResponse{}, nil
}
