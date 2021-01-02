package model

type ErrorType int

const (
	WRONG_INPUTS ErrorType = iota
	ERROR_IN_SAVING
	SUCCESS
)

const (
	CODE_WRONG_INPUTS    string = "10101"
	CODE_ERROR_IN_SAVING string = "10102"
	CODE_SUCCESS         string = "10103"
)

const (
	ERR_MSG_WRONG_INPUTS string = "Wrong inputs json"
	ERR_MSG_IN_SAVING    string = "Error in DB layer"
	MSG_SUCCESS_SAVE     string = "Successfull Saved the address in DB"
	MSG_UNSUCCESS_SAVE   string = "Failed in saving the address in DB"
)

type PersonModel struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	City    string `json:"city"`
	Pin     string `json:"pin"`
}

type PersonModelArray struct {
	PersonRecords []PersonModel `json:"book"`
}

type ResponseModel struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
