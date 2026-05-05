package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.html", map[string]any{
		"Title": "Master",
		"Body":  "<h1> Haii</h1>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)
	Body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(Body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)

	}
}

func TemplateAutoEscapeDisabled(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.html", map[string]any{
		"Title": "Master",
		"Body":  template.HTML("<h1> Haii</h1>"),
	})
}

func TestTemplateAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)
	Body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(Body))
}

func TemplateSimulasiTest(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.html", map[string]any{
		"Title": "Master",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateSimulasiXss(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateSimulasiTest),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)

	}
}
