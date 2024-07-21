package main

import (
	"fmt"
	"log"
)

func main() {
	cfg := &Config{
		ListenAddr: ":3000",
		StoreProducerFunc: func() Storer {
			return NewMemoryStore()
		},
	}
	s, err := NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	s.Store.Push([]byte("foobarbaz"))
	data, err := s.Store.Fetch(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
