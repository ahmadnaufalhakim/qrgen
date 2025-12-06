package main

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

func main() {

	// text := "8675309" //Numeric
	// text := "HELLO WORLD" //Alphanumeric
	// text := "だから僕は音楽をやめた" //Kanji
	text := "hello world😎" //Byte

	qrBuilder := qrcode.NewQRBuilder(text)
	qrCode, err := qrBuilder.
		WithErrorCorrectionLevel(qrconst.L).
		Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	qrCode.RenderPNG("testing.png", 15)
}
