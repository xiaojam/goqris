package util

import (
	"testing"
)

func TestCalculateCRC16_Success(t *testing.T) {
	input := "0002010102125204411153033605802ID5904Toko6007Jakarta"

	expected := "9369"

	result := CalculateCRC16(input)

	if result != expected {
		t.Errorf("Perhitungan CRC salah. Harusnya %s, dapat %s", expected, result)
	}
}

func TestCalculateCRC16_EmptyString(t *testing.T) {
	input := ""

	result := CalculateCRC16(input)

	expected := "FFFF"
	if result != expected {
		t.Errorf("CRC string kosong harusnya %s, dapat %s", expected, result)
	}
}
