package api

import (
	"AddressService/log"
	"AddressService/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var serviceProxy ServiceProxy

func init() {
	serviceProxy.initialiseConnection()
}

func CloseConnections() {
	serviceProxy.CloseClientConnect()
}

func handleGetHealth(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("Health Request")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"UP"}`))
	return
}

func handleAddAddress(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handleAddAddress Request")

	personDetails, err := parseRequestParams(w, r)
	if err != model.SUCCESS {
		log.Trace.Println("Error in parsing the request")
	}

	log.Trace.Printf("%s %s %s %s", personDetails.Name, personDetails.Id, personDetails.Phone, personDetails.Address)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var responseData model.ResponseModel
	responseData = serviceProxy.AddAddress(personDetails)
	json.NewEncoder(w).Encode(responseData)

}

func handleModifyAddress(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handleModifyAddress Request")

	personDetails, err := parseRequestParams(w, r)
	if err != model.SUCCESS {
		log.Trace.Println("Error in parsing the request")
	}

	log.Trace.Printf("%s %s %s %s", personDetails.Name, personDetails.Id, personDetails.Phone, personDetails.Address)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var responseData model.ResponseModel
	responseData = serviceProxy.ModifyAddress(personDetails.Id, personDetails)
	json.NewEncoder(w).Encode(responseData)

}
func handleSearchAddress(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handleModifyAddress Request")
	personId := r.FormValue("id")
	log.Trace.Printf("%s", personId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var personList model.PersonModelArray
	personList.PersonRecords = serviceProxy.SearchAddress(personId)

	json.NewEncoder(w).Encode(personList)

}
func handlePrintAllAddress(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handlePrintAllAddress Request")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var personList model.PersonModelArray
	personList.PersonRecords = serviceProxy.ListOfAddress()

	json.NewEncoder(w).Encode(personList)
}

func handleDeleteAddress(w http.ResponseWriter, r *http.Request) {
	log.Trace.Println("handleDeleteAddress Request")
	personId := r.FormValue("id")
	log.Trace.Printf("%s", personId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	responseData := serviceProxy.DeleteAddress(personId)
	json.NewEncoder(w).Encode(responseData)
}

func parseRequestParams(w http.ResponseWriter, r *http.Request) (model.PersonModel, model.ErrorType) {
	log.Trace.Println("parsing the input parameters ")
	var personDataRequestSet model.PersonModel

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil || reqBody == nil {
		log.Error.Println("Error Reading the body of the request %v", err)
		return personDataRequestSet, model.WRONG_INPUTS
	}

	err = json.Unmarshal(reqBody, &personDataRequestSet)
	if err != nil {
		log.Error.Println("Error Reading the body of the request %v", err)
		return personDataRequestSet, model.WRONG_INPUTS
	}

	return personDataRequestSet, model.SUCCESS
}
