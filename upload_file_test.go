package golangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	templateCache.ExecuteTemplate(w, "upload.form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {

	// HANDLING THE UPLOADED FILE
	
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	// create destination file
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// copy file to destination
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	// HANDLING OTHER FORM DATA
	name := r.PostFormValue("name")

	// RENDER DATA TO TEMPLATE
	templateCache.ExecuteTemplate(w, "upload.success.gohtml", map[string]any {
		"Name" : name,
		"File" : "/static/" + fileHeader.Filename,
	})
	
}


func TestUploadFileServer(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/upload-image.jpg
var uploadFileTest []byte

func TestUploadFileMock(t *testing.T) {
	// create multipart form data body
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "dipagalatian")
	file, _ := writer.CreateFormFile("file", "contohcontent.jpg")
	// use the embedded file content as the uploaded file content
	file.Write(uploadFileTest)
	writer.Close() 

	
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload-file", body)
	// set the content type header to multipart form data with the boundary from the writer
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
 
	bodyResponse, _ := io.ReadAll(response.Body)

	fmt.Println(string(bodyResponse))
}