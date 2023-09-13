package main

import (
	"crud/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")

}
