package service

//ResponsePayload is the response payload struct for all API
type ResponsePayload struct {
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

func constructResponsePayload(data interface{}, err error) ResponsePayload {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	return ResponsePayload{
		Data:         data,
		ErrorMessage: errorMessage,
	}
}
