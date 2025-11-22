package main

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/encoder"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrencode"
)

func main() {

	// InputString := "8675309" //Numeric
	InputString := "HELLO WORLD" //Alphanumeric
	// InputString := "だから僕は音楽をやめた" //Kanji
	// InputString := "Hello, 世界!" //Byte

	////////////////////////////
	QREncoder := encoder.NewEncoder(InputString)
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

	bitStrings := append([]string{
		qrencode.ModeIndicator(QREncoder.Mode()),
		qrencode.CharCountIndicator(
			QREncoder.Mode(),
			version,
			QREncoder.CharCount(),
		),
	}, encodedData...)
	fmt.Println(bitStrings)

	fmt.Println(qrencode.AssembleDataCodewords(
		ErrorCorrectionLevel,
		version,
		bitStrings,
	))

	// // Sanity check
	// // Error correction code words and block information
	// for version := range 40 {
	// 	for _, ecLevel := range []qrconst.ErrorCorrectionLevel{qrconst.L, qrconst.M, qrconst.Q, qrconst.H} {
	// 		ECBlock := tables.ECBlocks[ecLevel][version]
	// 		TotalDataCodeword := tables.DataCodewords[ecLevel][version]

	// 		calc := ECBlock.Group1Blocks*ECBlock.Group1DataCodewords + ECBlock.Group2Blocks*ECBlock.Group2DataCodewords
	// 		if TotalDataCodeword != calc {
	// 			panic(fmt.Sprintf("DataCodewords %d != calc %d", TotalDataCodeword, calc))
	// 		}
	// 	}
	// }
}
