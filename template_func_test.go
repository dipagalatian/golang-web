package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (p MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + p.Name
}


// CREATE TEMPLATE WITH FUNCTION
func CreateTemplateFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := `{{define "func"}} {{.SayHello "Gala"}} {{end}}`

	t := template.Must(template.New("func").Parse(tmpl))

	t.ExecuteTemplate(w, "func", MyPage{
		Name: "Dipa Galatian",
	})
}

func TestCreateTemplateFunc(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template-action", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateFunc(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// CREATE TEMPLATE WITH BUILT-IN GLOBAL FUNCTION
// built-in function: https://github.com/golang/go/blob/master/src/text/template/funcs.go
func CreateTemplateGlobalFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := `{{define "func"}} {{len .Name}} {{end}}`

	t := template.Must(template.New("func").Parse(tmpl))

	t.ExecuteTemplate(w, "func", MyPage{
		Name: "Dipa Galatian",
	})
}

func TestCreateTemplateGlobalFunc(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template-action", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateGlobalFunc(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// CREATE TEMPLATE WITH CUSTOM FUNCTION
// custom function: user-defined function that can be used in the template
// register custom function using Funcs method
// before parsing the template
func CreateTemplateCustomFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := `{{define "func"}} {{lower .Name}} {{end}}`

	t := template.New("func")
	t = t.Funcs(map[string]any {
		"upper" : func(value string) string {
			return strings.ToUpper(value)
		},
		"lower" : func(value string) string {
			return strings.ToLower(value)
		},
	})

	t = template.Must(t.Parse(tmpl))

	t.ExecuteTemplate(w, "func", MyPage{
		Name: "Dipa Galatian",
	})
}

func TestCreateTemplateCustomFunc(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template-action", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateCustomFunc(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// CREATE TEMPLATE WITH PIPE FUNCTION
// PIPE FUNCTION: the output of a function can be used as the input of another function
// Can have multiple functions in a pipe
func CreateTemplatePipeFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := `{{define "func pipe"}} {{sayHello .Name | upper}} {{end}}`

	t := template.New("func")

	t = t.Funcs(map[string]any {
		"upper" : func(value string) string {
			return strings.ToUpper(value)
		},
		"sayHello" : func(name string) string {
			return "Hello " + name
		},
	})

	t = template.Must(t.Parse(tmpl))

	t.ExecuteTemplate(w, "func pipe", MyPage{
		Name: "Dipa galatian",
	})
}

func TestCreateTemplatePipeFunc(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template-action", nil)
	recorder := httptest.NewRecorder()

	CreateTemplatePipeFunc(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}