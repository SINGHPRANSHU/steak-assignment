package exception

import "net/http"


type RestError struct {
	Message interface{} `json:"message"`
	Status int `json:"status"`

}


func CreateBadRequestRestError(message string) RestError  {
	return RestError{
		Message: message,
		Status: http.StatusBadRequest,
	}
}