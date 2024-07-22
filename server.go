package main

import "net/http"

type Config struct {
	ListenAddr        string
	StoreProducerFunc StoreProducerFunc
}

type Server struct {
	*Config
	topics map[string]Storer
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
		topics: make(map[string]Storer),
	}, nil
}

func (s *Server) Start() {
	http.ListenAndServe(s.ListenAddr, nil)
}

func (s *Server) createTopic(name string) {
	_, ok := s.topics[name]
	if !ok {
		s.topics[name] = s.StoreProducerFunc()
	}
}
