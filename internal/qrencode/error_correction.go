package qrencode

import (
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

func GroupDataCodewords(
	ecLevel qrconst.ErrorCorrectionLevel,
	version int,
	dataCodewords []string,
) ([][][]string, error) {
	ecBlock := tables.ECBlocks[ecLevel][version-1]

	groups := make([][][]string, 0, 2)

	// Fill group 1
	group1 := make([][]string, ecBlock.Group1Blocks)
	start := 0
	for block := range ecBlock.Group1Blocks {
		end := start + ecBlock.Group1DataCodewords
		group1[block] = append([]string{}, dataCodewords[start:end]...)
		start = end
	}
	groups = append(groups, group1)

	// Fill group 2 (if present)
	if ecBlock.Group2Blocks > 0 {
		group2 := make([][]string, ecBlock.Group2Blocks)
		for block := range ecBlock.Group2Blocks {
			end := start + ecBlock.Group2DataCodewords
			group2[block] = append([]string{}, dataCodewords[start:end]...)
			start = end
		}
		groups = append(groups, group2)
	}

	return groups, nil
}
