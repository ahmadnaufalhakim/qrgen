package qrencode

import "fmt"

func bitStringToByte(s string) (uint8, error) {
	if len(s) != 8 {
		return 0, fmt.Errorf("must be 8 bits")
	}

	b := uint8(0)
	for i := range 8 {
		b <<= 1
		if s[i] == '1' {
			b |= 1
		} else if s[i] != '0' {
			return 0, fmt.Errorf("invalid bit")
		}
	}

	return b, nil
}
