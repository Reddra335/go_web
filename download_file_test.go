package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {

	ReqFile := request.URL.Query().Get("file")
	if ReqFile == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+ReqFile+"\"")
	http.ServeFile(writer, request, "./recouse/"+ReqFile)

}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	ErrorT(err)
}
