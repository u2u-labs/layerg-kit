package pb

import (
	_ "github.com/u2u-labs/go-layerg-common/api"
	_ "github.com/u2u-labs/go-layerg-common/rtapi"
)

//go:generate protoc -I.. -I. -I../vendor -I../vendor/github.com/u2u-labs/go-layerg-common --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative peer.proto
