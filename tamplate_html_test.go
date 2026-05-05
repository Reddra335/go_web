package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	simpleTamplate := `<html><body>{{.}}</body></html>`
	// r, err := template.New("SIMPLE").Parse(simpleTamplate)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(simpleTamplate))
	t.ExecuteTemplate(writer, "SIMPLE", "Hello Bro")

}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	SimpleHtml(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleGoHtml(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFiles("./template/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Blud")
}
func TestSimpleGoHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	SimpleGoHtml(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TamplateDirectory(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseGlob("./template/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Blud")
}
func TestDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	TamplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// func TamplateEmbed(writer http.ResponseWriter, request *http.Request) {

// 	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
// 	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Blud")
// }
// func TestEmbed(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
// 	recorder := httptest.NewRecorder()
// 	TamplateEmbed(recorder, request)

// 	body, _ := io.ReadAll(recorder.Result().Body)
// 	fmt.Println(string(body))
// }
