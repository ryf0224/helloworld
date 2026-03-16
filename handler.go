package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handleGreet(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if strings.EqualFold(msg, "hi") {
		fmt.Fprintln(w, "helloworld")
	} else {
		http.Error(w, "error: expected 'hi'", http.StatusBadRequest)
	}
}
