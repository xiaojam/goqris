package util

import (
	"github.com/xiaojam/goqris/model"
	"strconv"
	"strings"
)

func ParseTLV(data string) []model.TLVItem {
	var items []model.TLVItem
	i := 0
	for i < len(data) {
		if i+4 > len(data) {
			break
		}
		tag := data[i : i+2]
		lengthStr := data[i+2 : i+4]
		length, _ := strconv.Atoi(lengthStr)

		valueEnd := i + 4 + length
		if valueEnd > len(data) {
			break
		}
		value := data[i+4 : valueEnd]

		items = append(items, model.TLVItem{Tag: tag, Length: lengthStr, Value: value})
		i = valueEnd
	}
	return items
}

func BuildQRISString(items []model.TLVItem) string {
	var builder strings.Builder
	for _, item := range items {
		if item.Tag != "63" {
			builder.WriteString(item.Tag)
			builder.WriteString(item.Length)
			builder.WriteString(item.Value)
		}
	}
	builder.WriteString("6304")
	return builder.String()
}
