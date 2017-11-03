package internal

import (
	"context"

	"github.com/pokstad/poki/pb"
	"google.golang.org/grpc"
)

// Server wraps the core with a grpc interface
type Server struct{}

// CreateDoc saves the document on the server
func (s *Server) CreateDoc(context.Context, *pb.Document) (*pb.DocumentRevision, error) {
	return nil, nil
}

// Register will associate this instance with the provided grpc server
func (s *Server) Register(gs *grpc.Server) {
	pb.RegisterPokiServer(gs, s)
}
