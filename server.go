package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run() {
	http.HandleFunc("/", handleGreet)
	fmt.Printf("Server listening on %s\n", s.addr)
	http.ListenAndServe(s.addr, nil)
}
