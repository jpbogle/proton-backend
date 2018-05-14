package handlers

import (
	"proton/utils"
	"net/http"
)

var RootHandler = utils.CreateRootHandler()

func init() {
	RootHandler.Get("/", func(ctx *utils.Context) {
		ctx.Respond("SUP", http.StatusOK)
	})

	apiHandler := utils.CreateHandler()
	RootHandler.AddHandler("/api", apiHandler)

	apiHandler.Get("/demo", func(ctx *utils.Context) {
		ctx.Respond("<h1>Here's an api route demo!</h1>", http.StatusOK)
	})

	apiHandler.Get("/:name", func(ctx *utils.Context) {
		ctx.Respondf("<h1>Welcome to %s</h1>", http.StatusOK, ctx.Fields["name"])
	})

}
