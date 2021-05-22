package store

import (
	"database/sql"
	"gta4roy/address_service/log"
	"gta4roy/address_service/protocol"

	_ "github.com/go-sql-driver/mysql"
)

type AddressDB struct {
	addressList []protocol.Person
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "gta4roy"
	dbPass := "71201"
	dbName := "address"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	log.Trace.Println("Service got connected to SQL Server")
	return db
}

func (store *AddressDB) Add(personDetails *protocol.Person) protocol.ServiceResponse {

	log.Trace.Println("Add Service Request ")
	db := dbConn()
	defer db.Close()

	var serviceResponse protocol.ServiceResponse
	serviceResponse.IsSuccess = true
	serviceResponse.Error = ""

	insForm, err := db.Prepare("INSERT INTO addressbook(id,name,address,phone,city,pin) VALUES(?,?,?,?,?,?)")
	insForm.Exec(personDetails.Id, personDetails.Name, personDetails.Address, personDetails.Phone, personDetails.City, personDetails.Pin)
	if err != nil {
		serviceResponse.IsSuccess = false
		serviceResponse.Error = err.Error()
		log.Trace.Println("Failed to execute insert request", err.Error())
	}
	return serviceResponse
}

func (store *AddressDB) ListAll() protocol.PersonList {

	var personList protocol.PersonList
	db := dbConn()
	defer db.Close()
	selDB, err := db.Query("SELECT * FROM addressbook ORDER BY id DESC")
	if err != nil {
		log.Trace.Println("Failed to execute select all request", err.Error())
	}
	var personArray []*protocol.Person
	for selDB.Next() {
		var personObj protocol.Person
		err = selDB.Scan(&personObj.Id, &personObj.Name, &personObj.Address, &personObj.Phone, &personObj.City, &personObj.Pin)
		if err != nil {
			log.Trace.Println("Failed to execute Scan  all request", err.Error())
		}

		personArray = append(personArray, &personObj)
	}

	personList.PersonsList = personArray
	return personList
}

func (store *AddressDB) Modify(modifyPersonDetails *protocol.ModifyPerson) protocol.ServiceResponse {

	db := dbConn()
	log.Trace.Println("Update Request")
	insForm, err := db.Prepare("UPDATE addressbook SET name=?,phone=?,address=?,city=?,pin=? WHERE id=?")
	if err != nil {
		log.Trace.Println("Failed to update request", err.Error())
	}
	insForm.Exec(modifyPersonDetails.ModifiedPerson.Name, modifyPersonDetails.ModifiedPerson.Phone, modifyPersonDetails.ModifiedPerson.Address, modifyPersonDetails.ModifiedPerson.City, modifyPersonDetails.ModifiedPerson.Pin, modifyPersonDetails.ModifiedPerson.Id)
	defer db.Close()
	var serviceResponse protocol.ServiceResponse
	serviceResponse.IsSuccess = true
	serviceResponse.Error = ""

	return serviceResponse
}

func (store *AddressDB) Search(personId *protocol.PersonID) protocol.PersonList {

	log.Trace.Println("Search Request")
	var personList protocol.PersonList
	db := dbConn()
	defer db.Close()
	selDB, err := db.Query("SELECT * FROM addressbook WHERE id =?", personId.Id)
	if err != nil {
		log.Trace.Println("Failed to search request", err.Error())
	}

	var personArray []*protocol.Person

	for selDB.Next() {

		var personObj protocol.Person
		err = selDB.Scan(&personObj.Id, &personObj.Name, &personObj.Address, &personObj.Phone, &personObj.City, &personObj.Pin)
		if err != nil {
			panic(err.Error())
		}

		personArray = append(personArray, &personObj)
	}

	personList.PersonsList = personArray
	return personList
}

func (store *AddressDB) Delete(personId *protocol.PersonID) protocol.ServiceResponse {

	log.Trace.Println("Delete Request")
	db := dbConn()
	delForm, err := db.Prepare("DELETE FROM addressbook WHERE id=?")
	if err != nil {
		log.Trace.Println("fail to Delete Request", err.Error())
	}
	delForm.Exec(personId.Id)
	defer db.Close()

	var serviceResponse protocol.ServiceResponse
	serviceResponse.IsSuccess = true
	serviceResponse.Error = ""

	return serviceResponse
}
