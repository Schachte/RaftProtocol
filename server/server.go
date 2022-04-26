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
	reminderHandler := re.New(*httpConfig, raft)
	http.HandleFunc("/reminder/add", reminderHandler.AddHandler)
	http.HandleFunc("/reminder/remove/:index", reminderHandler.RemoveHandler)
	http.HandleFunc("/reminder/retrieve/:index", reminderHandler.RetrieveHandler)

	raftHandler := ra.New(httpConfig, raft)
	http.HandleFunc("/raft/add", raftHandler.JoinCluster)
	http.HandleFunc("/raft/remove", raftHandler.RemoveFromCluster)
	http.HandleFunc("/raft/stats", raftHandler.Stats)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", httpConfig.BindAddr, httpConfig.BindPort), nil))
}
