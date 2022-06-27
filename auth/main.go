package main

import (
	"github.com/dheiro/adam-test/auth/config"
	"github.com/dheiro/adam-test/auth/routes"
)

func main() {
	db := config.InitDB()
	server := routes.InitRouter(db)
	server.Run()
}
