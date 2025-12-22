package qrconst

type Lookahead uint16

const (
	LookR  Lookahead = 1 << 0
	LookUR Lookahead = 1 << 1
	LookU  Lookahead = 1 << 2
	LookUL Lookahead = 1 << 3
	LookL  Lookahead = 1 << 4
	LookDL Lookahead = 1 << 5
	LookD  Lookahead = 1 << 6
	LookDR Lookahead = 1 << 7

	LookFinder     Lookahead = 1 << 8
	LookSeparator  Lookahead = 1 << 9
	LookAlignment  Lookahead = 1 << 10
	LookTiming     Lookahead = 1 << 11
	LookDarkModule Lookahead = 1 << 12

	LookStructural = LookFinder |
		LookSeparator |
		LookAlignment |
		LookTiming |
		LookDarkModule
)

func (l Lookahead) Has(las ...Lookahead) bool {
	for _, la := range las {
		if l&la == 0 {
			return false
		}
	}

	return true
}

func (l Lookahead) Lacks(las ...Lookahead) bool {
	for _, la := range las {
		if l&la != 0 {
			return false
		}
	}

	return true
}

func (l Lookahead) HasAny(las ...Lookahead) bool {
	for _, la := range las {
		if l&la != 0 {
			return true
		}
	}

	return false
}
