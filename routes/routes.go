package routes

import "api/controllers"

var Collection map[string]func

Collection["/users/add"] = users.UsersAdd