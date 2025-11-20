package qrencode

import (
	"fmt"

	"golang.org/x/text/encoding/japanese"
)

var sjisEncoder = japanese.ShiftJIS.NewEncoder()

type KanjiEncoder struct {
	s string
}

func NewKanjiEncoder(s string) *KanjiEncoder {
	return &KanjiEncoder{
		s: s,
	}
}

// Convert a kanji rune into an integer.
func kanjiRuneToInt(r rune) (int, error) {
	sjisBytes, err := sjisEncoder.Bytes([]byte(string(r)))
	if err != nil {
		return 0, err
	}

	sjisValue := uint16(sjisBytes[0])<<8 | uint16(sjisBytes[1])

	var base uint16
	if sjisValue >= 0x8140 && sjisValue <= 0x9FFC {
		base = 0x8140
	}
	if sjisValue >= 0xE040 && sjisValue <= 0xEBBF {
		base = 0xC140
	}

	sjisValue -= base
	msb := (sjisValue >> 8) & 0x00FF
	lsb := sjisValue & 0x00FF

	return int(msb*0x00C0 + lsb), nil
}

func (ke *KanjiEncoder) Encode() ([]string, error) {
	var bits []string
	for _, kanjiRune := range ke.s {
		kanjiInt, err := kanjiRuneToInt(kanjiRune)
		if err != nil {
			return nil, err
		}

		bits = append(bits, fmt.Sprintf("%013b", kanjiInt))
	}

	return bits, nil
}
