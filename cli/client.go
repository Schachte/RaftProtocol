package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/schachte/customraft/config"
	"github.com/schachte/customraft/proto"
	"google.golang.org/grpc"
)

var (
	serviceBindPort = flag.Int(config.HTTP_PORT, 0, "The port that the HTTP server will listen on")
	serviceBindAddr = flag.String(config.HTTP_ADDR, "localhost", "The IP address or loopback address of the HTTP server")
	action          = flag.String(config.ACTION, "default-action", "Tells the GRPC server what to do")
)

func main() {
	flag.Parse()

	serverAddr := fmt.Sprintf("%s:%d", *serviceBindAddr, *serviceBindPort+10)
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	grpcClient := proto.NewReminderServiceClient(conn)
	ctx := context.Background()
	req := &proto.AddReminderRequest{
		Title:       "TestTitle",
		Description: "TestDescription",
		Completed:   false,
	}

	req2 := &proto.GetLatestReminderRequest{}

	switch *action {
	case "ADD":
		log.Println("Adding New Reminder...")
		grpcClient.WriteReminder(ctx, req)
		return
	case "REMOVE":
	case "RETRIEVE_VALUE":
		grpcClient.RetrieveLatestReminder(ctx, req2)
		time.Sleep(10 * time.Second)
	}
	fmt.Println("waiting... <cntrl + c> to exit")
	for {
	}
}
