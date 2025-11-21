package main

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrencode"
)

func main() {

	// InputString := "8675309" //Numeric
	InputString := "HELLO WORLD" //Alphanumeric
	// InputString := "だから僕は音楽をやめた" //Kanji
	// InputString := "Hello, 世界!" //Byte

	////////////////////////////
	QREncoder := qrencode.NewEncoder(InputString)
	ErrorCorrectionLevel := qrconst.Q

	version, err := qrencode.DetermineVersion(
		QREncoder.Mode(),
		ErrorCorrectionLevel,
		QREncoder.CharCount(),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("version: %d-%c\n", version, ErrorCorrectionLevel)

	encodedData, err := QREncoder.Encode()
	if err != nil {
		fmt.Println(err)
		return
	}

	bits := append([]string{
		qrencode.ModeIndicator(QREncoder.Mode()),
		qrencode.CharCountIndicator(
			QREncoder.Mode(),
			version,
			QREncoder.CharCount(),
		),
	}, encodedData...)
	fmt.Println(bits)
}
