// main.go
package main

import (
	router2 "graphql_test/service/router"
)

func main() {
	r := router2.Router

	router2.SetRouter()

	r.Run(":1234")
}
