package handlers

import (
	// "encoding/json"
	"net/http"
	"proton/controllers"
	"proton/entities"
	"proton/utils"
)

func registerUser(apiHandler *utils.Handler) {
	// Fetch all users
	apiHandler.Get("/users", func(ctx *utils.Context) {
		users, err := controllers.GetUsers()
		if notFoundErr, ok := err.(*entities.NotFoundError); ok {
			ctx.Error(notFoundErr, http.StatusOK)
			return
		}
		if err != nil {
			ctx.Error(err, http.StatusInternalServerError)
			return
		}

		resp := entities.GetUsersResponse{
			Users: users,
		}
		ctx.RespondJson(resp, http.StatusOK)
	})

	// Fetch a user
	apiHandler.Get("/users/:username", func(ctx *utils.Context) {
		user, err := controllers.GetUser(ctx.Fields["username"])
		if err != nil {
			ctx.Error(err, http.StatusInternalServerError)
			return
		}

		resp := entities.GetUserResponse{
			User: user,
		}
		ctx.RespondJson(resp, http.StatusOK)
	})
}
