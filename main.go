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
var runAsGrpc = flag.Bool("grpc", false, "Initialize grpc server")

func initFlags() {
	flag.BoolVar(runAsClient, "c", false, "Enables running application in client-mode")
	flag.BoolVar(runAsFollower, "f", false, "Enables running application in follower-mode")
	flag.BoolVar(runAsGrpc, "g", false, "Enables running application as grpc server")
}

func main() {
	initFlags()
	flag.Parse()

	if *runAsFollower {
		raftPort := 7000
		runFollower(raftPort)
		return
	}

	if *runAsClient {
		runClient()
		return
	}

	if *runAsGrpc {

		return
	}

	fsm := &config.ReminderService{
		Reminders: []config.Reminder{},
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

	listener, grpcServer, err := newGrpcServer(r)

	if err != nil {
		log.Fatal(err)
	}

	go func(grpcServer *grpc.Server, l net.Listener) {
		grpcServer.Serve(l)
	}(grpcServer, listener)

	go logLeader(r)
	initiateRaftMuckery(r)
	time.Sleep(30 * time.Second)
}

func initiateRaftMuckery(r *raft.Raft) {
	fmt.Println("Waiting for second node to come online...")
	time.Sleep(5 * time.Second)

	r.AddVoter(raft.ServerID("node_2"), raft.ServerAddress("127.0.0.1:7000"), 0, 5*time.Second)
	fmt.Println("Waiting 5 seconds before shutting down master node...")

	time.Sleep(15 * time.Second)
	r.RemoveServer(raft.ServerID("node_1"), 0, 5*time.Second)
	fmt.Println("Master node removed!")
}

func runFollower(raftPort int) {
	fsm := &config.ReminderService{
		Reminders: []config.Reminder{},
	}

	raftConfig := &config.RaftConfig{
		BindAddr:           "127.0.0.1",
		BindPort:           raftPort,
		FiniteStateMachine: fsm,
		NodeIdentifier:     "node_2",
		StorageLocation:    "/tmp/raft",
	}

	r, err := config.SetupRaft(raftConfig, false)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	go logLeader(r)
	for {
	}
}

func runClient() {
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
	fmt.Println(fmt.Printf("Response from the server is: %s", resp))
}

func newGrpcServer(r *raft.Raft) (net.Listener, *grpc.Server, error) {
	var opts []grpc.ServerOption
	gsrv := grpc.NewServer(opts...)
	proto.RegisterReminderServiceServer(gsrv, &config.ReminderService{
		Raft: r,
	})

	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}
	return l, gsrv, nil
}

func logLeader(r *raft.Raft) {
	t := time.NewTicker(time.Duration(5) * time.Second)
	for {
		select {
		case <-t.C:
			if r.Leader() != "" {
				fmt.Println(r.LeaderWithID())
			}
		}
	}
}
