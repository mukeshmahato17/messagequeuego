package main

type Config struct {
	ListenAddr string
	Store      Storage
}

type Server struct {
	*Config
}

func NewServer(cfg *Config) (*Server, error) {
	return &Server{
		Config: cfg,
	}, nil
}
