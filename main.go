package main

import (
	"github.com/olusolaa/inventry/db"
	_ "github.com/olusolaa/inventry/docs"
	"github.com/olusolaa/inventry/server"
)

// @title        Inventory API
// @version      1.0
// @description  This is an inventory server. You can visit the GitHub repository at https://github.com/olusolaa/inventory

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host     inventory-api-v1.herokuapp.com
// @BasePath  /
// @securityDefinitions.basic  BasicAuth
func main() {

	DB := &db.PostgresDB{}
	DB.Init()

	s := &server.Server{
		DB: DB,
	}
	s.Start()
}
