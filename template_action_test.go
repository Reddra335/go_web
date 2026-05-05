package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateIf(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.html"))
	t.ExecuteTemplate(write, "if.html", map[string]any{
		"Name": "Rendi",
	})
}

func TestTemplateIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateIfOperator(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if_operator.html"))
	t.ExecuteTemplate(write, "if_operator.html", map[string]any{
		"Name":  "Rendi",
		"Angka": 80,
	})
}

func TestTemplateIfOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateIfOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateRange(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.html"))
	t.ExecuteTemplate(write, "range.html", map[string]any{
		"Name": []string{"Rendi", "Master", "Reddra"},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateWith(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.html"))
	t.ExecuteTemplate(write, "with.html", map[string]any{
		"Name":   "Rendi",
		"Addres": map[string]any{},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
