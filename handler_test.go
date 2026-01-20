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