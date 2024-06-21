package main

import (
	"Todo-List/routes"
	"log"
	"net/http"
)

const (
	port string = "4040"
)

func main() {

	if err := http.ListenAndServe(":"+port, routes.Init()); err != nil {
		log.Fatal(err)
	}
}
