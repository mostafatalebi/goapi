package main

import (
	"net/http"
)

type UserController struct {
}

func (ctrl *UserController) postAdd(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Posting the user add method."))
}
