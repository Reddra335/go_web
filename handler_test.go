package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(write http.ResponseWriter, request *http.Request) {
		//Logic Web
		fmt.Fprintln(write, `Lorem ipsum, dolor sit amet consectetur adipisicing elit. Illo ex consequuntur distinctio quas totam itaque! Iure, dolorum nam fugiat eligendi omnis veritatis reiciendis asperiores! Voluptatem, culpa illum! Sed, et recusandae. Saepe dolor odit maiores
error, in quas quo nulla sint deleniti commodi, aut corporis sit accusantium tenetur nesciunt? Iure perspiciatis, necessitatibus facere temporibus aliquam possimus et fugiat repudiandae. Facilis, maiores? Distinctio aliquam officia aspernatur, doloribus
iste quidem placeat, accusantium est aliquid rem eveniet exercitationem. Natus aspernatur quasi, magnam animi et a fugiat? Cupiditate corrupti iure, delectus et aut doloremque! Temporibus! Facere, eligendi eveniet voluptatum at blanditiis iusto dolores
culpa nihil sequi fugiat suscipit totam eum qui possimus sed dolorem corporis, quisquam atque magni maiores, in repellat provident quae eos. Magni? Ullam earum, nam consequuntur necessitatibus facere id, cupiditate dolorum unde iusto alias voluptate libero
facilis hic? Placeat ad, sed animi dolor ex expedita. Quam neque totam laudantium mollitia quasi nisi! Nemo tempore distinctio id voluptatum nihil expedita quasi repudiandae sit reiciendis. In quas, illo facilis, cumque porro, autem nisi mollitia nihil
nulla tenetur beatae laborum vitae quibusdam totam ipsa. Ducimus. A neque doloribus fuga aperiam architecto deleniti mollitia omnis sint, nihil ad at, quisquam beatae provident totam reprehenderit dolorem. Reprehenderit dolor quidem explicabo laborum
tempora in labore pariatur ipsam aut. Maiores odio consectetur iste quas cum nihil quaerat a, ab quis fugiat officiis ea aliquam ratione ut iusto sed incidunt voluptas reiciendis! Dolor magnam inventore exercitationem porro animi quas cumque? Rerum at
dolores saepe, nulla accusamus cupiditate! Dolorem repudiandae suscipit ab maxime sed! Dolorem, repudiandae. Veniam maxime earum mollitia quidem officiis vero, temporibus consequatur velit dolorum aspernatur, est, quam deserunt. Iste soluta hic accusamus
iusto autem dolor deleniti inventore eveniet sit nostrum excepturi nisi, incidunt sed nesciunt voluptas voluptatibus quibusdam, doloribus laborum? Labore sapiente neque perferendis quasi recusandae quas eum?`)

	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestHandlerMuc(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(writer, "Haiiii")

	})
	mux.HandleFunc("/beranda", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(writer, "Hello Asep")

	})
	mux.HandleFunc("/image/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(writer, "Hello Ayu")

	})
	mux.HandleFunc("/image/thumbnail", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(writer, "Reddra")

	})

	server := http.Server{
		Addr:    "http://localhost:8080/",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)

	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "http://localhost:8080/",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
