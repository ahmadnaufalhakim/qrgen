package qrencode

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

type AlphanumericEncoder struct {
	s string
}

func NewAlphanumericEncoder(s string) *AlphanumericEncoder {
	return &AlphanumericEncoder{
		s: s,
	}
}

// Convert an alphanumeric string of length 1,2 into an integer.
func alphanumericStrToInt(s string) int {
	result := 0
	for _, r := range s {
		result = (result * 45) + tables.AlphanumericValues[r]
	}

	return result
}

// Split alphanumeric string into groups of 2 (as per QR code specification).
func alphanumericSplit(s string) []string {
	var groups []string
	for i := 0; i < len(s); i += 2 {
		end := min(i+2, len(s))
		groups = append(groups, s[i:end])
	}

	return groups
}

// Convert one alphanumeric string group into its QR bit representation.
// 2 chars -> 11 bits, 1 char -> 6 bits
func encodeAlphanumericBits(group string, groupSize int) string {
	var binaryFormat string
	switch groupSize {
	case 2:
		binaryFormat = "%011b"
	case 1:
		binaryFormat = "%06b"
	}

	return fmt.Sprintf(binaryFormat, alphanumericStrToInt(group))
}

// Encode encodes the input string in QR Code Alphanumeric Mode.
//
// The string is split into groups of two characters. Each pair is
// encoded into an 11-bit value: (45 * value1 + value2).
// If the input has an odd number of characters, the final single
// character is encoded into a 6-bit value.
//
// The returned slice contains each group’s bit representation.
func (ae *AlphanumericEncoder) Encode() ([]string, error) {
	var bits []string
	for _, group := range alphanumericSplit(ae.s) {
		bits = append(bits, encodeAlphanumericBits(group, len(group)))
	}

	return bits, nil
}

// CharCount returns the number of characters in the input string.
// In Alphanumeric mode, character count is defined as the number of
// symbols in the allowed 45-character set.
func (ae *AlphanumericEncoder) CharCount() int {
	return len(ae.s)
}
