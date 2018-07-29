package main

import (
	"github.com/salapao2136/example-api/src/account"
	"github.com/salapao2136/example-api/src/config"

	"github.com/salapao2136/example-api/src/users"

	"github.com/labstack/echo"
)

func init() {

}
func main() {
	db := config.Init()
	e := echo.New()
	e.Use(ServerHeader)
	api := e.Group("/api") // prefix /api
	users.RouterUsers(api, db)
	account.RouterAccount(api, db)

	e.Logger.Fatal(e.Start(":8005"))
}
