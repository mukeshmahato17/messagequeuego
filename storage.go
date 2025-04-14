package main

import "fmt"

type Storage interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

type MemoryStore struct {
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (m *MemoryStore) Push(b []byte) (int, error) {
	m.data = append(m.data, b)
	return len(m.data) - 1, nil
}

func (m *MemoryStore) Fetch(offset int) ([]byte, error) {
	if len(m.data) < offset {
		return nil, fmt.Errorf("offset (%d) is too high", offset)
	}

	return m.data[offset], nil
}
