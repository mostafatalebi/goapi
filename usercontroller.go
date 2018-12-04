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
		rda := RedisAdapter{}
		err := rda.Set("user_"+firstName, firstName, 0)
		if err != nil {
			resp.Write([]byte("Error saving the user."))
		} else {
			resp.Write([]byte("User added successfully."))
		}
		
	} else {
		resp.Write([]byte("Fields incomplete."))
	}
}
