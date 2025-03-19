package utils

import (
	"fmt"
	"strings"
)

type Resp struct {
	encodedData string
	decodedData []string
}

func (r *Resp) EncodeData(decodedData []string) string {
	encodedData := ""

	for _, val := range decodedData {
		if encodedData == "" {
			encodedData = fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)
		} else {
			encodedData += fmt.Sprintf("$%d\r\n%s\r\n", len(val), val)
		}
	}

	if len(decodedData) > 1 {
		return fmt.Sprintf("*%d\r\n", len(decodedData)) + encodedData
	}

	return encodedData
}

func (r *Resp) DecodeData(encodedData string) []string {
	decodedData := strings.Fields(strings.ToLower(encodedData))
	return decodedData
}
