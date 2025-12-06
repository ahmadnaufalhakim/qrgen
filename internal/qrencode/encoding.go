package qrencode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

// DetermineVersion selects the smallest QR version that can fit `charCount`
// characters for the given encoding mode and error correction level.
//
// It uses two pointers approach:
// - One pointer scans upward from `minVersion`.
// - The other pointer scans downward from version 40
//
// The function returns the chosen version number (1-40).
// Returns error if no version is sufficient.
func DetermineVersion(
	encMode qrconst.EncodingMode,
	minVersion int,
	ecLevel qrconst.ErrorCorrectionLevel,
	charCount int,
) (int, error) {
	switch ecLevel {
	case qrconst.L, qrconst.M, qrconst.Q, qrconst.H:
		// Validate minVersion
		if minVersion < 1 || minVersion > 40 {
			return 0, fmt.Errorf("minVersion %d is out of range (1-40)", minVersion)
		}
		total := 40 - (minVersion - 1)

		charCapacity := tables.CharacterCapacities[encMode][ecLevel]

		for i := range total {
			version := i + (minVersion - 1)
			lowCharacterCapacity := charCapacity[version]
			highCharacterCapacity := charCapacity[40-version-1]

			// Low pointer check
			if lowCharacterCapacity >= charCount {
				return version + 1, nil
			}

			// High pointer check
			if highCharacterCapacity < charCount {
				if version != 0 {
					return 40 - version + 1, nil
				} else {
					return 0, fmt.Errorf("no version >= %d can encode %d characters",
						minVersion, charCount)
				}
			}
		}
		return 0, fmt.Errorf("no version >= %d can encode %d characters",
			minVersion, charCount)

	default:
		return 0, fmt.Errorf("invalid error correction level")
	}
}

// ModeIndicator returns the 4-bit mode indicator corresponding to
// the QR encoding mode (Numeric, Alphanumeric, Byte, Kanji, etc.).
func ModeIndicator(encMode qrconst.EncodingMode) string {
	s := strconv.FormatInt(int64(encMode), 2)
	return padBitString(s, 4)
}

// CharCountIndicator returns the bit string representing the number
// of input characters for this QR segment.
//
// QR Code specifications require that the number of bits used for the
// character count indicator depends on both:
//
//  1. the encoding mode (numeric, alphanumeric, byte, etc.)
//
//  2. the QR Version group:
//
// -  Version  1–9   -> Group 0
//
// -  Version 10–26  -> Group 1
//
// -  Version 27–40  -> Group 2
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

// AssembleDataCodewords takes the encoded data bit strings (mode indicator,
// char count indicator, and data bits) and converts them into properly
// padded 8-bit codewords according to the QR Code encoding specification.
//
// Steps performed here are:
//
//  1. Concatenate all bit strings into one continuous bitstream.
//  2. Add a terminator of up to 4 zeros (or fewer if space is limited).
//  3. Pad with additional zeros to align the bitstream to an 8-bit boundary.
//  4. If still shorter than total data capacity, append alternating pad bytes:
//     11101100 (0xEC)
//     00010001 (0x11)
//  5. Split the final bitstream into 8-bit codewords.
//
// The resulting slice contains *data codewords only* (no error correction),
// which will later be split into blocks and processed further.
func AssembleDataCodewords(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
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
