package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Form_Post(write http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	FirstName := request.PostForm.Get("FirstName")
	LastName := request.PostForm.Get("LastName")
	fmt.Fprintf(write, "Hello %s %s", FirstName, LastName)
}

func TestFormatPost(t *testing.T) {
	requestBody := strings.NewReader("FirstName=Rendi&LastName=Damara")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	Form_Post(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
