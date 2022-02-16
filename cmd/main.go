package main

import (
	"gifl/server"

	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
)

func main() {
	app := gear.New()
	app.UseHandler(logging.Default(true))

	router := gear.NewRouter()
	router.Get("/pi", server.Write)
	router.Get("/ip", server.Read)
	router.Otherwise(server.Other)

	app.UseHandler(router)

	logging.Info("server start...")
	if err := app.Listen(":9999"); err != nil {
		panic(err)
	}
}
