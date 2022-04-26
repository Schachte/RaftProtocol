package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/schachte/customraft/config"
	"github.com/schachte/customraft/proto"
	"google.golang.org/grpc"
)

var (
	serviceBindPort = flag.Int(config.HTTP_PORT, 0, "The port that the HTTP server will listen on")
	serviceBindAddr = flag.String(config.HTTP_ADDR, "localhost", "The IP address or loopback address of the HTTP server")
	action          = flag.String(config.ACTION, "default-action", "Tells the GRPC server what to do")
	key             = flag.String(config.KEY, "", "Key you want to add or remove from the K-V store")
	value           = flag.String(config.VALUE, "", "Value you want to add or remove from the K-V store")
)

func main() {
	flag.Parse()

	serverAddr := fmt.Sprintf("%s:%d", *serviceBindAddr, *serviceBindPort+10)
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	grpcClient := proto.NewKeyValueServiceClient(conn)
	ctx := context.Background()

	parsedAction := config.RpcCommand(*action)

	switch parsedAction {
	case config.ADD_KEY:
		log.Println("Adding/Updating New K/V Pair")
		addEntryReq := &proto.AddEntryRequest{
			Key:   *key,
			Value: *value,
		}
		resp, err := grpcClient.AddEntry(ctx, addEntryReq)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		return
	case config.REMOVE_KEY:
		log.Println("Removing K/V Pair")
		deleteEntryReq := &proto.DeleteKeyRequest{
			Key: *key,
		}
		resp, err := grpcClient.DeleteEntry(ctx, deleteEntryReq)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		return
	case config.RETRIEVE_VALUE:
		log.Println("Retrieving Value from Key")
		retrieveEntryReq := &proto.GetValueRequest{
			Key: *key,
		}
		resp, err := grpcClient.GetEntry(ctx, retrieveEntryReq)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		return
	case config.UPDATE_KEY:
		log.Println("Updating New K/V Pair")
		updateEntryReq := &proto.SetValueRequest{
			Key:   *key,
			Value: *value,
		}
		resp, err := grpcClient.UpdateEntry(ctx, updateEntryReq)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		return
	}

	fmt.Println("waiting... <cntrl + c> to exit")
	for {
	}
}
