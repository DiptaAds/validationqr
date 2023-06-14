package controller

import (
	"fmt"
	"image/color"
	"time"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(data string) (string, error) {
	qrCodee := fmt.Sprintf("%d.png", time.Now().UnixNano())
	err := qrcode.WriteColorFile(data, qrcode.Medium, 256, color.Black, color.White, qrCodee)
	if err != nil {
		return "", err
	}

	return qrCodee, nil
}
