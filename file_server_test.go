package golangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	// This will not work as expected
	// because the file server serves files relative to the root "/" -> server will try to find "/resources/static/filename" not "resources/filename"
	// so we need to strip the prefix first
	// mux.Handle("/static/", fileServer)

	// Correct way
	// Strip the "/static" prefix before passing the request to the file server
	// So when a request comes to "/static/filename", it will be stripped to "/filename" before reaching the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}

//go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {

	directory, _ := fs.Sub(resources, "resources")

	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}