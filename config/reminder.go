package config

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	"github.com/schachte/customraft/proto"
)

type ReminderService struct {
	Mu        sync.Mutex
	Reminders []Reminder
	Raft      *raft.Raft
}

type Reminder struct {
	Title       string
	Description string
	Completed   bool
}

func (r *ReminderService) Apply(l *raft.Log) interface{} {
	fmt.Println("Apply being called successfully!")
	data := strings.Split(string(l.Data), "-")
	completed, err := strconv.ParseBool(data[2])

	if err != nil {
		log.Fatal(err)
	}

	newReminder := &Reminder{
		Title:       data[0],
		Description: data[1],
		Completed:   completed,
	}

	fmt.Println("Getting called via Raft")
	r.Reminders = append(r.Reminders, *newReminder)

	fmt.Printf("Adding %v\n", newReminder)
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
	fmt.Println("The write reminder call has been sent to gRPC")
	fmt.Println(e.Raft.LeaderWithID())
	e.Raft.Apply([]byte(fmt.Sprintf("%s-%s-%t", req.GetTitle(), req.GetDescription(), req.GetCompleted())), time.Second)
	fmt.Println("Applied via raft.. returning")
	fmt.Println(e.Raft.LeaderWithID())
	return &proto.AddReminderResponse{
		CommitIndex: 400,
	}, nil
}

func (e *ReminderService) RetrieveLatestReminder(context context.Context, req *proto.GetLatestReminderRequest) (*proto.GetLatestReminderResponse, error) {
	fmt.Println("This is a retrieve latest reminder")
	return nil, nil
}
func (e *ReminderService) RetrieveAllReminders(context context.Context, req *proto.GetAllRemindersRequest) (*proto.GetAllRemindersResponse, error) {
	fmt.Println("This is a retrieve all reminders")
	return nil, nil
}
