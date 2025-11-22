package qrencode

import (
	"fmt"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

type NumericEncoder struct {
	s string
}

func NewNumericEncoder(s string) *NumericEncoder {
	return &NumericEncoder{
		s: s,
	}
}

// Convert a numeric string of length 1,2,3 into an integer.
func numericStrToInt(s string) int {
	result := 0
	for _, r := range s {
		result = (result * 10) + int(r-'0')
	}

	return result
}

// Split numeric string into groups of 3 (as per QR code specification).
func numericSplit(s string) []string {
	var groups []string
	for i := 0; i < len(s); i += 3 {
		end := min(i+3, len(s))
		groups = append(groups, s[i:end])
	}

	return groups
}

// Convert one numeric string group into its QR bit representation.
// 3 digits -> 10 bits, 2 digits -> 7 bits, 1 digit -> 4 bits
func encodeNumericBits(group string, groupSize int) string {
	var binaryFormat string
	switch groupSize {
	case 3:
		binaryFormat = "%010b"
	case 2:
		binaryFormat = "%07b"
	case 1:
		binaryFormat = "%04b"
	}

	return fmt.Sprintf(binaryFormat, numericStrToInt(group))
}

// Encode encodes the numeric string into a slice of bit strings,
// following QR Code Numeric Mode rules.
//
// The string is split into groups of 1–3 digits (as required by
// the QR specification). Each group is encoded into:
//   - 10 bits for 3 digits
//   - 7 bits  for 2 digits
//   - 4 bits  for 1 digit
//
// The returned slice contains each group’s bit representation.
func (ne *NumericEncoder) Encode() ([]string, error) {
	var bitStrings []string
	for _, group := range numericSplit(ne.s) {
		bitStrings = append(bitStrings, encodeNumericBits(group, len(group)))
	}

	return bitStrings, nil
}

// CharCount returns the number of characters in the numeric string.
// In Numeric mode, the QR spec defines the "character count" simply as
// the number of digits.
func (ne *NumericEncoder) CharCount() int {
	return len(ne.s)
}

// Mode returns numeric mode EncodingMode
func (ne *NumericEncoder) Mode() qrconst.EncodingMode {
	return qrconst.NumericMode
}
