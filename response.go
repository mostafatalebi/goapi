package main

import "encoding/json"

type Response struct {
	status  bool
	message string
	data    string
}

type ResponseJSONStruct struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (r *Response) SetStatus(status bool) *Response {
	r.status = status
	return r
}

func (r *Response) SetMessage(msg string) *Response {
	r.message = msg
	return r
}

func (r *Response) SetData(data string) *Response {
	r.data = data
	return r
}

func (r *Response) ToJSONString() ([]byte, error) {
	response := ResponseJSONStruct{}
	response.Message = r.message
	response.Status = r.status
	response.Data = r.data
	return json.Marshal(response)
}
