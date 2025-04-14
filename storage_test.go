package main

import (
	"fmt"
	"testing"
)

func TestMemoryStore(t *testing.T) {
	m := NewMemoryStore()
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foobaz_%d", i)
		latestOffset, err := m.Push([]byte(key))
		if err != nil {
			t.Error(err)
		}

		data, err := m.Fetch(latestOffset)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(string(data))
	}
}
