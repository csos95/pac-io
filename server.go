package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	address string
}

func newServer(address string) *server {
	return &server{address: address}
}

func (s *server) run() error {
	router := mux.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("site")))

	srv := http.Server{
		Handler:      router,
		Addr:         s.address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
