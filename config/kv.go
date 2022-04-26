package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/proto"
)

type Entry struct {
	Key   string
	Value string
}

type RpcInterface struct {
	Raft *raft.Raft
}

type FSM struct {
	Mu      sync.Mutex
	Entries map[string]string
}

func (r *FSM) Apply(l *raft.Log) interface{} {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	incomingCommand := &RaftCommand{}
	json.Unmarshal(l.Data, &incomingCommand)

	switch incomingCommand.Operation {
	case ADD_KEY:
		fmt.Println("ADDING KEY")
		incomingEntry := incomingCommand.Value
		err := r.add(incomingEntry.Key, incomingEntry.Value)
		if err != nil {
			log.Fatal("There was a problem processing: %v", err)
		}
		return r.Entries
	case REMOVE_KEY:
		fmt.Println("REMOVING KEY")
		incomingEntry := incomingCommand.Value
		r.delete(incomingEntry.Key)
		return r.Entries
	case UPDATE_KEY:
		fmt.Println("UPDATING KEY")
		incomingEntry := incomingCommand.Value
		err := r.add(incomingEntry.Key, incomingEntry.Value)
		if err != nil {
			log.Fatal("There was a problem processing: %v", err)
		}
		return r.Entries
	case RETRIEVE_VALUE:
		fmt.Println("RETRIEVING VALUE")
		incomingEntry := incomingCommand.Value
		value := r.get(incomingEntry.Key)
		log.Printf("The value is: %s\n", value)
		return r.Entries
	}

	log.Fatal("Invalid command was passed")
	return nil
}

func (r *FSM) Snapshot() (raft.FSMSnapshot, error) {
	fmt.Println("Snapshot Ignore")
	return nil, nil
}

func (r *FSM) Restore(s io.ReadCloser) error {
	fmt.Println("Restore Ignore")
	return nil
}

func (e *RpcInterface) AddEntry(_ context.Context, req *proto.AddEntryRequest) (*proto.KeyResponse, error) {
	log.Println("Called")
	cmd := &RaftCommand{
		Operation: ADD_KEY,
		Value: Entry{
			Key:   req.GetKey(),
			Value: req.GetValue(),
		},
	}

	log.Println("Command")
	log.Println(cmd)

	serializedData, err := json.Marshal(cmd)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Currently calling raft!")
	log.Println(req)
	res := e.Raft.Apply(serializedData, time.Second)
	if e := res.Error(); e != nil {
		log.Fatal(e)
	}

	log.Printf("Key %s added successfully to value %s\n", req.GetKey(), req.GetValue())
	return &proto.KeyResponse{
		Key:   req.GetKey(),
		Value: req.GetValue(),
	}, nil
}

func (e *RpcInterface) UpdateEntry(_ context.Context, req *proto.SetValueRequest) (*proto.KeyResponse, error) {
	cmd := &RaftCommand{
		Operation: UPDATE_KEY,
		Value: Entry{
			Key:   req.GetKey(),
			Value: req.GetValue(),
		},
	}

	serializedData, err := json.Marshal(cmd)
	if err != nil {
		log.Fatal(err)
	}

	res := e.Raft.Apply(serializedData, time.Second)
	if e := res.Error(); e != nil {
		log.Fatal(e)
	}

	log.Printf("Key %s updated successfully to value %s\n", req.GetKey(), req.GetValue())
	return &proto.KeyResponse{
		Key:   req.GetKey(),
		Value: req.GetValue(),
	}, nil
}

func (e *RpcInterface) DeleteEntry(_ context.Context, req *proto.DeleteKeyRequest) (*proto.KeyResponse, error) {
	cmd := &RaftCommand{
		Operation: REMOVE_KEY,
		Value: Entry{
			Key: req.GetKey(),
		},
	}

	serializedData, err := json.Marshal(cmd)
	if err != nil {
		log.Fatal(err)
	}

	res := e.Raft.Apply(serializedData, time.Second)
	if e := res.Error(); e != nil {
		log.Fatal(e)
	}

	log.Printf("Key %s removed successfully\n", req.GetKey())
	return &proto.KeyResponse{
		Key: req.GetKey(),
	}, nil
}

func (e *RpcInterface) GetEntry(_ context.Context, req *proto.GetValueRequest) (*proto.KeyResponse, error) {
	cmd := &RaftCommand{
		Operation: RETRIEVE_VALUE,
		Value: Entry{
			Key: req.GetKey(),
		},
	}

	serializedData, err := json.Marshal(cmd)
	if err != nil {
		log.Fatal(err)
	}

	res := e.Raft.Apply(serializedData, time.Second)
	if e := res.Error(); e != nil {
		log.Fatal(e)
	}

	value := res.Response().(map[string]string)[req.GetKey()]
	log.Printf("Value %s retrieved successfully from the key %s\n", value, req.GetKey())
	return &proto.KeyResponse{
		Key:   req.GetKey(),
		Value: value,
	}, nil
}

func (fsm *FSM) add(key, value string) error {
	fsm.Entries[key] = value
	return nil
}

func (fsm *FSM) get(key string) string {
	return fsm.Entries[key]
}

func (fsm *FSM) delete(key string) {
	delete(fsm.Entries, key)
}
