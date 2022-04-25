package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/config"
	"github.com/schachte/customraft/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var runAsClient = flag.Bool("client", false, "Run client instead of server")
var runAsFollower = flag.Bool("follower", false, "Run follower instead of leader")

func initFlags() {
	flag.BoolVar(runAsClient, "c", false, "Enables running application in client-mode")
	flag.BoolVar(runAsFollower, "f", false, "Enables running application in follower-mode")
}

func main() {
	initFlags()
	flag.Parse()
	if *runAsFollower {
		storage, err := config.New("/tmp/raft/node_2/storage")
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}

		fsm := &config.ReminderService{
			StorageHandler: storage,
			Reminders:      []config.Reminder{},
		}

		raftConfig := &config.RaftConfig{
			BindAddr:           "127.0.0.1",
			BindPort:           7000,
			FiniteStateMachine: fsm,
			NodeIdentifier:     "node_2",
			StorageLocation:    "/tmp/raft",
		}

		_, err = config.SetupRaft(raftConfig, false)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
		time.Sleep(10 * time.Second)
		fmt.Println("Removing the server node 1!")
		for {
		}
	}

	if *runAsClient {
		fmt.Println("Running in client mode, preparing gRPC server initialization...")
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		conn, err := grpc.Dial("127.0.0.1:4000", opts...)

		if err != nil {
			log.Fatal(err)
			return
		}

		client := proto.NewReminderServiceClient(conn)
		wrr := &proto.AddReminderRequest{
			Title:       "test",
			Description: "Desc",
			Completed:   false,
		}
		ctx := context.Background()
		resp, err := client.WriteReminder(ctx, wrr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)
		return
	}

	storage, err := config.New("/tmp/raft/node_1/storage")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fsm := &config.ReminderService{
		StorageHandler: storage,
		Reminders:      []config.Reminder{},
	}

	raftConfig := &config.RaftConfig{
		BindAddr:           "127.0.0.1",
		BindPort:           6000,
		FiniteStateMachine: fsm,
		NodeIdentifier:     "node_1",
		StorageLocation:    "/tmp/raft",
	}

	r, err := config.SetupRaft(raftConfig, true)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	var opts []grpc.ServerOption
	grpcServer, err := newGrpcServer(opts...)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func(grpcServer *grpc.Server, l net.Listener) {
		grpcServer.Serve(l)
	}(grpcServer, l)

	go logLeader(r)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("This called")
	r.AddVoter(raft.ServerID("node_2"), raft.ServerAddress("127.0.0.1:7000"), 0, 5*time.Second)
	r.RemoveServer(raft.ServerID("node_1"), 0, 5*time.Second)
	for {
	}
}

func newGrpcServer(opts ...grpc.ServerOption) (*grpc.Server, error) {
	gsrv := grpc.NewServer(opts...)
	proto.RegisterReminderServiceServer(gsrv, &config.ReminderService{})
	return gsrv, nil
}

func logLeader(r *raft.Raft) {
	t := time.NewTicker(time.Duration(1) * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println(r.Leader())
		}
	}
}
