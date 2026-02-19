package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (m *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Log execute before handler...")
	m.Handler.ServeHTTP(w, r)
	fmt.Println("Log execute after handler...")
}

type ErrorHandlerMiddlerware struct {
	Handler http.Handler
}

func (m *ErrorHandlerMiddlerware) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error handled in middleware: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error, due to: %s", err)
		}
	}()
	
	m.Handler.ServeHTTP(w, r)
}

func TestMiddlewareServer(t *testing.T) {
	mux := http.NewServeMux()

	// All handlers will be wrapped with LogMiddleware, so the log will be executed before and after the handler

	// MIDDLEWARE FOR LOGGING
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler executed successfully")
		fmt.Fprint(w, "Testing middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Foo executeed successfully")
		fmt.Fprint(w, "Hello Foo")
	})

	// MIDDLEWARE FOR ERROR HANDLING
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Panic executed successfully")
		panic("Oops something went wrong")
	})

	// Using struct to wrap the handler with middleware

	// First middleware to logging
	logMiddleware := &LogMiddleware{Handler: mux}

	// Second middleware to error handling, wrap the log middleware
	errorHandlerMiddleware := &ErrorHandlerMiddlerware{Handler: logMiddleware}

	server := http.Server{
		Addr: "localhost:8080",
		// Alternative without using struct
		// Handler: &LogMiddleware{Handler: mux},

		// Only use the Log middleware
		// Handler: logMiddleware,

		// Use both Log and Error handling middleware
		Handler: errorHandlerMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}