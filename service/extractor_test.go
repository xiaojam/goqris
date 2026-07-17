package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestExtractFromImage_Error_NotAnImage(t *testing.T) {
	invalidData := bytes.NewReader([]byte("ini adalah file teks palsu, bukan gambar"))

	_, err := ExtractFromImage(invalidData)

	if err == nil {
		t.Error("Harusnya error karena format bukan gambar, tapi ternyata sukses")
	}
}

func TestExtractFromPath_Error_FileNotFound(t *testing.T) {
	fakePath := "./folder_ngasal/gambar_ilang.png"

	_, err := ExtractFromPath(fakePath)

	if err == nil {
		t.Error("Harusnya error karena file tidak ada, tapi ternyata sukses")
	}
}

func TestExtractFromURL_Error_HTTPStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	_, err := ExtractFromURL(server.URL)

	if err == nil {
		t.Error("Harusnya error karena HTTP status 404, tapi ternyata sukses")
	}
}

func TestExtractFromPath_Success(t *testing.T) {
	path := "./testdata/sample_qris.png"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Skip("Melewati Test: file gambar asli ./testdata/sample_qris.png belum disiapkan.")
	}

	result, err := ExtractFromPath(path)

	if err != nil {
		t.Fatalf("Harusnya sukses baca gambar, tapi dapat error: %v", err)
	}
	if result == "" {
		t.Error("Hasil decode QR Code kosong")
	}
}

func TestExtractFromURL_Success(t *testing.T) {
	path := "./testdata/sample_qris.png"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Skip("Melewati Test: file gambar asli ./testdata/sample_qris.png belum disiapkan.")
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileBytes, _ := os.ReadFile(path)
		w.WriteHeader(http.StatusOK)
		w.Write(fileBytes)
	}))
	defer server.Close()

	result, err := ExtractFromURL(server.URL)

	if err != nil {
		t.Fatalf("Harusnya sukses download & baca QR, tapi dapat error: %v", err)
	}
	if result == "" {
		t.Error("Hasil decode QR Code dari URL kosong")
	}
}
