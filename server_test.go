package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	sever := http.Server{
		Addr: "localhost:8080",
	}
	fmt.Println("http://localhost:8080/")

	err := sever.ListenAndServe()

	if err != nil {

		panic(err)
	}
}
