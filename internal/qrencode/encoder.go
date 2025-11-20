package qrencode

type Encoder interface {
	Encode() []string
}

func NewEncoder(s string) Encoder {
	encodingMode := DetermineEncodingMode(s)

	switch encodingMode {
	case NumericMode:
		return NewNumericEncoder(s)

	case AlphanumericMode:
		return NewAlphanumericEncoder(s)

		// case KanjiMode:
		// 	return NewKanjiEncoder(s)

		// case ByteMode:
		// 	return NewByteEncoder(s)

		// case ECIMode:
		// 	return NewECIEncoder(s)
	}

	return nil
}
