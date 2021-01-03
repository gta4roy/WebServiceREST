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
	var data protocol.ModifyPerson
	data.ModifiedPerson = &protocol.Person{}
	data.ModifiedPerson.Name = person.Name
	data.ModifiedPerson.Id = person.Id
	data.ModifiedPerson.Phone = person.Phone
	data.ModifiedPerson.Address = person.Address
	data.ModifiedPerson.City = person.City
	data.ModifiedPerson.Pin = person.Pin
	data.Id = &protocol.PersonID{}
	data.Id.Id = personId

	resp, _ := proxy.addressServiceClient.Modify(context.Background(), &data)

	var response model.ResponseModel
	if resp.IsSuccess {
		response.Status = model.CODE_SUCCESS
		response.Message = model.MSG_SUCCESS_SAVE
	}
	return response

}

func (proxy *ServiceProxy) ListOfAddress() []model.PersonModel {

	var emptyParams protocol.EmptyParams
	resp, _ := proxy.addressServiceClient.ListAll(context.Background(), &emptyParams)

	var listOfPersons []model.PersonModel
	for _, personVal := range resp.PersonsList {
		var data model.PersonModel
		data.Name = personVal.Name
		data.Id = personVal.Id
		data.Phone = personVal.Phone
		data.Address = personVal.Address
		data.City = personVal.City
		data.Pin = personVal.Pin
		listOfPersons = append(listOfPersons, data)
	}
	return listOfPersons
}

func (proxy *ServiceProxy) DeleteAddress(personId string) model.ResponseModel {

	var data protocol.PersonID
	data.Id = personId

	resp, _ := proxy.addressServiceClient.Delete(context.Background(), &data)

	var response model.ResponseModel
	if resp.IsSuccess {
		response.Status = model.CODE_SUCCESS
		response.Message = model.MSG_SUCCESS_SAVE
	}
	return response
}

func (proxy *ServiceProxy) SearchAddress(personId string) []model.PersonModel {

	var data protocol.PersonID
	data.Id = personId
	resp, _ := proxy.addressServiceClient.Search(context.Background(), &data)

	var listOfPersons []model.PersonModel
	for _, personVal := range resp.PersonsList {
		var data model.PersonModel
		data.Name = personVal.Name
		data.Id = personVal.Id
		data.Phone = personVal.Phone
		data.Address = personVal.Address
		data.City = personVal.City
		data.Pin = personVal.Pin
		listOfPersons = append(listOfPersons, data)
	}
	return listOfPersons
}
