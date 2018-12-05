package main

import (
	"net/http"
)

type UserController struct {
}

func (ctrl *UserController) postAdd(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	firstName := req.PostForm.Get("first_name")
	if len(firstName) > 0 {
		rda := Model{}
		res := rda.TTL(100).Save("user_"+firstName, firstName)
		if res {
			resp.Write([]byte("Error saving the user."))
		}
		resp.Write([]byte("User added successfully."))

	} else {
		resp.Write([]byte("Fields incomplete."))
	}
}
