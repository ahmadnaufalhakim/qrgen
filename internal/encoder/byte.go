package encoder

import (
	"strconv"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

type ByteEncoder struct {
	s string
}

func NewByteEncoder(s string) *ByteEncoder {
	return &ByteEncoder{
		s: s,
	}
}

// Encode encodes the input string using QR Code Byte Mode.
//
// Each byte of the UTF-8 input string is encoded into an 8-bit value,
// as required by the QR specification for Byte mode. Multi-byte UTF-8
// characters will produce multiple encoded bytes.
func (be *ByteEncoder) Encode() ([]string, error) {
	bitStrings := make([]string, len(be.s))
	for i := 0; i < len(be.s); i++ {
		b := strconv.FormatInt(int64(be.s[i]), 2)
		bitStrings[i] = padBitString(b, 8)
	}

	return bitStrings, nil
}

// CharCount returns the number of bytes in the UTF-8 string.
// In QR Byte mode, the "character count" is defined as the number of
// encoded bytes, not Unicode characters.
func (be *ByteEncoder) CharCount() int {
	return len([]byte(be.s))
}

// Mode returns byte mode EncodingMode
func (be *ByteEncoder) Mode() qrconst.EncodingMode {
	return qrconst.ByteMode
}
