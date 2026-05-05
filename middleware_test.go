package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Sebelum Middleware di execute")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("Sesudah Middleware di executed")

}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(writer, request)

}
func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprint(writer, "Hello Middleware")

	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Foo")
		fmt.Fprint(writer, "Hello Foo")

	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello bro")
		panic("jir")
	})
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	errorMiddleware := &ErrorHandler{
		Handler: logMiddleware,
	}

	var server = http.Server{
		Addr:    "localhost:8080",
		Handler: errorMiddleware,
	}
	err := server.ListenAndServe()
	ErrorT(err)
}
