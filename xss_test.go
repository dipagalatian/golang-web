package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CreateAutoEscapeXSS(w http.ResponseWriter, r *http.Request) {
	templateCache.ExecuteTemplate(w, "xss.gohtml", map[string]any{
		"Title":   "Auto Escape XSS in Go Template",
		"Content": "<script>alert('XSS Attack!');</script>",
		// "Content" : "<p>This is a paragraph.</p>",
	})
}

func TestCreateAutoEscapeXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/auto-escape-xss", nil)
	recorder := httptest.NewRecorder()

	CreateAutoEscapeXSS(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))	
}

func TestAutoEscapeServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/auto-escape-xss", CreateAutoEscapeXSS)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// DISABLE AUTO ESCAPE XSS USING template.HTML TYPE
// WARNING: Disabling auto-escaping can expose your application to XSS attacks.
// Template.HTML, template.JS and template.CSS will bypass the auto-escaping mechanism
// Do this only if you are sure that the content is safe and trusted
// to prevent XSS attacks.
// In real-world applications, always sanitize and validate any user-generated content
// before rendering it in templates.

func CreateAutoEscapeXSSDisabled(w http.ResponseWriter, r *http.Request) {
	templateCache.ExecuteTemplate(w, "xss.gohtml", map[string]any{
		"Title":   "Auto Escape XSS in Go Template",
		// "Content": "<script>alert('XSS Attack!');</script>",
		"Content" : template.HTML("<p>This is a paragraph.<script>alert('XSS Attack!');</script></p>"),
	})
}

func TestCreateAutoEscapeXSSDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/auto-escape-xss-disabled", nil)
	recorder := httptest.NewRecorder()

	CreateAutoEscapeXSSDisabled(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))	
}

func TestAutoEscapeXSSDisabledServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(CreateAutoEscapeXSSDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}