/*
Most of the time, we don’t want to have a single handler to handle all the requests like
in listing 3.6; instead we want to use different handlers instead for different URLs. To
Listing 3.6 Handling requests
Handlers and handler functions 57
do this, we don’t specify the Handler field in the Server struct like handler_2 (which means it will
use the DefaultServeMux as the handler); we use the http.Handle function to attach
a handler to DefaultServeMux. Notice that some of the functions like Handle are
functions for the http package and also methods for ServeMux. These functions are
actually convenience functions; calling them simply calls DefaultServeMux’s corresponding
functions. If you call http.Handle you’re actually calling DefaultServeMux’s
Handle method.
*/
package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}
func main() {
	hello := HelloHandler{} //creating a handler which you don't need if you use handlefunc in example 4
	world := WorldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
