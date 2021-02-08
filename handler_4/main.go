/*
We talked about handlers, but what are handler functions? Handler functions are
functions that behave like handlers. Handler functions have the same signature as the
ServeHTTP method; that is, they accept a ResponseWriter and a pointer to a Request.
The following listing shows how this works with our server.

How does this work? Go has a function type named HandlerFunc, which will adapt a
function f with the appropriate signature into a Handler with a method f. For example,
take the hello function:
func hello(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello!")
}
If we do this:
helloHandler := HandlerFunc(hello)
then helloHandler becomes a Handler

*/

package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "at root")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

type ByeHandler struct{}

func (b *ByeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bye!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	helloHandler := http.HandlerFunc(hello) //creating a handler func and then passing it
	/*http.HandlerFunc(func(arg1 http.ResponseWriter, arg2 *http.Request) {

	})*/

	http.HandleFunc("/hola", helloHandler)
	/*http.HandleFunc(pattern, func(arg1 http.ResponseWriter, arg2 *http.Request) {

	})*/

	http.HandleFunc("/globe", world) //pass the handle directly
	/*http.HandleFunc(pattern, func(arg1 http.ResponseWriter, arg2 *http.Request) {

	})*/

	bye := ByeHandler{} // ByeHandler is a handler as it has implemented the ServeHTTP
	http.Handle("/adios/all/", &bye)
	//http.Handle(pattern, handler)

	http.HandleFunc("/", root)

	server.ListenAndServe()

	//http.ListenAndServe(":8080", helloHandler)
	//http.ListenAndServe(":8080", &bye)
	//http.ListenAndServe(":8080", world(w, r)) -not feasible
	//by doing this, it only execute first listenandserve

	//http.ListenAndServe(addr, handler)
}

/*
we registered the helloHandler to the URL /hello instead of /
hello/. For any registered URLs that don’t end with a slash (/), ServeMux will try to
match the exact URL pattern. If the URL ends with a slash (/), ServeMux will see if the
requested URL starts with any registered URL.
If we’d registered the URL /hello/ instead, then when /hello/there is requested, if
ServeMux can’t find an exact match, it’ll start looking for URLs that start with /hello/.
There’s a match, so helloHandler will be called.
*/
