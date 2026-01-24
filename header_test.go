package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {

	// headers are case insensitive, but in r.Header.Get is not case sensitive
	contentType := r.Header.Get("content-type")

	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/get-header", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {

	// to add custom header in response from server to client
	w.Header().Add("X-Powered-By", "dipa galatian")

	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/response-header", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	// to get custom header from response
	// the header is not case sensitive
	customHeader := recorder.Header().Get("x-powered-by")

	fmt.Println(customHeader)

	response := recorder.Result()
	body,_  := io.ReadAll(response.Body)

	fmt.Println(string(body))
	
}