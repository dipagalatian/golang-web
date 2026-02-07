package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// CREATE TEMPLATE WITH IF ACTION
// see templates/if.gohtml for the template file

func CreateTemplateIfAction(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", map[string]any{
		"Title" : "If Action in Go Template",
		// "Name" : nil,
		"Name" : "Dipa Galatian",
	})
	
}

func TestCreateTemplateIfAction(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/create-template-if", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateIfAction(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

// CREATE TEMPLATE WITH COMPARATOR OPERATOR
// see templates/comparator.gohtml for the template file

func CreateTemplateOperator(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	t.ExecuteTemplate(w, "comparator.gohtml", map[string]any{
		"Title" : "Comparator Operator in Go Template",
		"FinalValue" : 50,
	})
	
}

func TestCreateTemplateOperator(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/create-template-operator", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateOperator(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

// CREATE TEMPLATE WITH RANGE ACTION
// see templates/range.gohtml for the template file
func CreateTemplateRangeAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]any{
		"Title" : "Range Action in Go Template",
		"Hobbies" : []string{
			"Gaming",
			"Reading",
			"Coding",
			"Traveling",
		},
	})
}

func TestCreateTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/create-template-range", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateRangeAction(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

// CREATE TEMPLATE WITH WITH ACTION
// see templates/address.gohtml for the template file
func CreateTemplateWithAction(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(w, "address.gohtml", map[string]any{
		"Title" : "With Action in Go Template",
		"Name" : "Dipa Galatian",
		"Address" : map[string]string{
			"Street" : "Jl. Merdeka No. 123",
			"City" : "Jakarta",
			"Country" : "Indonesia",
		},
	})
}

func TestCreateTemplateWithAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/with-template", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateWithAction(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}