package main

import (
	"testing"
	"time"
)

func mustPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic, got none")
		}
	}()
	f()
}

func TestMyWaitGroup_Wait(t *testing.T) {
	wg := NewMyWaitGroup(100)
	wg.Add(20)

	for range 20 {
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
		}()
	}

	wg.Wait()
}

func TestMyWaitGroup_TwoWait(t *testing.T) {
	wg := NewMyWaitGroup(100)
	wg.Add(20)

	for range 20 {
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
		}()
	}

	wg.Wait()

	wg.Add(20)

	for range 20 {
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
		}()
	}

	wg.Wait()
}

func TestMyWaitGroup_DoneWithoutAddPanics(t *testing.T) {
	wg := NewMyWaitGroup(0)
	mustPanic(t, func() { wg.Done() })
}

func TestMyWaitGroup_AddOverLimitPanics(t *testing.T) {
	wg := NewMyWaitGroup(1)
	mustPanic(t, func() { wg.Add(2) })
}
