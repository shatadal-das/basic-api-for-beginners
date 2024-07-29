package router

import (
	"net/http"
)

func Net_Http(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handler1)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello World!"))
}
