package main

import (
	"net/http"
)

type UserController struct {
}

func (ctrl *UserController) postAdd(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.PostForm.Get("username")
	password := req.PostForm.Get("password")
	passwordConfirm := req.PostForm.Get("password_confirm")
	response := Response{}
	UVal := UserValidation{}
	UVal.Validator.Required("username", "password", "password_confirm")

	if UVal.Validator.HasRequiredFields(&req.PostForm) == false {
		response.SetMessage(UVal.Validator.LastError())
		resp.WriteHeader(HTTP_BAD_REQUEST)
		respJSON, _ := response.ToJSONString()
		resp.Write(respJSON)
		return
	}

	UVal.Min(4).Max(32).UserName(username)
	UVal.Min(6).Max(32).Password(password)
	UVal.Compare(password, passwordConfirm)

	if UVal.Succeed() {
		rda := Model{}
		rda.SetEngine(engineRedis)
		res := rda.TTL(100).Save("user_"+username, username)
		if res {
			response.SetStatus(true)
			response.SetMessage("User added successfully")
			response.SetData("")
			respJSON, _ := response.ToJSONString()
			resp.Write(respJSON)
			return
		}
		response.SetStatus(false)
		response.SetMessage("Failed to insert user into storage")
		response.SetData("")
		respJSON, _ := response.ToJSONString()
		resp.Write(respJSON)
		return

	} else {
		response.SetMessage(UVal.Validator.LastError())
		response.SetStatus(false)
		respJSON, _ := response.ToJSONString()
		resp.Write(respJSON)
	}
}
