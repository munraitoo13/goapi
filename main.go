package main

import (
	"log"
	"net/http"

	"github.com/munraitoo13/goapi/handlers"
	"github.com/munraitoo13/goapi/internal/middleware"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc // middlewares type
var middlewares = []Middleware{
	middleware.TokenAuthMiddleware, // auth middleware
}

const PORT = "6000" // server port

func main() {
	var handler http.HandlerFunc = handlers.HandleClientProfile // initial handler

	// loops through middlewares adding each one to the handler
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	http.HandleFunc("/client/profile", handler) // handlefunc for endpoints. path, handler

	log.Println("Server running on port:", PORT)  // logging
	log.Fatal(http.ListenAndServe(":"+PORT, nil)) // listen and serve starts the server
}
