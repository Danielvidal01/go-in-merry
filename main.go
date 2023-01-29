package main

import (
	"fmt"
	"net/http"

	"going-marry.com/goingmerry"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func main() {
	goingmerry.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World From GET")
	})
	goingmerry.Post("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World From POST")
	})
	goingmerry.RegisterMiddleware(helloWorld, goingmerry.Route{Method: "GET", Endpoint: "/"})
	goingmerry.Ahoy(3000)

}
