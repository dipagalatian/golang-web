package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// CREATE TEMPLATE FROM STRING

func CreateTemplateHTMLString(w http.ResponseWriter, r *http.Request) {
	templateString := `<html><body><h1>{{.}}</h1></body></html>`

	// Manual error handling
	// t, err := template.New("TEMPLATE 1").Parse(templateString)
	//  if err != nil {
	// 	panic(err)
	//  }

	// Using Must to handle error internally
	t := template.Must(template.New("TEMPLATE 1").Parse(templateString))

	 t.ExecuteTemplate(w, "TEMPLATE 1", "Hello World from template")
} 

func TestCreateTemplateHtmlString(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/create-template-s", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateHTMLString(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	expectedBody := `<html><body><h1>Hello World from template</h1></body></html>`

	if string(body) != expectedBody {
		t.Errorf("body response wrong, got %s want %s", string(body), expectedBody)
	}
	fmt.Println(string(body))
	
}

func TestCreateTemplateServer(t *testing.T) {

	server := http.Server{
		Addr: "localhost:8080", 
		Handler: http.HandlerFunc(CreateTemplateHTMLString),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}


// CREATE TEMPLATE FROM FILE

func CreateTemplateFromFile(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Body data for template file")
	
}

func TestCreateTemplateFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/create-template-f", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateFromFile(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// CREATE TEMPLATE MULTIPLE FILES FROM DIRECTORY

func CreateTemplateFromDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Body data for template dir")
}

func TestCreateTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/create-template-d", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateFromDirectory(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// CREATE TEMPLATE FROM DIR USING GO EMBED

// embed declaration is moved to template_caching_test.go

func CreateTemplateFromDirWithEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Body data for template embed")
}

func TestCreateTemplateDirWithEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/create-template-e", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateFromDirWithEmbed(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
