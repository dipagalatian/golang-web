package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Successfully redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	// create any logic before redirecting
	fmt.Println("Redirecting to another route...")

	// redirect to another route or external url
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectToExternal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirecting to external url...")
	http.Redirect(w, r, "https://roadmap.sh/", http.StatusTemporaryRedirect)
}

func TestRedirectServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-external", RedirectToExternal)
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}