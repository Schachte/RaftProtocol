package store

import (
	"fmt"
	"log"
	"net"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/config"
	"github.com/schachte/customraft/proto"
	"google.golang.org/grpc"
)

func New(httpConfig config.HttpConfig, r *raft.Raft) {
	GRPC_PORT := httpConfig.BindPort + 10
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", httpConfig.BindAddr, GRPC_PORT))
	fmt.Printf("Running grpc server on port %d\n", GRPC_PORT)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterKeyValueServiceServer(s, &config.RpcInterface{
		Raft: r,
	})
	s.Serve(lis)
}
