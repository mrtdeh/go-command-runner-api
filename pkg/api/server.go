package api_server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	Host   string
	Port   uint
	Router http.Handler
}

func (s *HttpServer) Serve() error {

	// Release Mode
	endPoint := fmt.Sprintf("%s:%d", s.Host, s.Port)

	server := &http.Server{
		Addr:        endPoint,
		Handler:     s.Router,
		IdleTimeout: time.Second * 10,
	}

	log.Println("listen on ", endPoint)

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("error in server listener : %w", err)
	}

	return nil
}
