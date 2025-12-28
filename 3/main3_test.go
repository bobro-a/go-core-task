package main

import (
	"sync"
	"testing"
)

func TestStringIntMap_Add(t *testing.T) {
	m := NewStringIntMap()
	tests := []struct {
		key   string
		value int
	}{
		{key: "a", value: 1},
		{key: "b", value: 2},
		{key: "a", value: 0},
	}
	for _, tt := range tests {
		m.Add(tt.key, tt.value)
		if _, ok := m.data[tt.key]; !ok {
			t.Errorf("key %s was not added\n", tt.key)
		}
	}
	if m.data["a"] != 0 {
		t.Errorf("key 'a' has not been overwritten")
	}
}

func TestStringIntMap_Get(t *testing.T) {
	m := NewStringIntMap()
	tests := []struct {
		key   string
		value int
	}{
		{key: "a", value: 1},
		{key: "b", value: 2},
		{key: "c", value: 3},
	}
	for _, tt := range tests {
		m.Add(tt.key, tt.value)
		if val, ok := m.data[tt.key]; !ok {
			t.Errorf("value with key %s is missing\n", tt.key)
		} else if val != tt.value {
			t.Errorf("key value must be %d, not %d\n", tt.value, val)
		}
	}
}

func TestStringIntMap_Exists(t *testing.T) {
	m := NewStringIntMap()
	tests := []struct {
		key   string
		value int
	}{
		{key: "a", value: 1},
		{key: "b", value: 2},
		{key: "c", value: 3},
	}
	for _, tt := range tests {
		m.Add(tt.key, tt.value)
		if _, ok := m.data[tt.key]; !ok {
			t.Errorf("value with key %s is missing\n", tt.key)
		}
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	m := NewStringIntMap()
	tests := []struct {
		key   string
		value int
	}{
		{key: "a", value: 1},
		{key: "b", value: 2},
		{key: "c", value: 3},
	}
	for _, tt := range tests {
		m.Add(tt.key, tt.value)
	}
	newData := m.Copy()
	newData["a"] = 4
	if m.data["a"] == 4 {
		t.Errorf(" Copy function returns the same memory area")
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	m := NewStringIntMap()
	tests := []struct {
		key   string
		value int
	}{
		{key: "a", value: 1},
		{key: "b", value: 2},
		{key: "c", value: 3},
	}
	for _, tt := range tests {
		m.Add(tt.key, tt.value)
		m.Remove(tt.key)
		if _, ok := m.data[tt.key]; ok {
			t.Errorf("value with key %s has not been deleted\n", tt.key)
		}
	}
}

func TestConcurrencyStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	numGoroutines := 100
	operationsPerGoroutine := 1000

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				key := "key_" + string(rune(id%10))
				m.Add(key, j)
			}
		}(i)
	}

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				key := "key_" + string(rune(id%10))
				m.Get(key)
			}
		}(i)
	}

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				key := "key_" + string(rune(id%10))
				m.Exists(key)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine; j++ {
				m.Copy()
			}
		}()
	}

	for i := 0; i < numGoroutines/10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operationsPerGoroutine/10; j++ {
				key := "key_" + string(rune((id+j)%10))
				m.Remove(key)
			}
		}(i)
	}

	wg.Wait()

	m.Add("test_key", 42)
	val, ok := m.Get("test_key")

	if !ok {
		t.Error("Get должен вернуть true для добавленного ключа")
	}
	if val != 42 {
		t.Errorf("Ожидал значение 42, получил %d", val)
	}
}
