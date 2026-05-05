package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.html"))
	t.ExecuteTemplate(writer, "layout", map[string]any{
		"Name": "Rendi",
	})
}

func TestTemplateLayout(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateLayout(recorder, req)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
