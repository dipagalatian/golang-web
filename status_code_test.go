package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Name is empty")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=dipagalatian", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("status:", response.Status)
	fmt.Println(string(body))
	
}