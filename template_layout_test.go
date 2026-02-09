package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// CREATE TEMPLATE WITH LAYOUT ACTION
// see templates/layouts/base.gohtml and templates/layouts/home.gohtml for the template files
func CreateTemplateLayoutAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/layout.gohtml", "./templates/header.gohtml", "./templates/footer.gohtml"))
	t.ExecuteTemplate(w, "layout", map[string]any{
		"Title" : "Layout Action in Go Template",
		"Name" : "Dipa Galatian",

	})
	// Different layout execution in the same file on templates/layout.gohtml
	// t.ExecuteTemplate(w, "test", map[string]any{
	// 	"Title" : "Layout Action in Go Template",
	// 	"Name" : "Dipa Galatian",
	// })
}

func TestCreateTemplateLayoutAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/layout-template", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateLayoutAction(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}