package go_web

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

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.html", nil)
}
func Upload(writer http.ResponseWriter, request *http.Request) {
	file, FileHeader, err := request.FormFile("file")
	ErrorT(err)
	fileDestination, err := os.Create("./recouse/" + FileHeader.Filename)
	ErrorT(err)
	_, err = io.Copy(fileDestination, file)
	ErrorT(err)

	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload_succes.html", map[string]any{
		"Name": name,
		"File": "/static/" + FileHeader.Filename,
	})
}
func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./recouse"))))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed her/4.png
var upload []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Rendi Damara")
	file, _ := writer.CreateFormFile("file", "contoh.jpg")
	file.Write(upload)
	writer.Close()

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	Upload(recorder, req)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
