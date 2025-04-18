package main

import (
	"fmt"
	"sync"
)

type Storage interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

type MemoryStore struct {
	mu   sync.RWMutex
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (m *MemoryStore) Push(b []byte) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data = append(m.data, b)
	return len(m.data) - 1, nil
}

func (m *MemoryStore) Fetch(offset int) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if offset < 0 {
		return nil, fmt.Errorf("offset cannot be smaller than 0")
	}

	if len(m.data)-1 < offset {
		return nil, fmt.Errorf("offset (%d) is too high", offset)
	}

	return m.data[offset], nil
}
