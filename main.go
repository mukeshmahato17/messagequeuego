package main

import (
	"fmt"
	"log"
)

func main() {
	cfg := Config{
		ListenAddr: ":3000",
		Store:      NewMemoryStore(),
	}
	s, err := NewServer(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	s.Store.Push([]byte("foobar"))
	offset, err := s.Store.Fetch(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(offset)
}
