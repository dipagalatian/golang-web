package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandlerServer(t *testing.T) {

	// create handler for HTTP Request
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {

		// logic web -> response writer
		fmt.Fprint(w, "Hello Server Go!")
	}

	// create server
	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	// run server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}

func TestServerMux(t *testing.T) {

	// create server mux
	// to handle multiple route
	// like router in other framework
	// NewServeMux is to create multiple handlerFunc or handler
	// NewServeMux also implements the Handler interface like HandlerFunc
	mux := http.NewServeMux()

	// register route and handler func
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from root")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello from hi")
	})

	// endpoints with "/" at the end will match all subpaths
	// for example: /images/thumbnails, /images/anything
	// this is similar to using wildcard in other framework
	// the longest pattern match will be used not the first one defined
	// the order of definition does not matter in go net/http
	
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Images")
	})

	mux.HandleFunc("/images/thumnails/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumnails")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}
	
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}

func TestRequest(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.Host)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}