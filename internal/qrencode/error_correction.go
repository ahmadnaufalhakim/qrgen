package qrencode

import (
	"fmt"
	"strings"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

// AssembleDataBlocks splits the data codewords into the exact block structure
// required by the QR Code specification for the given version and error
// correction level.
//
// A QR version may have:
// - Group 1 blocks, each containing N1 data codewords; and
// - Group 2 blocks, each containing N2 data codewords (usually N1 + 1)
//
// The function returns a flat slice of data blocks, where each block is
// a slice of data codewords (8-bit strings).
func AssembleDataBlocks(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	dataCodewords []string,
) ([][]string, error) {
	ecBlockInfo := tables.ECBlockInfos[ecLevel][version-1]
	group1Blocks := ecBlockInfo.Group1Blocks
	group2Blocks := ecBlockInfo.Group2Blocks
	n1 := ecBlockInfo.Group1DataCodewordsPerBlock
	n2 := ecBlockInfo.Group1DataCodewordsPerBlock

	dataBlocks := make([][]string, group1Blocks+group2Blocks)

	start := 0

	// Fill with group 1 blocks
	for i := range group1Blocks {
		end := start + n1
		if end > len(dataCodewords) {
			return nil, fmt.Errorf("insufficient data codewords for Group 1 blocks")
		}

		dataBlocks[i] = append([]string{}, dataCodewords[start:end]...)
		start = end
	}

	// Fill with group 2 blocks (if any)
	for i := range group2Blocks {
		end := start + n2
		if end > len(dataCodewords) {
			return nil, fmt.Errorf("insufficient data codewords for Group 2 blocks")
		}

		dataBlocks[group1Blocks+i] = append([]string{}, dataCodewords[start:end]...)
		start = end
	}

	// Safety check (should always match)
	if start != len(dataCodewords) {
		return nil, fmt.Errorf("data codeword count mismatch: leftover = %d", len(dataCodewords)-start)
	}

	return dataBlocks, nil
}

// MessagePolynomial converts a slice of 8-bit data codewords (as strings)
// into a slice of uint8 representing the message polynomial for error correction.
func MessagePolynomial(dataBlock []string) ([]uint8, error) {
	var m []uint8
	for _, dataCodeword := range dataBlock {
		b, err := bitStringToByte(dataCodeword)
		if err != nil {
			return nil, err
		}
		m = append(m, b)
	}

	return m, nil
}

// GeneratorPolynomial returns the generator polynomial of degree n for
// Reed-Solomon error correction, computing it on-demand if not stored
// in-memory.
func GeneratorPolynomial(n int) ([]uint8, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n must be positive")
	}

	if g, ok := tables.GeneratorPolynomial[n]; ok {
		return g, nil
	}

	var polynomials [][]uint8
	for i := n; i > 0; i-- {
		if g, ok := tables.GeneratorPolynomial[i]; ok {
			polynomials = append(polynomials, g)
			break
		} else {
			polynomials = append(polynomials, []uint8{1, tables.AntilogGF256[i-1]})
		}
	}

	prod, err := polynomialsProduct(polynomials...)
	if err != nil {
		return nil, err
	}

	tables.GeneratorPolynomial[n] = prod

	return prod, nil
}

// GenerateErrorCorrectionBlocks computes the error correction codewords
// for each data block of a QR Code using the appropriate generator polynomial.
func GenerateErrorCorrectionBlocks(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	dataBlocks [][]string,
) ([][]uint8, error) {
	ecBlockInfo := tables.ECBlockInfos[ecLevel][version-1]
	n := ecBlockInfo.ECCodewordsPerBlock
	g, err := GeneratorPolynomial(n)
	if err != nil {
		return nil, err
	}

	ecBlocks := make([][]uint8, ecBlockInfo.Group1Blocks+ecBlockInfo.Group2Blocks)
	for i, dataBlock := range dataBlocks {
		m, err := MessagePolynomial(dataBlock)
		if err != nil {
			return nil, err
		}

		ecCodewords := divideTwoPolynomials(m, g)
		ecBlocks[i] = ecCodewords
	}

	return ecBlocks, nil
}

// InterleaveBlocks interleaves the data and error correction codewords
// according to the QR Code specification and returns the final message
// bit string representing the full message ready for placement in the
// QR matrix.
//
// It handles both group 1 and group 2 blocks, performs column-wise
// interleaving, appends the error correction codewords, and adds the
// remainder bits for the given QR version.
func InterleaveBlocks(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	dataBlocks [][]string,
	ecBlocks [][]uint8,
) (string, error) {
	if len(dataBlocks) != len(ecBlocks) {
		return "", fmt.Errorf("number of data blocks must be equal to the number of error correction blocks")
	}

	ecBlockInfo := tables.ECBlockInfos[ecLevel][version-1]
	ecCodewordsPerBlock := ecBlockInfo.ECCodewordsPerBlock
	group1Blocks := ecBlockInfo.Group1Blocks
	group2Blocks := ecBlockInfo.Group2Blocks
	group1DataCodewordsPerBlock := ecBlockInfo.Group1DataCodewordsPerBlock
	group2DataCodewordsPerBlock := ecBlockInfo.Group2DataCodewordsPerBlock
	group1DataCodewords := group1Blocks * group1DataCodewordsPerBlock
	group2DataCodewords := group2Blocks * group2DataCodewordsPerBlock

	// Interleave the data codewords
	totalDataCodewords := group1DataCodewords + group2DataCodewords
	interleavedDataCodewords := make([]uint8, totalDataCodewords)
	dataCols := max(group1DataCodewordsPerBlock, group2DataCodewordsPerBlock)
	group1Offset := group1Blocks
	for j := range dataCols {
		if j < group1DataCodewordsPerBlock {
			for i := range group1Blocks {
				b, err := bitStringToByte(dataBlocks[i][j])
				if err != nil {
					return "", err
				}

				interleavedDataCodewords[(group1Blocks+group2Blocks)*j+i] = b
			}
			group1Offset = group1Blocks
		} else {
			group1Offset = 0
		}

		if j < group2DataCodewordsPerBlock {
			for i := range group2Blocks {
				b, err := bitStringToByte(dataBlocks[group1Blocks+i][j])
				if err != nil {
					return "", err
				}

				interleavedDataCodewords[(group1Blocks+group2Blocks)*j+group1Offset+i] = b
			}
		}
	}

	// Interleave the error correction codewords
	totalECCodewords := (group1Blocks + group2Blocks) * ecCodewordsPerBlock
	interleavedECCodewords := make([]uint8, totalECCodewords)
	ecCols := ecCodewordsPerBlock
	for j := range ecCols {
		for i := range group1Blocks + group2Blocks {
			interleavedECCodewords[(group1Blocks+group2Blocks)*j+i] = ecBlocks[i][j]
		}
	}

	// Build the final message string
	var finalMessageBuilder strings.Builder
	for _, dataCodeword := range interleavedDataCodewords {
		finalMessageBuilder.WriteString(byteToBitString(dataCodeword))
	}
	for _, ecCodeword := range interleavedECCodewords {
		finalMessageBuilder.WriteString(byteToBitString(ecCodeword))
	}
	finalMessageBuilder.WriteString(strings.Repeat("0", tables.RemainderBits[version-1]))

	return finalMessageBuilder.String(), nil
}
