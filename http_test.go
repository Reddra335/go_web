package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var host1 = "http://localhost:8080/"

func HandlerHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello Guys")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	HandlerHello(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
