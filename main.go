package main

import (
	"go-architecture-mysql/api/config"
	"go-architecture-mysql/api/route"
)

func init() {
	config.SetupConfiguration()
}

// @title Go Architecture with MySQL
// @version 1.0
// @description Go Architecture
// @termsOfService http://swagger.io/terms/

// @contact.name Muhammad Rais Adlani
// @contact.url https://gitlab.com/mraisadlani
// @contact.email mraisadlani@gmail.com

// @license.name MIT
// @license.url https://gitlab.com/mraisadlani/go-architecture-mysql/-/blob/main/LICENSE

// @securityDefinitions.apikey BearerAuth
// @in Header
// @name Authorization

// @host localhost:9000
// @BasePath /api
func main() {
	db := config.InitDB()
	route.CreateHandler(db)
}