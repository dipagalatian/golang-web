package golangweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templates embed.FS

var templateCache = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func CreateTemplateWithCaching(w http.ResponseWriter, r *http.Request) {
	templateCache.ExecuteTemplate(w, "simple.gohtml", "Hello from template caching")
}

func TestCreateTemplateWithCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template-caching", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateWithCaching(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}