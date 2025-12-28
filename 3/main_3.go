package main

import (
	"maps"
	"sync"
)

type StringIntMap struct {
	data map[string]int
	m    sync.RWMutex
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (si *StringIntMap) Add(key string, value int) { //перезатирает старое значение, если оно есть по ключу
	si.m.Lock()
	si.data[key] = value
	si.m.Unlock()
}

func (si *StringIntMap) Remove(key string) {
	si.m.Lock()
	delete(si.data, key)
	si.m.Unlock()
}
func (si *StringIntMap) Copy() map[string]int {
	si.m.RLock()
	defer si.m.RUnlock()
	newMap := maps.Clone(si.data)
	return newMap
}

func (si *StringIntMap) Exists(key string) bool {
	si.m.RLock()
	defer si.m.RUnlock()
	if _, ok := si.data[key]; ok {
		return true
	}
	return false
}

func (si *StringIntMap) Get(key string) (int, bool) {
	si.m.RLock()
	defer si.m.RUnlock()
	if val, ok := si.data[key]; ok {
		return val, ok
	}
	return 0, false
}
