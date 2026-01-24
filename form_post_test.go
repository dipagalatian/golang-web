package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// this func is a shortcut for the manual proses (parseform + get value)
	// r.PostFormValue("first_name")

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)	
}

func TestFormPost(t *testing.T) {

	bodyRequest := strings.NewReader("first_name=dipa&last_name=galatian")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/post-name", bodyRequest) 
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body,_  := io.ReadAll(response.Body)

	fmt.Println(string(body))
	
}