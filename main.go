package main

import (
	"flag"
	"log"

	"github.com/schachte/customraft/config"
	"github.com/schachte/customraft/server"
)

var (
	bootstrap       = flag.Bool(config.BOOTSTRAP, false, "Creates and initializes files if needed")
	serviceBindPort = flag.Int(config.HTTP_PORT, 0, "The port that the HTTP server will listen on")
	raftBindPort    = flag.Int(config.RAFT_PORT, 0, "The port that the Raft server will listen on")
	baseLoc         = flag.String(config.BASE_LOCATION, "/tmp/raft", "Default node identifier for Raft")
	serviceBindAddr = flag.String(config.HTTP_ADDR, "localhost", "The IP address or loopback address of the HTTP server")
	raftBindAddr    = flag.String(config.RAFT_ADDR, "localhost", "The IP address or loopback address of the Raft server")
	nodeIdentifier  = flag.String(config.NODE_IDENTIFIER, "default-identifier", "Default node identifier for Raft")
)

func main() {
	flag.Parse()

	raftConfiguration := &config.RaftConfig{
		Bootstrap:       *bootstrap,
		NodeIdentifier:  *nodeIdentifier,
		StorageLocation: *baseLoc,
		BindAddr:        *raftBindAddr,
		BindPort:        *raftBindPort,
	}

	log.Printf("Raft Configuration: %v\n", raftConfiguration)

	httpConfiguration := &config.HttpConfig{
		BindAddr:       *serviceBindAddr,
		BindPort:       *serviceBindPort,
		NodeIdentifier: *nodeIdentifier,
	}

	log.Printf("HTTP Configuration: %v\n", httpConfiguration)

	raft, err := config.SetupRaft(raftConfiguration)
	if err != nil {
		log.Fatal(err)
	}

	server.InitializeServer(httpConfiguration, raft)
}
