package handlers

import "net/http"

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Adding user.."))
}