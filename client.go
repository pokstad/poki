package poki

import (
	"context"
	"net"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
	"github.com/pokstad/poki/pb"
	"google.golang.org/grpc"
)

// Client connects to a poki server and allows requests to be made
type Client struct {
	client pb.PokiClient
}

// DialServer establishes a connection to the poki server
func DialServer(ctx context.Context, addr net.Addr) (*Client, error) {
	cc, err := grpc.DialContext(ctx, addr.String())
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return &Client{
		client: pb.NewPokiClient(cc),
	}, nil
}

// StoreDocument will attempt to store the document in the poki server. Upon
// successful completion, a revision ID will be returned.
func (c *Client) StoreDocument(id string, payload *any.Any) (string, error) {
	return "", nil
}
