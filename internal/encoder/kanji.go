package encoder

import (
	"fmt"
	"strconv"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"golang.org/x/text/encoding/japanese"
)

var sjisEncoder = japanese.ShiftJIS.NewEncoder()

type KanjiEncoder struct {
	s string
}

func NewKanjiEncoder(s string) *KanjiEncoder {
	return &KanjiEncoder{
		s: s,
	}
}

// Convert a kanji rune into an integer.
func kanjiRuneToInt(r rune) (int, error) {
	sjisBytes, err := sjisEncoder.Bytes([]byte(string(r)))
	if err != nil {
		return 0, err
	}

	sjisValue := uint16(sjisBytes[0])<<8 | uint16(sjisBytes[1])

	var base uint16
	switch {
	case sjisValue >= 0x8140 && sjisValue <= 0x9FFC:
		base = 0x8140
	case sjisValue >= 0xE040 && sjisValue <= 0xEBBF:
		base = 0xC140
	default:
		return 0, fmt.Errorf("rune %q not in QR Kanji Shift-JIS ranges", r)
	}

	sub := sjisValue - base
	msb := (sub >> 8) & 0x00FF
	lsb := sub & 0x00FF

	val := int(msb)*0xC0 + int(lsb)
	if val < 0 || val > 0x1FFF {
		return 0, fmt.Errorf("calculated kanji value out of 13-bit range: %d", val)
	}

	return val, nil
}

// Encode encodes the input string into QR Code Kanji Mode.
//
// Each rune is converted into its corresponding Shift JIS value.
// Valid QR Kanji must fall within the ranges used by the QR spec.
// The Shift JIS value is then transformed into a 13-bit code word.
//
// If a rune cannot be mapped to a valid QR Kanji code, an error is returned.
func (ke *KanjiEncoder) Encode() ([]string, error) {
	runes := []rune(ke.s)
	bitStrings := make([]string, len(runes))

	for i, r := range runes {
		kanjiInt, err := kanjiRuneToInt(r)
		if err != nil {
			return nil, err
		}

		b := strconv.FormatInt(int64(kanjiInt), 2)
		bitStrings[i] = padBitString(b, 13)
	}

	return bitStrings, nil
}

// CharCount returns the number of Kanji characters in the string.
// In Kanji mode, each rune corresponds to one encoded Kanji symbol.
func (ke *KanjiEncoder) CharCount() int {
	return len([]rune(ke.s))
}

// Mode returns kanji mode EncodingMode
func (ke *KanjiEncoder) Mode() qrconst.EncodingMode {
	return qrconst.KanjiMode
}
