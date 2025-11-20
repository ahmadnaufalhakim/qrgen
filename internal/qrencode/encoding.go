package qrencode

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
	"golang.org/x/text/encoding/japanese"
)

type EncodingMode int

const (
	NumericMode      EncodingMode = 0b_0001
	AlphanumericMode EncodingMode = 0b_0010
	ByteMode         EncodingMode = 0b_0100
	KanjiMode        EncodingMode = 0b_1000
	ECIMode          EncodingMode = 0b_0111
)

func DetermineEncodingMode(s string) EncodingMode {
	if IsNumeric(s) {
		return NumericMode
	}

	if IsAlphanumeric(s) {
		return AlphanumericMode
	}

	if IsShiftJIS(s) {
		return KanjiMode
	}

	return ByteMode
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
