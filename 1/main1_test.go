package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
)

func TestValueAndType(t *testing.T) {
	got := ValueAndType(42, "%d")
	if !strings.Contains(got, "Type: int") {
		t.Fatalf("expected type int in %q", got)
	}
	if !strings.Contains(got, "Value: 42") {
		t.Fatalf("expected value 42 in %q", got)
	}
}

func TestCombinedString(t *testing.T) {
	v := DefaultVars()
	got := CombinedString(v)

	mustContain := []string{
		"Dec: 42",
		"Oct: 052",
		"Hex: 0X2A",
		"Float: 3.14",
		"Str: Golang",
		"Bool: true",
		"Cpx: (1+2i)",
	}
	for _, sub := range mustContain {
		if !strings.Contains(got, sub) {
			t.Fatalf("combined string must contain %q; got %q", sub, got)
		}
	}
}

func TestToRunes(t *testing.T) {
	s := "aБ"
	r := ToRunes(s)
	if len(r) != 2 {
		t.Fatalf("expected 2 runes, got %d", len(r))
	}
	if string(r) != s {
		t.Fatalf("expected %q, got %q", s, string(r))
	}
}

func TestInsertSaltMiddle(t *testing.T) {
	tests := []struct {
		name string
		in   string
		salt string
		want string
	}{
		{name: "even", in: "abcd", salt: "X", want: "abXcd"},
		{name: "odd", in: "abc", salt: "X", want: "aXbc"},
		{name: "empty", in: "", salt: "go-2024", want: "go-2024"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(InsertSaltMiddle([]rune(tt.in), []rune(tt.salt)))
			if got != tt.want {
				t.Fatalf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestSaltAndHash(t *testing.T) {
	in := "abcd"
	salt := "go-2024"

	gotModified, gotHash := SaltAndHash(in, salt)

	wantModified := "abgo-2024cd"
	if gotModified != wantModified {
		t.Fatalf("modified: want %q, got %q", wantModified, gotModified)
	}

	wantHash := fmt.Sprintf("%x", sha256.Sum256([]byte(wantModified)))
	if gotHash != wantHash {
		t.Fatalf("hash: want %q, got %q", wantHash, gotHash)
	}
}
