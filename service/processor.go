package service

import (
	"fmt"
	"io"

	"goqris/model"
	"goqris/util"
)

func ProcessFromString(originalString string, nominal string) (*model.ProcessResponse, error) {
	if len(originalString) < 8 {
		return nil, fmt.Errorf("format string QRIS tidak valid")
	}

	originalTLV := util.ParseTLV(originalString)
	var newTLV []model.TLVItem

	for _, item := range originalTLV {
		if item.Tag == "01" && item.Value == "11" {
			item.Value = "12"
		}
		if item.Tag == "63" {
			continue
		}
		newTLV = append(newTLV, item)
	}

	nominalLength := fmt.Sprintf("%02d", len(nominal))
	newTLV = append(newTLV, model.TLVItem{Tag: "54", Length: nominalLength, Value: nominal})

	tempString := util.BuildQRISString(newTLV)
	crcValue := util.CalculateCRC16(tempString)

	newTLV = append(newTLV, model.TLVItem{Tag: "63", Length: "04", Value: crcValue})
	finalString := tempString + crcValue

	return &model.ProcessResponse{
		OriginalString: originalString,
		NewString:      finalString,
		OriginalTLV:    originalTLV,
		NewTLV:         newTLV,
		CRCValue:       crcValue,
		Nominal:        nominal,
	}, nil
}

func ProcessFromImage(file io.Reader, nominal string) (*model.ProcessResponse, error) {
	qrString, err := ExtractFromImage(file)
	if err != nil {
		return nil, err
	}
	return ProcessFromString(qrString, nominal)
}

func ProcessFromPath(path string, nominal string) (*model.ProcessResponse, error) {
	qrString, err := ExtractFromPath(path)
	if err != nil {
		return nil, err
	}
	return ProcessFromString(qrString, nominal)
}

func ProcessFromURL(url string, nominal string) (*model.ProcessResponse, error) {
	qrString, err := ExtractFromURL(url)
	if err != nil {
		return nil, err
	}
	return ProcessFromString(qrString, nominal)
}
