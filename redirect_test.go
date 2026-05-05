package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func Redirect_to(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")

}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/redirect_to", http.StatusTemporaryRedirect)
}
func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "https://www.google.com/", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect_from", RedirectFrom)
	mux.HandleFunc("/redirect_to", Redirect_to)
	mux.HandleFunc("/redirect_out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
