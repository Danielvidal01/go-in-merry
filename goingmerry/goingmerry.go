package goingmerry

import (
	"fmt"
	"net/http"
)

type Route struct {
	Method   string
	Endpoint string
}
type RequestHandeler = func(res http.ResponseWriter, req *http.Request)
type Middleware struct {
	Exec      RequestHandeler
	WhiteList []Route
}

var routes []Route
var middlewares []Middleware

func Ahoy(port int) {
	parsedPort := fmt.Sprintf(":%v", port)
	fmt.Printf("⚡ server is running at http://localhost%s ⚡", parsedPort)
	err := http.ListenAndServe(parsedPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func RegisterMiddleware(exec RequestHandeler, requestWhiteList ...Route) {
	middleware := Middleware{Exec: exec, WhiteList: requestWhiteList}
	middlewares = append(middlewares, middleware)
}

func registerRoute(r Route) {
	fmt.Printf("ROUTE REGISTERED %v \n", r)
	routes = append(routes, r)
}
func includes(values []Route, value Route) bool {
	for _, val := range values {
		if val == value {
			return true
		}
	}

	return false
}

func execMiddlewares(route Route, cb RequestHandeler, res http.ResponseWriter, req *http.Request) {
	if len(middlewares) == 0 {
		cb(res, req)
	}
	for _, middleware := range middlewares {
		if includes(middleware.WhiteList, route) {
			cb(res, req)
			return
		}
		middleware.Exec(res, req)
		cb(res, req)
	}
}
func execMethod(route Route, res http.ResponseWriter, req *http.Request, cb RequestHandeler) {
	if route.Method != req.Method {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte("405 - method not allowed"))
		return
	}

	execMiddlewares(route, cb, res, req)
}

func Get(endpoint string, cb RequestHandeler) {
	route := Route{Method: "GET", Endpoint: endpoint}
	registerRoute(route)
	http.HandleFunc(endpoint, func(res http.ResponseWriter, req *http.Request) {
		execMethod(route, res, req, cb)
	})
}
func Post(endpoint string, cb RequestHandeler) {
	route := Route{Method: "POST", Endpoint: endpoint}
	registerRoute(route)
	http.HandleFunc(route.Endpoint, func(res http.ResponseWriter, req *http.Request) {
		execMethod(route, res, req, cb)
	})
}
func Delete(endpoint string, cb RequestHandeler) {
	route := Route{Method: "DELETE", Endpoint: endpoint}
	registerRoute(route)
	http.HandleFunc(route.Endpoint, func(res http.ResponseWriter, req *http.Request) {
		execMethod(route, res, req, cb)
	})
}

func Patch(endpoint string, cb RequestHandeler) {
	route := Route{Method: "PATCH", Endpoint: endpoint}
	registerRoute(route)
	http.HandleFunc(route.Endpoint, func(res http.ResponseWriter, req *http.Request) {
		execMethod(route, res, req, cb)
	})
}

func Put(endpoint string, cb RequestHandeler) {
	route := Route{Method: "PUT", Endpoint: endpoint}
	registerRoute(route)
	http.HandleFunc(route.Endpoint, func(res http.ResponseWriter, req *http.Request) {
		execMethod(route, res, req, cb)
	})
}
