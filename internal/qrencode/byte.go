package qrencode

import "fmt"

type ByteEncoder struct {
	s string
}

func NewByteEncoder(s string) *ByteEncoder {
	return &ByteEncoder{
		s: s,
	}
}

func (be *ByteEncoder) Encode() ([]string, error) {
	var bits []string
	for i := 0; i < len(be.s); i++ {
		bits = append(bits, fmt.Sprintf("%08b", be.s[i]))
	}

	return bits, nil
}
