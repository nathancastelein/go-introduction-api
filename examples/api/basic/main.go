package main

import (
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"hello": "world"}`))
}

func main() {
	http.HandleFunc("GET /hello", HelloWorld)

	http.ListenAndServe(":8080", nil)
}
