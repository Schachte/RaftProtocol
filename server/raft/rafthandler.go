package raft

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/config"
)

type request struct {
	NodeIdentifier string
	FullAddress    string
}

type response struct {
	Message string
}

type server struct {
	Raft         *raft.Raft
	Identifier   string
	ServerConfig *config.HttpConfig
}

func New(httpConfig *config.HttpConfig, r *raft.Raft) *server {
	return &server{
		Raft:         r,
		Identifier:   httpConfig.NodeIdentifier,
		ServerConfig: httpConfig,
	}
}

func (s *server) JoinCluster(w http.ResponseWriter, r *http.Request) {
	if s.Raft.State() != raft.Leader {
		fmt.Fprintf(w, "Must be leader\n")
		return
	}

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		return
	}
	var incomingPayload request
	json.NewDecoder(r.Body).Decode(&incomingPayload)
	fmt.Println(incomingPayload)
	fmt.Println(incomingPayload.FullAddress)
	f := s.Raft.AddVoter(raft.ServerID(incomingPayload.NodeIdentifier), raft.ServerAddress(incomingPayload.FullAddress), 0, time.Second)
	if e := f.Error(); e != nil {
		log.Fatal(e)
		return
	}

	Configuration := s.Raft.GetConfiguration().Configuration()
	for _, server := range Configuration.Servers {
		fmt.Println(server.Suffrage)
	}

	handleRequest(w, r, func() ([]byte, error) {
		payload := &response{
			Message: fmt.Sprintf("Added new node to %s - %s @ %s.", s.Identifier, incomingPayload.NodeIdentifier, incomingPayload.FullAddress),
		}
		return json.Marshal(payload)
	})
}

func (s *server) RemoveFromCluster(w http.ResponseWriter, r *http.Request) {
	if s.Raft.State() != raft.Leader {
		fmt.Fprintf(w, "Must be leader\n")
		return
	}

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		return
	}

	var incomingPayload request
	json.NewDecoder(r.Body).Decode(&incomingPayload)
	fmt.Println(incomingPayload)
	f := s.Raft.RemoveServer(raft.ServerID(incomingPayload.NodeIdentifier), 0, time.Second)
	if e := f.Error(); e != nil {
		log.Fatal(e)
		return
	}

	handleRequest(w, r, func() ([]byte, error) {
		payload := &response{
			Message: fmt.Sprintf("%s removed.\n", incomingPayload.NodeIdentifier),
		}
		return json.Marshal(payload)
	})
}

func (s *server) Stats(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, func() ([]byte, error) {
		return json.Marshal(s.Raft.Stats())
	})
}

func handleRequest(w http.ResponseWriter, r *http.Request, fn func() ([]byte, error)) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := fn()
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
