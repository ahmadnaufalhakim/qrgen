package qrencode

import (
	"fmt"
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
	return fmt.Sprintf("%04b", encMode)
}

func CharCountIndicator(
	encMode qrconst.EncodingMode,
	version int,
	charCount int,
) string {
	var charCountFormat string
	var idx int

	if version >= 1 && version <= 9 {
		idx = 0
	} else if version >= 10 && version <= 26 {
		idx = 1
	} else if version >= 27 && version <= 40 {
		idx = 2
	}
	charCountFormat = fmt.Sprintf("%%0%db", tables.CharacterCountIndicatorBits[encMode][idx])

	return fmt.Sprintf(charCountFormat, charCount)
}

func AssembleDataCodewords(
	ecLevel qrconst.ErrorCorrectionLevel,
	version int,
	bitStrings []string,
) ([]string, error) {
	ecBlock := tables.ECBlocks[ecLevel][version-1]
	totalDataCodewords := ecBlock.Group1Blocks*ecBlock.Group1DataCodewords + ecBlock.Group2Blocks*ecBlock.Group2DataCodewords

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
		pads := []int{236, 17}

		totalPadBytes := (totalDataCodewords*8 - sb.Len()) / 8
		for i := range totalPadBytes {
			sb.WriteString(fmt.Sprintf("%08b", pads[i%2]))
		}
	}

	// Group bit string into 8-bit codewords
	var dataCodewords []string
	for i := range totalDataCodewords {
		dataCodewords = append(dataCodewords, sb.String()[i*8:i*8+8])
	}

	return dataCodewords, nil
}
