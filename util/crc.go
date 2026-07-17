package util

import (
	"fmt"
	"strings"
)

func CalculateCRC16(data string) string {
	crc := 0xFFFF
	for i := 0; i < len(data); i++ {
		crc ^= int(data[i]) << 8
		for j := 0; j < 8; j++ {
			if (crc & 0x8000) != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}
	crc &= 0xFFFF
	return strings.ToUpper(fmt.Sprintf("%04X", crc))
}
