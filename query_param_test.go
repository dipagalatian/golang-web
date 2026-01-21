package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request){

	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
	
}

func TestQueryParam(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=dipagalatian", nil)
	recorder := httptest.NewRecorder()
	
	SayHello(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(body)
	
}

func GetName(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("firstname")
	lastName := r.URL.Query().Get("lastname")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
} 

func TestMultipleQueryParam(t *testing.T) {
	
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/getname?firstname=dipa&lastname=galatian", nil)
	recorder := httptest.NewRecorder()
	
	GetName(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	
}

func GetMultipleNameValues(w http.ResponseWriter, r * http.Request) {
	var query = r.URL.Query()
	var names []string  = query["name"]

	fmt.Fprintf(w, "Hello %s", strings.Join(names, " "))
}

func TestMultipleQueryValues(t *testing.T) {
	
	request := httptest.NewRequest(http.MethodGet, "http://localhost/hello?name=dipa&name=galatian&name=aja", nil)
	recorder := httptest.NewRecorder()

	GetMultipleNameValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	
}