package config

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/proto"
)

type ReminderService struct {
	StorageHandler PersistenceStore
	Mu             sync.Mutex
	Reminders      []Reminder
	Raft           *raft.Raft
}

type Reminder struct {
	Title       string
	Description string
	Completed   bool
}

func (r *ReminderService) Apply(l *raft.Log) interface{} {
	r.Mu.Lock()
	fmt.Println("Getting called via Raft")
	fmt.Println(l.Data)
	r.Mu.Unlock()
	return nil
}

func (r *ReminderService) Snapshot() (raft.FSMSnapshot, error) {
	fmt.Println("Snapshot Ignore")
	return nil, nil
}

func (r *ReminderService) Restore(s io.ReadCloser) error {
	fmt.Println("Restore Ignore")
	return nil
}

// Because we are treating the reminder service as the FSM for RAFT, we must implement the methods to make it polymorphic
var _ raft.FSM = &ReminderService{}

// WriteReminder is responsible for appending a reminder
// request to the persistence store for the current Event
func (e *ReminderService) WriteReminder(context context.Context, req *proto.AddReminderRequest) (*proto.AddReminderResponse, error) {
	e.Raft.Apply([]byte("tester"), time.Second)
	// Take an incoming request and write it to the persistence store of Reminder instance
	// Because this is modifying the state of our application, we need to go through RAFT to ensure the write
	// was cross-replicated across all the nodes without any issues

	// Construct a reminder from the incoming GRPC request
	// newReminder := &cfg.Reminder{
	// 	Title:       req.GetTitle(),
	// 	Description: req.GetDescription(),
	// 	Completed:   false,
	// }

	// err := e.StorageHandler.Store(newReminder)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(e.Raft.LeaderWithID())
	fmt.Println("This is a write reminder")
	return &proto.AddReminderResponse{
		CommitIndex: 400,
	}, nil
}

// RetrieveLatestReminder is responsible for retrieving the most
// recently appended reminder that was added to the reminder log
func (e *ReminderService) RetrieveLatestReminder(context context.Context, req *proto.GetLatestReminderRequest) (*proto.GetLatestReminderResponse, error) {
	fmt.Println("This is a retrieve latest reminder")
	return nil, nil
}

// RetrieveAllReminders is responsible for retrieving all reminders
// that are stored in the entire reminder list
func (e *ReminderService) RetrieveAllReminders(context context.Context, req *proto.GetAllRemindersRequest) (*proto.GetAllRemindersResponse, error) {
	fmt.Println("This is a retrieve all reminders")
	return nil, nil
}
