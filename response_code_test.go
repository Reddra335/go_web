package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(write http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		write.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(write, "Namanya Kosong")
	} else {
		fmt.Fprintf(write, "Hello %s", name)
	}

}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}
func TestResponseCodevalid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Reddra", nil)
	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))

}
