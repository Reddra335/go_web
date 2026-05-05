package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Konten-Type")

	fmt.Fprintln(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)

	request.Header.Add("Konten-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()
	Body, _ := io.ReadAll(response.Body)

	fmt.Println(string(Body))

}

func HeaderResponse(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Add("Reddra_By", "Rendi Damara")
	fmt.Fprintln(writer, "Mantap")

}

func TestHeaderResponse(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)

	request.Header.Add("Konten-Type", "application/json")
	recorder := httptest.NewRecorder()

	HeaderResponse(recorder, request)
	response := recorder.Result()
	Body, _ := io.ReadAll(response.Body)

	fmt.Println(string(Body))
	fmt.Println(response.Header.Get("Reddra_By"))

}
