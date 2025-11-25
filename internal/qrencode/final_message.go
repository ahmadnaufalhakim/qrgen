package qrencode

import (
	"fmt"
	"strings"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

func StructureFinalMessage(
	ecLevel qrconst.ErrorCorrectionLevel,
	version int,
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
