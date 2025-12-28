package main

import (
	"crypto/sha256"
	"fmt"
)

type Vars struct {
	NumDecimal     int
	NumOctal       int
	NumHexadecimal int
	Pi             float64
	Name           string
	IsActive       bool
	ComplexNum     complex64
}

func DefaultVars() Vars {
	return Vars{
		NumDecimal:     42,
		NumOctal:       052,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     1 + 2i,
	}
}

func ValueAndType(v any, valueFmt string) string {
	return fmt.Sprintf("Value: "+valueFmt+", Type: %T", v, v)
}

func DescribeAll(v Vars) []string {
	return []string{
		ValueAndType(v.NumDecimal, "%d"),
		ValueAndType(v.NumOctal, "%#o"),
		ValueAndType(v.NumHexadecimal, "%#X"),
		ValueAndType(v.Pi, "%.2f"),
		ValueAndType(v.Name, "%s"),
		ValueAndType(v.IsActive, "%t"),
		ValueAndType(v.ComplexNum, "%v"),
	}
}

func CombinedString(v Vars) string {
	return fmt.Sprintf(
		"Dec: %d | Oct: %#o | Hex: %#X | Float: %.2f | Str: %s | Bool: %t | Cpx: %v",
		v.NumDecimal, v.NumOctal, v.NumHexadecimal, v.Pi, v.Name, v.IsActive, v.ComplexNum,
	)
}

func ToRunes(s string) []rune {
	return []rune(s)
}

func InsertSaltMiddle(runes []rune, salt []rune) []rune {
	mid := len(runes) / 2

	out := make([]rune, 0, len(runes)+len(salt))
	out = append(out, runes[:mid]...)
	out = append(out, salt...)
	out = append(out, runes[mid:]...)
	return out
}

func SHA256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", sum)
}

func SaltAndHash(s string, salt string) (modified string, hashHex string) {
	r := ToRunes(s)
	salted := InsertSaltMiddle(r, []rune(salt))
	modified = string(salted)
	hashHex = SHA256Hex(modified)
	return modified, hashHex
}

func main() {
	v := DefaultVars()

	for _, line := range DescribeAll(v) {
		fmt.Println(line)
	}

	combined := CombinedString(v)
	fmt.Println("Combined:", combined)

	runes := ToRunes(combined)
	fmt.Println("Runes length:", len(runes))

	modified, hashHex := SaltAndHash(combined, "go-2024")
	fmt.Println("Modified:", modified)
	fmt.Println("SHA256:", hashHex)
}
