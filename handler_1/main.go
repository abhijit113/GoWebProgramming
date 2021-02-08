package main

import (
	"fmt"
	"net/http"
)

//We just created a handler
//and attached it to our server, so we’re no longer using any multiplexers. This
//means there’s no longer any URL matching to route the request to a particular handler,
//so all requests going into the server will go to this handler.
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(addr, handler)
}
