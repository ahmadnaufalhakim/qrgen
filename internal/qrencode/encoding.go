package qrencode

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
	"golang.org/x/text/encoding/japanese"
)

func DetermineEncodingMode(s string) qrconst.EncodingMode {
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

func DetermineVersion(
	encMode qrconst.EncodingMode,
	errCorrectionLevel qrconst.ErrorCorrectionLevel,
	charCount int,
) (int, error) {
	switch errCorrectionLevel {
	case qrconst.L, qrconst.M, qrconst.Q, qrconst.H:
		charCapacity := tables.CharacterCapacities[encMode][errCorrectionLevel]

		for version := range 40 {
			lowCharacterCapacity := charCapacity[version]
			highCharacterCapacity := charCapacity[40-version-1]

			if lowCharacterCapacity >= charCount {
				return version + 1, nil
			}
			if highCharacterCapacity < charCount {
				if version != 0 {
					return 40 - version + 1, nil
				} else {
					return 0, fmt.Errorf("no version found that can encode %d characters", charCount)
				}
			}
		}
		return 0, fmt.Errorf("no version found that can encode %d characters", charCount)

	default:
		return 0, fmt.Errorf("invalid error correction level")
	}
}

func ModeIndicator(encMode qrconst.EncodingMode) string {
	return fmt.Sprintf("%04b", encMode)
}

func CharCountIndicator(
	encMode qrconst.EncodingMode,
	version int,
	charCount int,
) string {
	var charCountFormat string
	var idx int

	if version >= 1 && version <= 9 {
		idx = 0
	} else if version >= 10 && version <= 26 {
		idx = 1
	} else if version >= 27 && version <= 40 {
		idx = 2
	}
	charCountFormat = fmt.Sprintf("%%0%db", tables.CharacterCountIndicatorBits[encMode][idx])

	return fmt.Sprintf(charCountFormat, charCount)
}
