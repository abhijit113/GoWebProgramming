package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}
func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

/*
Instead of using HandleFunc to register the handler functions, we use the method
names:
mux.GET("/hello/:name", hello)
In this case we’re registering a URL for the GET method to the hello function. If we
send a GET request, the hello function will be called; if we send any other HTTP
Listing 3.12 Using HttpRouter
66 CHAPTER 3 Handling requests
requests it won’t. Notice that the URL now has something called a named parameter.
These named parameters can be replaced by any values and can be retrieved later by
the handler.
func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}
The handler function has changed too; instead of taking two parameters, we now take
a third, a Params type. Params contain the named parameters, which we can retrieve
using the ByName method
*/
