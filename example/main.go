package main

import (
	"fmt"
	"log"
	"os"

	"goqris/service"
)

func main() {
	// ---------------------------------------------------------
	// SKENARIO 1: DARI STRING QRIS
	// ---------------------------------------------------------

	nominal := "125000"

	rawString := "0002010102115204411153033605802ID5904Toko6007Jakarta6304A1B2"

	resStr, err := service.ProcessFromString(rawString, nominal)
	if err != nil {
		log.Println("Gagal proses string:", err)
	} else {
		fmt.Printf("Nominal      : Rp %s\n", nominal)
		fmt.Printf("String Lama  : %s\n", resStr.OriginalString)
		fmt.Printf("String Baru  : %s\n", resStr.NewString)
		fmt.Printf("CRC Baru     : %s\n", resStr.CRCValue)
	}

	// ---------------------------------------------------------
	// SKENARIO 2: DARI URL INTERNET
	// ---------------------------------------------------------
	urlGambar := "https://raw.githubusercontent.com/xiaojam/goqris/main/example/qris.jpg"

	resUrl, err := service.ProcessFromURL(urlGambar, nominal)
	if err != nil {
		log.Println("Gagal baca QR dari URL:", err)
	} else {
		fmt.Println("Sukses baca dari URL!")
		fmt.Println("String Baru  :", resUrl.NewString)
	}

	// ---------------------------------------------------------
	// SKENARIO 3: DARI FILE LOKAL
	// ---------------------------------------------------------
	dummyPath := "./qris.jpg"
	_ = os.WriteFile(dummyPath, []byte("ini bukan gambar"), 0644)
	defer os.Remove(dummyPath)

	resPath, err := service.ProcessFromPath(dummyPath, nominal)
	if err != nil {
		log.Println("Gagal baca file lokal:", err)
	} else {
		fmt.Println("Sukses baca dari file lokal!")
		fmt.Println("String Baru  :", resPath.NewString)
	}
}
