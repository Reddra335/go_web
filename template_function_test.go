package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(Name string) string {
	return "Hello " + Name + " My Name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Rendi Damara",
	})
}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateFunction(recorder, req)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"Upper": func(value string) string {
			return strings.ToUpper(value)

		},
	})
	t = template.Must(t.Parse(`{{Upper .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Rendi Damara",
	})
}

func TestTemplateCreateGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateCreateGlobal(recorder, req)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateCreateGlobalPipline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"Upper": func(value string) string {
			return strings.ToUpper(value)

		},
	})
	t = template.Must(t.Parse(`{{sayHello .Name | Upper}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Rendi Damara",
	})
}

func TestTemplateCreateGlobalPipline(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateCreateGlobalPipline(recorder, req)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
