package encoder

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
	"golang.org/x/text/encoding/japanese"
)

type Encoder interface {
	Encode() ([]string, error)
	CharCount() int
	Mode() qrconst.EncodingMode
}

func NewEncoder(s string, encMode *qrconst.EncodingMode) (Encoder, error) {
	encodingMode := determineEncodingMode(s)
	if encMode != nil {
		if !isEncodingModeValid(s, *encMode) {
			return nil, fmt.Errorf(
				"input string %q is not valid for encoding mode %s",
				s,
				encMode.String(),
			)
		}

		encodingMode = *encMode
	}

	switch encodingMode {
	case qrconst.NumericMode:
		return NewNumericEncoder(s), nil

	case qrconst.AlphanumericMode:
		return NewAlphanumericEncoder(s), nil

	case qrconst.KanjiMode:
		return NewKanjiEncoder(s), nil

	case qrconst.ByteMode:
		return NewByteEncoder(s), nil

		// case qrconst.ECIMode:
		// 	return NewECIEncoder(s)
	}

	return nil, fmt.Errorf("unknown encoding mode: %v", encodingMode)
}

func determineEncodingMode(s string) qrconst.EncodingMode {
	isNumericEncodable := true
	isAlphanumericEncodable := true
	isKanjiEncodable := true

	for _, r := range s {
		if isNumericEncodable || isAlphanumericEncodable || isKanjiEncodable {
			if isNumericEncodable {
				isNumericEncodable = isNumeric(r)
			}
			if isAlphanumericEncodable {
				isAlphanumericEncodable = isAlphanumeric(r)
			}
			if isKanjiEncodable {
				isKanjiEncodable = isKanji(r)
			}
		} else {
			return qrconst.ByteMode
		}
	}

	switch {
	case isNumericEncodable:
		return qrconst.NumericMode
	case isAlphanumericEncodable:
		return qrconst.AlphanumericMode
	// case isKanjiEncodable:
	// 	return qrconst.KanjiMode
	default:
		return qrconst.ByteMode
	}
}

func isNumeric(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}

	return true
}

func isAlphanumeric(r rune) bool {
	if _, ok := tables.AlphanumericValues[r]; !ok {
		return false
	}

	return true
}

func isKanji(r rune) bool {
	sjisEncoder := japanese.ShiftJIS.NewEncoder()

	sjisBytes, err := sjisEncoder.Bytes([]byte(string(r)))

	if err != nil || len(sjisBytes) != 2 {
		return false
	} else {
		sjisValue := (uint16(sjisBytes[0]) << 8) | uint16(sjisBytes[1])
		if !(sjisValue >= 0x8140 && sjisValue <= 0x9FFC) &&
			!(sjisValue >= 0xE040 && sjisValue <= 0xEBBF) {
			return false
		}

		// Convert Shift-JIS -> JIS X 0208 row/column (QR Spec)
		var row, col int

		// Row
		if sjisBytes[0] <= 0x9F {
			row = int(sjisBytes[0]-0x81)*2 + 1
		} else {
			row = int(sjisBytes[0]-0xC1)*2 + 1
		}
		if sjisBytes[1] >= 0x9F {
			row++
		}

		// Column
		if sjisBytes[1] >= 0x9F {
			col = int(sjisBytes[1] - 0x7E)
		} else {
			col = int(sjisBytes[1] - 0x40)
		}

		// QR in Kanji mode uses JIS X 0208 rows/cols 1..94
		return (row >= 1 && row <= 94) && (col >= 1 && col <= 94)
	}
}

func isEncodingModeValid(s string, encMode qrconst.EncodingMode) bool {
	if encMode == qrconst.ByteMode {
		return true
	}

	isValid := true
	var validate func(r rune) bool
	switch encMode {
	case qrconst.NumericMode:
		validate = isNumeric
	case qrconst.AlphanumericMode:
		validate = isAlphanumeric
	case qrconst.KanjiMode:
		validate = isKanji
	}

	for _, r := range s {
		if isValid {
			isValid = validate(r)
		} else {
			return false
		}
	}

	return isValid
}
