package main

import (
	"WP_ApiDemo/apiV1"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	// initialize the server

	r := apiV1.InitServer()
	r.Run(":8080")
}
