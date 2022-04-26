package reminder

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
	fmt.Printf("Running grpc server on port %d", GRPC_PORT)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterReminderServiceServer(s, &config.RpcInterface{
		Raft: r,
	})

	log.Println("Serving GRPC service...")
	log.Println("Served!")
	log.Printf("GRPC Server IP: %s\n", httpConfig.BindAddr)
	log.Printf("GRPC Server Port: %d\n", GRPC_PORT)

	s.Serve(lis)
}
