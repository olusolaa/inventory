package main

import (
	"github.com/olusolaa/inventry/cache"
	"github.com/olusolaa/inventry/db"
	_ "github.com/olusolaa/inventry/docs"
	"github.com/olusolaa/inventry/server"
)

// @title        Inventory API
// @version      1.0
// @description  This is a movie server. You can visit the GitHub repository at https://github.com/olusola/inventory

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.basic  BasicAuth
func main() {

	var moviesCache = cache.NewRedisCache("localhost:6379", 1, "", 100)

	DB := &db.PostgresDB{}
	DB.Init()

	s := &server.Server{
		DB:    DB,
		Cache: moviesCache,
	}
	s.Start()
}
