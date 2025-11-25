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

func byteToBitString(b uint8) string {
	buf := [8]uint8{}
	for i := range 8 {
		// Check bit from lrft
		if b&(1<<(7-i)) != 0 {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}

	return string(buf[:])
}
