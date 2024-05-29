package utils

import (
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(url string) ([]byte, error) {
	return qrcode.Encode(url, qrcode.Medium, 256)
}
