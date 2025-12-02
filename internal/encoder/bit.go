package encoder

func padBitString(s string, width int) string {
	if len(s) >= width {
		return s
	}

	pad := width - len(s)
	buf := make([]byte, width)
	for i := range pad {
		buf[i] = '0'
	}
	copy(buf[pad:], s)

	return string(buf)
}
