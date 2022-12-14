package main

import (
	"fmt"
	"net/http"
)

type App struct {
}

func (app App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	/* switch req.URL {
	case "/products":
		switch req.Method {
		case "GET":
		case "POST":
		}
	case "/customers":
		switch req.Method {
		case "GET":
		case "POST":
		}
	} */
	fmt.Printf("%s - %s\n", req.Method, req.URL)
	res.Write([]byte("Hello World!"))
}

func main() {
	app := App{}
	http.ListenAndServe(":8080", app)
}

/*
	GET - http://localhost:8080/products
	POST - http://localhost:8080/products

	GET - http://localhost:8080/customers
	POST - http://localhost:8080/customers
*/
