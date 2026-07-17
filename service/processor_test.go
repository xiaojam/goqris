package service

import (
	"testing"
)

func TestProcessFromString_Success(t *testing.T) {
	staticQRIS := "0002010102115204411153033605802ID5904Toko6007Jakarta6304A1B2"
	nominal := "50000"

	response, err := ProcessFromString(staticQRIS, nominal)

	if err != nil {
		t.Fatalf("Harusnya sukses, tapi dapat error: %v", err)
	}

	if response.CRCValue == "" {
		t.Error("CRCValue tidak boleh kosong")
	}

	foundDynamic := false
	for _, item := range response.NewTLV {
		if item.Tag == "01" && item.Value == "12" {
			foundDynamic = true
			break
		}
	}
	if !foundDynamic {
		t.Error("Tag 01 tidak berubah menjadi 12 (Dinamis)")
	}
}

func TestProcessFromString_Error_InvalidLength(t *testing.T) {
	invalidQRIS := "000"
	nominal := "50000"

	response, err := ProcessFromString(invalidQRIS, nominal)

	if err == nil {
		t.Error("Harusnya menghasilkan error karena panjang string tidak valid, tapi ternyata sukses")
	}

	if response != nil {
		t.Error("Response harusnya nil jika terjadi error")
	}
}
