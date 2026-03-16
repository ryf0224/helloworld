package main

import (
	"fmt"
	"net/http"
)

func handleGreet(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "hi" {
		fmt.Fprintln(w, "helloworld")
	} else {
		http.Error(w, "error: expected 'hi'", http.StatusBadRequest)
	}
}
