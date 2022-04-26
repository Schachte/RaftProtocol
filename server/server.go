package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/config"
	ra "github.com/schachte/customraft/server/raft"
	re "github.com/schachte/customraft/server/reminder"
)

func InitializeServer(httpConfig *config.HttpConfig, raft *raft.Raft) {
	go re.New(*httpConfig, raft)

	raftHandler := ra.New(httpConfig, raft)
	http.HandleFunc("/raft/add", raftHandler.JoinCluster)
	http.HandleFunc("/raft/remove", raftHandler.RemoveFromCluster)
	http.HandleFunc("/raft/stats", raftHandler.Stats)

	log.Println("Listening...")
	httpServerConfig := fmt.Sprintf("%s:%d", httpConfig.BindAddr, httpConfig.BindPort)
	log.Printf("We are going to start the HTTP server on %s\n", httpServerConfig)
	http.ListenAndServe(httpServerConfig, nil)
	for {
	}
}
