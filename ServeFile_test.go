package go_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./recouse/valid.html")
	} else {
		http.ServeFile(writer, request, "./recouse/invalid.html")

	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed recouse/valid.html
var OkHtml string

//go:embed recouse/invalid.html
var NoHtml string

func ServeEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprintln(writer, OkHtml)
	} else {
		fmt.Fprintln(writer, NoHtml)

	}
}
func TestServeEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
