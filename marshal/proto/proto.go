package proto

import (
	protobuf "github.com/golang/protobuf/proto"
	"github.com/pokstad/poki"
	pb "github.com/pokstad/poki/pb"
)

func MarshalPost(p poki.Post) ([]byte, error) {

	pbpost := &pb.Post{
		Raw: p.Raw,
		Meta: &pb.Post_MetaData{
			Title: p.Meta.Title,
		},
		Path: p.Path,
	}

	return protobuf.Marshal(pbpost)
}
