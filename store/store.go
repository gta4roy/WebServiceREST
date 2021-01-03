package store

import (
	"AddressService/protocol"
)

type AddressDB struct {
	addressList []protocol.Person
}

func (store *AddressDB) Add(personDetails *protocol.Person) protocol.ServiceResponse {
	store.addressList = append(store.addressList, *personDetails)
	var serviceResponse protocol.ServiceResponse
	serviceResponse.IsSuccess = true
	serviceResponse.Error = ""

	return serviceResponse
}

func (store *AddressDB) ListAll() []protocol.Person {

	var personList protocol.PersonList
	//*personList.PersonsList = store.addressList
	return personList
}

func (store *AddressDB) Modify(modifyPersonDetails *protocol.ModifyPerson) protocol.ServiceResponse {

	for _, personDet := range store.addressList {
		if personDet.Id == modifyPersonDetails.Id.Id {
			personDet.Name = modifyPersonDetails.ModifiedPerson.Name
			personDet.Id = modifyPersonDetails.ModifiedPerson.Id
			personDet.Phone = modifyPersonDetails.ModifiedPerson.Phone
			personDet.Address = modifyPersonDetails.ModifiedPerson.Address
			personDet.City = modifyPersonDetails.ModifiedPerson.City
			personDet.Pin = modifyPersonDetails.ModifiedPerson.Pin
		}
	}

	var serviceResponse protocol.ServiceResponse
	serviceResponse.IsSuccess = true
	serviceResponse.Error = ""

	return serviceResponse
}

func (store *AddressDB) Search(personId *protocol.PersonID) protocol.PersonList {

	var personList protocol.PersonList

	for _, personDet := range store.addressList {
		if personDet.Id == personId.Id {
			//personList.PersonsList = append(*(personList.PersonsList), personDet)
			break
		}
	}

	return personList
}

func (store *AddressDB) Delete(personId *protocol.PersonID) protocol.ServiceResponse {

	//var personList protocol.PersonList

	for _, personDet := range store.addressList {
		if personDet.Id == personId.Id {
			break
		}
	}

	var serviceResponse protocol.ServiceResponse
	serviceResponse.IsSuccess = true
	serviceResponse.Error = ""

	return serviceResponse
}
