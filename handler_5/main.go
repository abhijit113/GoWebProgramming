/*

We have our usual hello handler function. We also have a log function, which takes
in a HandlerFunc and returns a HandlerFunc. Remember that hello is a HandlerFunc,
so this sends the hello function into the log function; in other words it chains the log
and the hello functions.
log(hello)
The log function returns an anonymous function that takes a ResponseWriter and a
pointer to a Request, which means that the anonymous function is a HandlerFunc.
Inside the anonymous function, we print out the name of the HandlerFunc (in this
Listing 3.10 Chaining two handler functions
http.HandlerFunc=return func(w http.ResponseWriter, r *http.Request)
Handlers and handler functions 61
case it’s main.hello), and then call it. As a result, we’ll get hello! in the browser and a
printed statement on the console that says this:
*/

package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path)
}
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello))

	//HandleFunc(pattern, HandlerFunc)
	server.ListenAndServe()

}
