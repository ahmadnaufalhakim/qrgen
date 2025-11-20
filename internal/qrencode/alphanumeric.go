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

func (ae *AlphanumericEncoder) Encode() ([]string, error) {
	var bits []string
	for _, group := range alphanumericSplit(ae.s) {
		bits = append(bits, encodeAlphanumericBits(group, len(group)))
	}

	return bits, nil
}
