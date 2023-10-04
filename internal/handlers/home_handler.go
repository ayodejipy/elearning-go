package handlers

import "net/http"

func GetHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}