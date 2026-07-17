# GoQRIS

GoQRIS is a lightweight, modular Golang library designed to process and manipulate QRIS (Quick Response Code Indonesian Standard) codes. This library focuses on converting a Static QRIS into a Dynamic QRIS by dynamically injecting a transaction amount (Tag 54) and accurately recalculating the mandatory Checksum (CRC16-CCITT).

This library is built with a strictly Zero-Dependency approach in its `go.mod`. 

## Features

- **Convert Static to Dynamic:** Automatically updates the initiation method indicator (Tag 01) from 11 (Static) to 12 (Dynamic).
- **Nominal Injection:** Seamlessly injects the transaction amount (Tag 54) into the TLV (Tag-Length-Value) structure.
- **Auto CRC16 Recalculation:** Automatically drops the old checksum and recalculates a highly accurate CRC16-CCITT checksum for the new string.
- **Multi-Input Support:** Extract and process QRIS from multiple sources:
  - Raw String
  - Local File Path (`.png`, `.jpg`)
  - Internet Image URL
  - `io.Reader` (Perfect for HTTP API `multipart/form-data` uploads).

## Scanner Engine Origin

The internal QR Code reader engine used in this library is an extracted, vendored version of [makiuchi-d/gozxing](https://github.com/makiuchi-d/gozxing), which is a Golang port of the renowned ZXing (Zebra Crossing) library. 

To maintain the strict zero-dependency rule in our `go.mod` file, the necessary computer vision and Reed-Solomon error correction components from `gozxing` were manually copied and embedded directly into the `gozxing/` directory of this repository. This vendor-in-tree approach guarantees that this library remains robust and immune to any future upstream changes, repository deletions, or broken dependencies, giving you full control over the core logic.

## Installation

To install the library, simply run the following `go get` command:

```bash
go get github.com/xiaojam/goqris
```

## Usage and Sample Implementation

Below is a complete, ready-to-run example demonstrating how to use the library across different scenarios (String, URL, and Local File). 

You can create a `main.go` file in your project and run this code:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"goqris/service" 
)

func main() {
	// ---------------------------------------------------------
	// FROM STRING
	// ---------------------------------------------------------

	nominal := "125000"

	rawString := "0002010102115204411153033605802ID5904Toko6007Jakarta6304A1B2"
	
	resStr, err := service.ProcessFromString(rawString, nominal)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("Nominal     : Rp %s\n", nominal)
		fmt.Printf("Old String  : %s\n", resStr.OriginalString)
		fmt.Printf("New String  : %s\n", resStr.NewString)
		fmt.Printf("New CRC     : %s\n", resStr.CRCValue)
	}

	// ---------------------------------------------------------
	// FROM URL
	// ---------------------------------------------------------
	urlGambar := "https://raw.githubusercontent.com/xiaojam/goqris/main/example/qris.jpg"
	
	resUrl, err := service.ProcessFromURL(urlGambar, nominal)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("New String  :", resUrl.NewString)
	}
	
	// ---------------------------------------------------------
	// FROM LOCAL PATH
	// ---------------------------------------------------------
	dummyPath := "./qris.jpg"
	_ = os.WriteFile(dummyPath, []byte("this is not image"), 0644)
	defer os.Remove(dummyPath) 

	resPath, err := service.ProcessFromPath(dummyPath, nominal)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("New String  :", resPath.NewString)
	}
}
```

### Implementing in HTTP APIs (Using io.Reader)

If you are building a REST API where users upload QRIS images, you can directly pass the `multipart.File` into the library using `ProcessFromImage`:

```go
// Inside your HTTP Handler
file, _, err := r.FormFile("qris_image")
if err != nil {
    http.Error(w, "Failed to get file", http.StatusBadRequest)
    return
}
defer file.Close()

// file is of type multipart.File, which implements io.Reader
response, err := service.ProcessFromImage(file, "200000")
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

fmt.Fprintln(w, "Dynamic QRIS: ", response.NewString)
```

## Project Structure

- `model/` : Contains the TLV item structures and Process Response schemas.
- `util/` : Houses the TLV string parser engine and the CRC16-CCITT cryptographic recalculation logic.
- `service/` : The Core Business Logic handling the transformation of Static QRIS to Dynamic.
- `gozxing/` : An internally vendored, zero-dependency QR code decoding engine ported from makiuchi-d/gozxing.

## License

This project is licensed under the MIT License.
