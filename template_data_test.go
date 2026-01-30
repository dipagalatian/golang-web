package golangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// CREATE TEMPLATE WITH DYNAMIC DATA USING MAP
// see templates/name.gohtml for the template file
func CreateTemplateDynamicDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", map[string]any{
		"Title" : "Dynamic Title from Go",
		"Name" : "Dipa Galatian",
	})
}

func TestCreateTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template-data", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateDynamicDataMap(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// CREATE TEMPLATE WITH DYNAMIC DATA USING STRUCT
type Address struct {
	Street string
}
type Page struct {
	Title string
	Name string
	Address Address
}


func CreateTemplateDynamicDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Title data from struch",
		Name:  "Dipa Galatian",
		Address: Address{
			Street: "Jalan Merdeka No.123",
		},
	})
}

func TestCreateTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/template-data", nil)
	recorder := httptest.NewRecorder()

	CreateTemplateDynamicDataStruct(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestCreateTemplateDataServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(CreateTemplateDynamicDataStruct),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}