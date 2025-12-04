package qrencode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

func DetermineVersion(
	encMode qrconst.EncodingMode,
	ecLevel qrconst.ErrorCorrectionLevel,
	charCount int,
) (int, error) {
	switch ecLevel {
	case qrconst.L, qrconst.M, qrconst.Q, qrconst.H:
		charCapacity := tables.CharacterCapacities[encMode][ecLevel]

		for version := range 40 {
			lowCharacterCapacity := charCapacity[version]
			highCharacterCapacity := charCapacity[40-version-1]

			if lowCharacterCapacity >= charCount {
				return version + 1, nil
			}
			if highCharacterCapacity < charCount {
				if version != 0 {
					return 40 - version + 1, nil
				} else {
					return 0, fmt.Errorf("no version found that can encode %d characters", charCount)
				}
			}
		}
		return 0, fmt.Errorf("no version found that can encode %d characters", charCount)

	default:
		return 0, fmt.Errorf("invalid error correction level")
	}
}

func ModeIndicator(encMode qrconst.EncodingMode) string {
	s := strconv.FormatInt(int64(encMode), 2)
	return padBitString(s, 4)
}

func CharCountIndicator(
	encMode qrconst.EncodingMode,
	version int,
	charCount int,
) string {
	var idx int

	if version >= 1 && version <= 9 {
		idx = 0
	} else if version >= 10 && version <= 26 {
		idx = 1
	} else if version >= 27 && version <= 40 {
		idx = 2
	}

	bits := tables.CharacterCountIndicatorBits[encMode][idx]
	b := strconv.FormatInt(int64(charCount), 2)

	return padBitString(b, bits)
}

func AssembleDataCodewords(
	ecLevel qrconst.ErrorCorrectionLevel,
	version int,
	bitStrings []string,
) ([]string, error) {
	ecBlockInfo := tables.ECBlockInfos[ecLevel][version-1]
	totalDataCodewords := ecBlockInfo.Group1Blocks*ecBlockInfo.Group1DataCodewordsPerBlock + ecBlockInfo.Group2Blocks*ecBlockInfo.Group2DataCodewordsPerBlock

	var sb strings.Builder
	for _, bitString := range bitStrings {
		sb.WriteString(bitString)
	}

	// Add a terminator of 0s (if necessary)
	terminatorLength := min(4, totalDataCodewords*8-sb.Len())
	if terminatorLength < 0 {
		return nil, fmt.Errorf("input bits exceed data capacity")
	}
	sb.WriteString(strings.Repeat("0", terminatorLength))

	// Add more 0s to make the length of the bit string
	// a multiple of 8
	remainderLength := (8 - sb.Len()%8) % 8
	sb.WriteString(strings.Repeat("0", remainderLength))

	// Add pad bytes if the bit string is still too short
	if sb.Len() < totalDataCodewords*8 {
		pads := []string{"11101100", "00010001"}

		totalPadBytes := (totalDataCodewords*8 - sb.Len()) / 8
		for i := range totalPadBytes {
			sb.WriteString(pads[i%2])
		}
	}

	// Group bit string into 8-bit codewords
	var dataCodewords []string
	for i := range totalDataCodewords {
		dataCodewords = append(dataCodewords, sb.String()[i*8:i*8+8])
	}

	return dataCodewords, nil
}
