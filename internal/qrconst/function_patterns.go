package qrconst

type FunctionPattern int

const (
	FPUnoccupied FunctionPattern = iota
	FPFinder
	FPSeparator
	FPAlignment
	FPTiming
	FPDarkModule
	FPFormatInfo
	FPVersionInfo
	FPMessageBit
)

func (fp FunctionPattern) IsOccupied() bool {
	return fp != FPUnoccupied
}

func (fp FunctionPattern) IsUnoccupied() bool {
	return fp == FPUnoccupied
}

func (fp FunctionPattern) IsMessage() bool {
	return fp == FPMessageBit
}

func (fp FunctionPattern) IsFormat() bool {
	return fp == FPFormatInfo
}

func (fp FunctionPattern) IsVersion() bool {
	return fp == FPVersionInfo
}

func (fp FunctionPattern) IsDarkModule() bool {
	return fp == FPDarkModule
}

func (fp FunctionPattern) IsTiming() bool {
	return fp == FPTiming
}
