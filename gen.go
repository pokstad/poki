package poki

//go:generate protoc -I=$GOPATH/src/github.com/pokstad/poki/pb --go_out=plugins=grpc:$GOPATH/src $GOPATH/src/github.com/pokstad/poki/pb/poki.proto
