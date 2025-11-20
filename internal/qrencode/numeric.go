package qrencode

import (
	"fmt"
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

func (ne *NumericEncoder) Encode() ([]string, error) {
	var bits []string
	for _, group := range numericSplit(ne.s) {
		bits = append(bits, encodeNumericBits(group, len(group)))
	}

	return bits, nil
}
