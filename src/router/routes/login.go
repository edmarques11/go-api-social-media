package routes

import (
	"api/src/controllers"
	"net/http"
)

var login = Route{
	Uri:         "/login",
	Method:      http.MethodPost,
	Function:    controllers.Login,
	RequireAuth: false,
}
