package main

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrencode"
)

func main() {

	NumericString := "8675309"
	fmt.Println(qrencode.DetermineEncodingMode(NumericString))

	AlphanumericString := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0192837465 $%*+-./:"
	fmt.Println(qrencode.DetermineEncodingMode(AlphanumericString))

	KanjiString := "だから僕は音楽をやめた"
	fmt.Println(qrencode.DetermineEncodingMode(KanjiString))

	ByteString := "Hello, 世界!"
	fmt.Println(qrencode.DetermineEncodingMode(ByteString))

	////////////////////////////
	QREncoder := qrencode.NewEncoder(NumericString)
	fmt.Println(QREncoder.Encode())
}
