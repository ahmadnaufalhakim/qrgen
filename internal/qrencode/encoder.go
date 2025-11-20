package qrencode

import "github.com/ahmadnaufalhakim/qrgen/internal/qrconst"

type Encoder interface {
	Encode() ([]string, error)
}

func NewEncoder(s string) Encoder {
	encodingMode := DetermineEncodingMode(s)

	switch encodingMode {
	case qrconst.NumericMode:
		return NewNumericEncoder(s)

	case qrconst.AlphanumericMode:
		return NewAlphanumericEncoder(s)

	case qrconst.KanjiMode:
		return NewKanjiEncoder(s)

	case qrconst.ByteMode:
		return NewByteEncoder(s)

		// case qrconst.ECIMode:
		// 	return NewECIEncoder(s)
	}

	return nil
}
