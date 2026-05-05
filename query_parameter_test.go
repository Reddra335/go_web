package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(write http.ResponseWriter, request *http.Request) {

	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(write, "Hello")

	} else {
		fmt.Fprintf(write, "Hello %s", name)
	}

}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Reddra", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleQueryParameter(write http.ResponseWriter, request *http.Request) {
	FirstName := request.URL.Query().Get("FirstName")
	SecondName := request.URL.Query().Get("SecondName")

	fmt.Fprintf(write, "Hello %s %s", FirstName, SecondName)

}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/uji?FirstName=Reddra&SecondName=Master", nil)

	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}
func TestMultipleParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/uji?name=Rendi&name=Damara", nil)

	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
