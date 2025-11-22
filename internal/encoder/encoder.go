package encoder

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
	"golang.org/x/text/encoding/japanese"
)

type Encoder interface {
	Encode() ([]string, error)
	CharCount() int
	Mode() qrconst.EncodingMode
}

func NewEncoder(s string) Encoder {
	encodingMode := determineEncodingMode(s)

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

func determineEncodingMode(s string) qrconst.EncodingMode {
	isNumeric := true
	isAlphanumeric := true
	isShiftJIS := true

	sjisEncoder := japanese.ShiftJIS.NewEncoder()

	for _, r := range s {
		if isNumeric || isAlphanumeric || isShiftJIS {
			if isNumeric {
				if r < '0' || r > '9' {
					isNumeric = false
				}
			}

			if isAlphanumeric {
				if _, ok := tables.AlphanumericValues[r]; !ok {
					isAlphanumeric = false
				}
			}

			if isShiftJIS {
				sjisBytes, err := sjisEncoder.Bytes([]byte(string(r)))

				if err != nil || len(sjisBytes) != 2 {
					isShiftJIS = false
				} else {
					sjisValue := (uint16(sjisBytes[0]) << 8) | uint16(sjisBytes[1])
					if !(sjisValue >= 0x8140 && sjisValue <= 0x9FFC) &&
						!(sjisValue >= 0xE040 && sjisValue <= 0xEBBF) {
						isShiftJIS = false
					}
				}
			}
		} else {
			return qrconst.ByteMode
		}
	}

	if isNumeric {
		return qrconst.NumericMode
	}
	if isAlphanumeric {
		return qrconst.AlphanumericMode
	}
	if isShiftJIS {
		return qrconst.KanjiMode
	}
	return qrconst.ByteMode
}
