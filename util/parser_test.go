package util

import (
	"testing"

	"goqris/model"
)

func TestParseTLV_Success(t *testing.T) {
	input := "000201010211"

	items := ParseTLV(input)

	if len(items) != 2 {
		t.Fatalf("Harusnya mendapat 2 item TLV, tapi dapat %d", len(items))
	}
	if items[0].Tag != "00" || items[0].Value != "01" {
		t.Errorf("Parsing Tag 00 gagal. Dapat: %v", items[0])
	}
}

func TestParseTLV_Error_IncompleteString(t *testing.T) {
	input := "5904To"

	items := ParseTLV(input)

	if len(items) != 0 {
		t.Errorf("Harusnya item kosong karena format terpotong, tapi terbaca %d item", len(items))
	}
}

func TestBuildQRISString_Success(t *testing.T) {
	items := []model.TLVItem{
		{Tag: "00", Length: "02", Value: "01"},
		{Tag: "63", Length: "04", Value: "A1B2"},
	}

	result := BuildQRISString(items)

	expected := "0002016304"
	if result != expected {
		t.Errorf("Build string gagal. Harusnya '%s', dapat '%s'", expected, result)
	}
}
