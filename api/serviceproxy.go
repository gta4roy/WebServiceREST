package api

import (
	"AddressService/log"
	"AddressService/model"
	"AddressService/protocol"
	"context"

	"google.golang.org/grpc"
)

type ServiceProxy struct {
	clientConnect        *grpc.ClientConn
	addressServiceClient protocol.AddressServiceClient
}

func (proxy *ServiceProxy) initialiseConnection() {
	opts := grpc.WithInsecure()

	var err error
	proxy.clientConnect, err = grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Error.Println(err)
	}

	proxy.addressServiceClient = protocol.NewAddressServiceClient(proxy.clientConnect)
}

func (proxy *ServiceProxy) CloseClientConnect() {
	proxy.clientConnect.Close()
}

func (proxy *ServiceProxy) AddAddress(person model.PersonModel) model.ResponseModel {

	var data protocol.Person
	data.Name = person.Name
	data.Id = person.Id
	data.Phone = person.Phone
	data.Address = person.Address
	data.City = person.City
	data.Pin = person.Pin

	resp, _ := proxy.addressServiceClient.Add(context.Background(), &data)

	var response model.ResponseModel
	if resp.IsSuccess {
		response.Status = model.CODE_SUCCESS
		response.Message = model.MSG_SUCCESS_SAVE
	}
	return response
}

func (proxy *ServiceProxy) ModifyAddress(personId string, person model.PersonModel) model.ResponseModel {

}

func (proxy *ServiceProxy) ListOfAddress() []model.PersonModel {

}

func (proxy *ServiceProxy) DeleteAddress(personId string) model.ResponseModel {

}

func (proxy *ServiceProxy) SearchAddress(personId string) []model.PersonModel {

}
