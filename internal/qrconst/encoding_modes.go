package qrconst

type EncodingMode int

const (
	NumericMode      EncodingMode = 0b_0001
	AlphanumericMode EncodingMode = 0b_0010
	ByteMode         EncodingMode = 0b_0100
	KanjiMode        EncodingMode = 0b_1000
	ECIMode          EncodingMode = 0b_0111
)

func (em EncodingMode) String() string {
	switch em {
	case NumericMode:
		return "Numeric"
	case AlphanumericMode:
		return "Alphanumeric"
	case ByteMode:
		return "Byte"
	case KanjiMode:
		return "Kanji"
	case ECIMode:
		return "ECI"
	}
	return "Unknown"
}
