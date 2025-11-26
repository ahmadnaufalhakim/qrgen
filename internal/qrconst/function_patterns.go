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
)

func (fp FunctionPattern) IsOccupied() bool {
	return fp != FPUnoccupied
}

func (fp FunctionPattern) IsUnoccupied() bool {
	return fp == FPUnoccupied
}
