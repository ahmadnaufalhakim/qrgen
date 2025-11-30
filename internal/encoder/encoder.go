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
	isKanji := true

	sjisEncoder := japanese.ShiftJIS.NewEncoder()

	for _, r := range s {
		if isNumeric || isAlphanumeric || isKanji {
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

			if isKanji {
				sjisBytes, err := sjisEncoder.Bytes([]byte(string(r)))

				if err != nil || len(sjisBytes) != 2 {
					isKanji = false
				} else {
					sjisValue := (uint16(sjisBytes[0]) << 8) | uint16(sjisBytes[1])
					if !(sjisValue >= 0x8140 && sjisValue <= 0x9FFC) &&
						!(sjisValue >= 0xE040 && sjisValue <= 0xEBBF) {
						isKanji = false
						continue
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
					isKanji = (row >= 1 && row <= 94) && (col >= 1 && col <= 94)
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
	if isKanji {
		return qrconst.KanjiMode
	}
	return qrconst.ByteMode
}
