package qrconst

type Lookahead uint8

const (
	LookR  Lookahead = 1 << 7
	LookUR Lookahead = 1 << 6
	LookU  Lookahead = 1 << 5
	LookUL Lookahead = 1 << 4
	LookL  Lookahead = 1 << 3
	LookDL Lookahead = 1 << 2
	LookD  Lookahead = 1 << 1
	LookDR Lookahead = 1 << 0
)
