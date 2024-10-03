package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:         "/user",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		Uri:         "/user",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequireAuth: true,
	},
	{
		Uri:         "/user/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.GetUserById,
		RequireAuth: true,
	},
	{
		Uri:         "/user/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		Uri:         "/user/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		Uri:         "/user/{userId}/follow",
		Method:      http.MethodPost,
		Function:    controllers.UserFollow,
		RequireAuth: true,
	},
}
