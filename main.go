package main

import (
	"tugas7/config"
	"tugas7/routers"
)

func main() {
	config.EnsureTableExists()
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}