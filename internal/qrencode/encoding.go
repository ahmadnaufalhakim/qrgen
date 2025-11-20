package qrencode

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
	"golang.org/x/text/encoding/japanese"
)

func DetermineEncodingMode(s string) qrconst.EncodingMode {
	if IsNumeric(s) {
		return qrconst.NumericMode
	}

	if IsAlphanumeric(s) {
		return qrconst.AlphanumericMode
	}

	if IsShiftJIS(s) {
		return qrconst.KanjiMode
	}

	return qrconst.ByteMode
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if _, ok := tables.AlphanumericValues[r]; !ok {
			return false
		}
	}

	return true
}

func IsShiftJIS(s string) bool {
	sjisEncoder := japanese.ShiftJIS.NewEncoder()
	for _, r := range s {
		sjisBytes, err := sjisEncoder.Bytes([]byte(string(r)))
		if err != nil {
			return false
		}

		// Check if the character can be converted to 2-bytes Shift JIS character
		if len(sjisBytes) != 2 {
			return false
		}

		// Check if the Shift JIS byte value is in the range
		// of 0x8140 to 0x9FFC, or 0xE040 to 0xEBBF
		sjisValue := (uint16(sjisBytes[0]) << 8) | uint16(sjisBytes[1])
		if (sjisValue >= 0x8140 && sjisValue <= 0x9FFC) ||
			(sjisValue >= 0xE040 && sjisValue <= 0xEBBF) {
			continue
		} else {
			return false
		}
	}

	return true
}
