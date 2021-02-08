package main

import (
	"fmt"
	"net/http"
)

//we are using our own handler instead of the default multiplexer, we’ll be
//able to respond, as shown in this listing.

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}

// If you go to http://localhost:8080 in your browser you’ll see Hello World!
//Here’s the tricky bit: if you go to http://localhost:8080/anything/at/all you’ll still
//get the same response! Why this is so should be quite obvious. We just created a handler
//and attached it to our server, so we’re no longer using any multiplexers. This
//means there’s no longer any URL matching to route the request to a particular handler,
// so all requests going into the server will go to this handler

//This is the reason why we’d normally use a multiplexer. Most of the time we want
//the server to respond to more than one request, depending on the request URL. Naturally
//if you’re writing a very specialized server for a very specialized purpose, simply
//creating one handler will do the job marvelously.
