package main

import (
	"fmt"
	"io"
	"net/http"

	"go-in-merry.com/merry"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Worldasdasd")
}
func getExample(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Get Example")
}
func postExample(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "POST Example")
}
func putExample(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "PUT Example")
}
func main() {
	merry.RegisterMiddleware(helloWorld, merry.WhiteListItem{Method: "GET", Endpoint: "/"})
	merry.Get("/", getExample)
	merry.Post("/", postExample)
	merry.Put("/", putExample)
	merry.Get("/contacts", getExample)
	merry.Sail(3000)

}
