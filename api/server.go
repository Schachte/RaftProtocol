package api

import (
	"context"
	"fmt"

	"github.com/hashicorp/raft"

	"github.com/schachte/customraft/config"
	proto "github.com/schachte/customraft/proto"
)

type RpcInterface struct {
	Raft    raft.Raft
	Service config.ReminderService
}

func (rpc *RpcInterface) WriteReminder(c context.Context, reminder *proto.AddReminderRequest) (*proto.AddReminderResponse, error) {
	fmt.Println("Writing reminder!")
	result, err := rpc.Service.WriteReminder(c, reminder)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (rpc *RpcInterface) RetrieveLatestReminder(c context.Context, reminder *proto.GetLatestReminderRequest) (*proto.GetLatestReminderResponse, error) {
	return nil, nil
}

func (rpc *RpcInterface) RetrieveAllReminders(c context.Context, reminder *proto.GetAllRemindersRequest) (*proto.GetAllRemindersResponse, error) {
	return nil, nil
}
