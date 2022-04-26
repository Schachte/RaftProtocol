package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/config"
	ra "github.com/schachte/customraft/server/raft"
	store "github.com/schachte/customraft/server/store"
)

func InitializeServer(httpConfig *config.HttpConfig, raft *raft.Raft) {
	go store.New(*httpConfig, raft)

	raftHandler := ra.New(httpConfig, raft)
	http.HandleFunc("/raft/add", raftHandler.JoinCluster)
	http.HandleFunc("/raft/remove", raftHandler.RemoveFromCluster)
	http.HandleFunc("/raft/stats", raftHandler.Stats)

	httpServerConfig := fmt.Sprintf("%s:%d", httpConfig.BindAddr, httpConfig.BindPort)
	log.Printf("We are going to start the HTTP server on %s\n", httpServerConfig)
	http.ListenAndServe(httpServerConfig, nil)
}
