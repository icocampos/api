package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequered: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		AuthRequered: false,
	}, {
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		AuthRequered: false,
	}, {
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequered: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequered: false,
	},
}
