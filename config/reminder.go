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

type HttpConfig struct {
	BindAddr       string
	BindPort       int
	NodeIdentifier string
}

type ReminderWrapper struct {
	Mu              sync.Mutex
	CurrentReminder Reminder
	Command         string
}

type Reminder struct {
	Title       string
	Description string
	Completed   bool
}

type RpcInterface struct {
	Raft *raft.Raft
}

func (r *ReminderWrapper) Apply(l *raft.Log) interface{} {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	data := string(l.Data)
	if data == "RETRIEVE" {
		fmt.Println(r.CurrentReminder)
		fmt.Printf("Retrieving %s\n", r.CurrentReminder.Title)
		return r.CurrentReminder.Title
	}
	fmt.Println("Apply being called successfully!")
	deserializedReminder := &proto.AddReminderRequest{}
	json.Unmarshal(l.Data, &deserializedReminder)
	r.CurrentReminder = Reminder{
		Title:       deserializedReminder.Title,
		Description: deserializedReminder.Description,
		Completed:   deserializedReminder.Completed,
	}
	fmt.Printf("Current reminder is: %s\n", r.CurrentReminder.Title)
	return r.CurrentReminder
}

func (r *ReminderWrapper) Snapshot() (raft.FSMSnapshot, error) {
	fmt.Println("Snapshot Ignore")
	return nil, nil
}

func (r *ReminderWrapper) Restore(s io.ReadCloser) error {
	fmt.Println("Restore Ignore")
	return nil
}

func (e *RpcInterface) WriteReminder(context context.Context, req *proto.AddReminderRequest) (*proto.AddReminderResponse, error) {
	serializedData, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	applyFuture := e.Raft.Apply(serializedData, time.Second)
	if err := applyFuture.Error(); err != nil {
		log.Fatal(err)
	}
	log.Println("Applied!")
	return &proto.AddReminderResponse{
		CommitIndex: 400,
	}, nil
}

func (e *RpcInterface) RetrieveLatestReminder(context context.Context, req *proto.GetLatestReminderRequest) (*proto.GetLatestReminderResponse, error) {
	result := e.Raft.Apply([]byte("RETRIEVE"), time.Second)
	if err := result.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The value returned is: %s\n", result.Response().(string))
	return &proto.GetLatestReminderResponse{
		ReadAtIndex: 0,
		Reminder:    result.Response().(string),
	}, nil
}

func (e *RpcInterface) RetrieveAllReminders(context context.Context, req *proto.GetAllRemindersRequest) (*proto.GetAllRemindersResponse, error) {
	fmt.Println("This is a retrieve all reminders")
	return nil, nil
}
