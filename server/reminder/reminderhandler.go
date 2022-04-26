package reminder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/config"
	"github.com/schachte/customraft/fsm"
)

type request struct {
	Operation int
	Value     config.Reminder
}

type server struct {
	Raft       *raft.Raft
	Identifier string
}

func New(httpConfig config.HttpConfig, r *raft.Raft) *server {
	return &server{
		Raft:       r,
		Identifier: httpConfig.NodeIdentifier,
	}
}

func (s *server) RemoveHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, there\n")
}

func (s *server) RetrieveHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, there\n")
}

func (s *server) AddHandler(w http.ResponseWriter, r *http.Request) {
	if s.Raft.State() != raft.Leader {
		fmt.Fprintf(w, "Error, must be leader\n")
		return
	}

	var incomingPayload request
	json.NewDecoder(r.Body).Decode(&incomingPayload)
	jsonPayload := &fsm.Command{
		Value:     incomingPayload.Value,
		Operation: fsm.ADD_REMINDER,
	}

	byteLoad, err := json.Marshal(jsonPayload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("About to call RAFT")
	s.Raft.Apply(byteLoad, time.Second*3)
	fmt.Println("Called RAFT")
	fmt.Fprintf(w, "Hello, there\n")
}
