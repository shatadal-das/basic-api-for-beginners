package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Gorilla_Mux(port string) {
	router := mux.NewRouter()

	router.HandleFunc("/hello", handler2).Methods("POST")

	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
