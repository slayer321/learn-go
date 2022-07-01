package main

import (
	"finalassignment/database"
	"finalassignment/routes"
	"fmt"
)

func main() {

	fmt.Println("Strating with MongoDB connection")

	database.Init()

	fmt.Println("Server is getting started")

	r := routes.Router()

	r.Run(":5000")
	fmt.Println("Listening at port 5000")
}
