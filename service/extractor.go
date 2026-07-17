package service

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"

	"github.com/xiaojam/goqris/gozxing"
	"github.com/xiaojam/goqris/gozxing/qrcode"
)

func ExtractFromImage(file io.Reader) (string, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("gagal decode gambar: %v", err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", fmt.Errorf("gagal memproses bitmap: %v", err)
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", fmt.Errorf("gagal membaca QR Code: %v", err)
	}

	return result.GetText(), nil
}

func ExtractFromPath(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return ExtractFromImage(file)
}

func ExtractFromURL(urlStr string) (string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return ExtractFromImage(resp.Body)
}
