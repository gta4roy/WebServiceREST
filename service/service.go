package main

import (
	"AddressService/log"
	"AddressService/protocol"
	"AddressService/store"
	"AddressService/util"
	"context"
	"net"

	"google.golang.org/grpc"
)

const (
	port      = ":50051"
	aggregate = "Address"
	event     = "AddressAdded"
)

type server struct {
	storage store.AddressDB
}

func (s *server) Add(ctx context.Context, in *protocol.Person) (*protocol.ServiceResponse, error) {

	var response protocol.ServiceResponse
	response = s.storage.Add(in)
	log.Trace.Println("Adding an address ..  :")
	go normalPublishEvent()
	return &response, nil
}

func (s *server) Modify(ctx context.Context, in *protocol.ModifyPerson) (*protocol.ServiceResponse, error) {

	var response protocol.ServiceResponse
	response = s.storage.Modify(in)
	log.Trace.Println("Modifying an address ..  :")
	go normalPublishEvent()
	return &response, nil
}

func (s *server) ListAll(ctx context.Context, in *protocol.EmptyParams) (*protocol.PersonList, error) {

	var listOfAddress protocol.PersonList
	listOfAddress = s.storage.ListAll()
	log.Trace.Println("Search an address ..  :")
	go normalPublishEvent()
	return &listOfAddress, nil
}

func (s *server) Delete(ctx context.Context, in *protocol.PersonID) (*protocol.ServiceResponse, error) {

	var response protocol.ServiceResponse
	response = s.storage.Delete(in)
	log.Trace.Println("Deleting an address ..  :")
	go normalPublishEvent()
	return &response, nil
}

func (s *server) Search(ctx context.Context, in *protocol.PersonID) (*protocol.PersonList, error) {

	var listOfAddress protocol.PersonList
	listOfAddress = s.storage.Search(in)
	log.Trace.Println("Search an address ..  :")
	go normalPublishEvent()
	return &listOfAddress, nil
}

func normalPublishEvent() {
	//natsConnection, _ := nats.Connect(nats.DefaultURL)
	//log.Println("Connected to " + nats.DefaultURL)
	//defer natsConnection.Close()

	subjectNotQueue := "Order.TestEvent"
	//data := "String message from the client"
	//natsConnection.Publish(subjectNotQueue, []byte(data))
	log.Trace.Println("Published message on suject :" + subjectNotQueue)
}
func init() {
	// Parse log level from command line
	logLevel := util.GetProperty(util.LogLevel)
	// Calling the SetLogLevel with the command-line argument
	log.SetLogLevel(logLevel, "server.txt")
	log.Trace.Println("Loging initialised")

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error.Println("Failed to listen %v ", err)
	}
	s := grpc.NewServer()
	protocol.RegisterAddressServiceServer(s, &server{})
	log.Trace.Println("Server listening on the port :", port)
	s.Serve(lis)
}
