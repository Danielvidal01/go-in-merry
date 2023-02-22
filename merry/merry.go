package merry

import (
	"fmt"
	"log"
	"net/http"

	"go-in-merry.com/helpers"
)

type RequestHandeler = func(res http.ResponseWriter, req *http.Request)
type WhiteListItem struct {
	Endpoint string
	Method   string
}
type Middleware struct {
	handler   RequestHandeler
	WhiteList []WhiteListItem
}

type Executor = map[string]RequestHandeler

type Controller = map[string]Executor

var controllers = make(Controller)

var middlewares []Middleware

var serverIsRunning = false

// criar um group router **
func Sail(port int) {
	if serverIsRunning {
		log.Panic("SERVER ALREADY RUNNING")
	}
	parsedPort := fmt.Sprintf(":%v", port)
	iterateControllers()
	err := http.ListenAndServe(parsedPort, nil)
	fmt.Printf("⚡ server is running at http://localhost%s ⚡ \n", parsedPort)
	serverIsRunning = true
	if err != nil {
		fmt.Println(err)
	}
}

func HandleEndpoint(endpoint string, method string, handler RequestHandeler) {
	if serverIsRunning {
		log.Panic("SERVER ALREADY RUNNING")
	}
	if controllers[endpoint] == nil {
		executor := make(Executor)
		executor[method] = handler
		controllers[endpoint] = executor
		return
	}
	if controllers[endpoint][method] != nil {
		log.Panicf("at %s Method %s already in use", endpoint, method)
	}
	controllers[endpoint][method] = handler
}
func iterateControllers() {
	for endpoint := range controllers {
		execRoute(endpoint)
	}
}

func execRoute(endpoint string) {
	http.HandleFunc(endpoint, func(res http.ResponseWriter, req *http.Request) {
		method := req.Method
		if controllers[endpoint][method] == nil {
			res.WriteHeader(http.StatusMethodNotAllowed)
			res.Write([]byte("405 - method not allowed"))
			return
		}
		whiteListItem := WhiteListItem{Endpoint: endpoint, Method: method}
		execMiddlewares(whiteListItem, controllers[endpoint][method], res, req)
	})
}
func RegisterMiddleware(handler RequestHandeler, requestWhiteList ...WhiteListItem) {
	middleware := Middleware{handler: handler, WhiteList: requestWhiteList}
	middlewares = append(middlewares, middleware)
}

func execMiddlewares(whiteListItem WhiteListItem, handler RequestHandeler, res http.ResponseWriter, req *http.Request) {
	if len(middlewares) == 0 {
		handler(res, req)
	}
	for _, middleware := range middlewares {
		if helpers.Includes(middleware.WhiteList, whiteListItem) {
			handler(res, req)
			return
		}
		middleware.handler(res, req)
		handler(res, req)
	}
}
func Get(endpoint string, cb RequestHandeler) {
	HandleEndpoint(endpoint, "GET", cb)
}
func Post(endpoint string, cb RequestHandeler) {
	HandleEndpoint(endpoint, "POST", cb)
}
func Delete(endpoint string, cb RequestHandeler) {
	HandleEndpoint(endpoint, "DELETE", cb)
}
func Patch(endpoint string, cb RequestHandeler) {
	HandleEndpoint(endpoint, "PATCH", cb)
}
func Put(endpoint string, cb RequestHandeler) {
	HandleEndpoint(endpoint, "PUT", cb)
}
