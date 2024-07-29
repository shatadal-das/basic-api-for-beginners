package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Chi(port string) {
	r := chi.NewRouter()

	r.Post("/hello", handler4)

	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}

func handler4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
