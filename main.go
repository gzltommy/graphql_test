// main.go
package main

import (
	"graphql_test/router"
)

func main () {
	r := router.Router

	router.SetRouter()

	r.Run(":1234")
}