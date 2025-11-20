package qrconst

type EncodingMode int

const (
	NumericMode      EncodingMode = 0b_0001
	AlphanumericMode EncodingMode = 0b_0010
	ByteMode         EncodingMode = 0b_0100
	KanjiMode        EncodingMode = 0b_1000
	ECIMode          EncodingMode = 0b_0111
)
