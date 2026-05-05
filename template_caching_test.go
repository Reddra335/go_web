package go_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.html
var MyTemplate embed.FS

var myTemplates = template.Must(template.ParseFS(MyTemplate, "templates/*.html"))

func GoCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.html", "Hello")
}

func TestGoCaching(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	GoCaching(recorder, req)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
