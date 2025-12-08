package qrcode

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

type QRCode struct {
	Version     int
	ECLevel     qrconst.ErrorCorrectionLevel
	Size        int
	MessageBits string
	Modules     [][]bool
	Patterns    [][]qrconst.FunctionPattern
	MaskNum     int
}

func NewQRCode(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	messageBitString string,
) *QRCode {
	size := (version-1)*4 + 21
	messageBits := messageBitString
	modules := make([][]bool, size)
	patterns := make([][]qrconst.FunctionPattern, size)
	for i := range size {
		modules[i] = make([]bool, size)
		patterns[i] = make([]qrconst.FunctionPattern, size)
	}
	maskNum := 0

	return &QRCode{
		Version:     version,
		ECLevel:     ecLevel,
		Size:        size,
		MessageBits: messageBits,
		Modules:     modules,
		Patterns:    patterns,
		MaskNum:     maskNum,
	}
}
