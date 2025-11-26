package qrcode

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
	"github.com/ahmadnaufalhakim/qrgen/internal/tables"
)

type QRCode struct {
	Version  int
	ECLevel  qrconst.ErrorCorrectionLevel
	Size     int
	Modules  [][]bool
	Patterns [][]qrconst.FunctionPattern
}

func NewQRCode(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	messageBitString string,
) *QRCode {
	size := (version-1)*4 + 21
	modules := make([][]bool, size)
	patterns := make([][]qrconst.FunctionPattern, size)
	for i := range size {
		modules[i] = make([]bool, size)
		patterns[i] = make([]qrconst.FunctionPattern, size)
	}

	qrcode := &QRCode{
		Version:  version,
		ECLevel:  ecLevel,
		Size:     size,
		Modules:  modules,
		Patterns: patterns,
	}

	qrcode.PlaceFinderPatterns()
	qrcode.PlaceSeparators()
	qrcode.PlaceAlignmentPattern()
	qrcode.PlaceTimingPattern()
	qrcode.PlaceDarkModule()
	qrcode.ReserveFormatInformationArea()
	qrcode.ReserveVersionInformationArea()

	return qrcode
}

func (qr *QRCode) PlaceFinderPatterns() {
	startPos := [3][2]int{
		{0, 0},
		{0, qr.Size - 7},
		{qr.Size - 7, 0},
	}

	for _, pos := range startPos {
		for i := range 7 {
			for j := range 7 {
				isOuter := (i == 0 || i == 6) || (j == 0 || j == 6)
				isInner := (i >= 2 && i <= 4) && (j >= 2 && j <= 4)

				var module bool
				if isOuter || isInner {
					module = true
				} else {
					module = false
				}

				qr.Modules[pos[0]+i][pos[1]+j] = module
				qr.Patterns[pos[0]+i][pos[1]+j] = qrconst.FPFinder
			}
		}
	}
}

func (qr *QRCode) PlaceSeparators() {
	startPos := [3][2]int{
		{0, 0},
		{0, qr.Size - 7},
		{qr.Size - 7, 0},
	}

	for i, j := 0, 0; i < 8 && j < 8; i, j = i+1, j+1 {
		// Top left
		qr.Modules[startPos[0][0]+i][startPos[0][1]+7] = false
		qr.Modules[startPos[0][0]+7][startPos[0][1]+j] = false
		qr.Patterns[startPos[0][0]+i][startPos[0][1]+7] = qrconst.FPSeparator
		qr.Patterns[startPos[0][0]+7][startPos[0][1]+j] = qrconst.FPSeparator

		// Top right
		qr.Modules[startPos[1][0]+i][startPos[1][1]-1] = false
		qr.Modules[startPos[1][0]+7][startPos[1][1]-1+j] = false
		qr.Patterns[startPos[1][0]+i][startPos[1][1]-1] = qrconst.FPSeparator
		qr.Patterns[startPos[1][0]+7][startPos[1][1]-1+j] = qrconst.FPSeparator

		// Bottom left
		qr.Modules[startPos[2][0]-1+i][startPos[2][1]+7] = false
		qr.Modules[startPos[2][0]-1][startPos[2][1]+j] = false
		qr.Patterns[startPos[2][0]-1+i][startPos[2][1]+7] = qrconst.FPSeparator
		qr.Patterns[startPos[2][0]-1][startPos[2][1]+j] = qrconst.FPSeparator
	}
}

func (qr *QRCode) PlaceAlignmentPattern() {
	version := qr.Version
	alignmentRowCols := tables.AlignmentPatternLocations[version-1]
	alignmentCenterPos := make([][2]int, len(alignmentRowCols)*len(alignmentRowCols))
	for i, r := range alignmentRowCols {
		for j, c := range alignmentRowCols {
			alignmentCenterPos[i*len(alignmentRowCols)+j] = [2]int{r, c}
		}
	}

	for _, centerPos := range alignmentCenterPos {
		if qr.Patterns[centerPos[0]][centerPos[1]] != qrconst.FPUnoccupied {
			continue
		}

		for i := range 5 {
			for j := range 5 {
				isOuter := (i == 0 || i == 4) || (j == 0 || j == 4)
				isInner := i == 2 && j == 2

				var module bool
				if isOuter || isInner {
					module = true
				} else {
					module = false
				}

				qr.Modules[centerPos[0]-2+i][centerPos[1]-2+j] = module
				qr.Patterns[centerPos[0]-2+i][centerPos[1]-2+j] = qrconst.FPAlignment
			}
		}
	}

}

func (qr *QRCode) PlaceTimingPattern() {
	startPos := [3][2]int{
		{0, 0},
		{0, qr.Size - 7},
		{qr.Size - 7, 0},
	}

	module := true
	// Place vertical timing pattern if not occupied by another function pattern
	for i := startPos[0][0] + 6; i < startPos[2][0]; i++ {
		if qr.Patterns[i][6].IsUnoccupied() {
			qr.Modules[i][6] = module
			qr.Patterns[i][6] = qrconst.FPTiming
		}
		module = !module
	}
	// Place horizontal timing pattern if not occupied by another function pattern
	for j := startPos[0][0] + 6; j < startPos[2][0]; j++ {
		if qr.Patterns[6][j].IsUnoccupied() {
			qr.Modules[6][j] = module
			qr.Patterns[6][j] = qrconst.FPTiming
		}
		module = !module
	}
}

func (qr *QRCode) PlaceDarkModule() {
	version := qr.Version
	qr.Modules[4*version+9][8] = true
	qr.Patterns[4*version+9][8] = qrconst.FPDarkModule
}

func (qr *QRCode) ReserveFormatInformationArea() {
	startPos := [3][2]int{
		{0, 0},
		{0, qr.Size - 7},
		{qr.Size - 7, 0},
	}

	for i, j := 0, 0; i < 9 && j < 9; i, j = i+1, j+1 {
		// Reserve two-module strip near the top-left finder pattern
		if qr.Patterns[startPos[0][0]+i][startPos[0][1]+8].IsUnoccupied() {
			qr.Patterns[startPos[0][0]+i][startPos[0][1]+8] = qrconst.FPFormatInfo
		}
		if qr.Patterns[startPos[0][0]+8][startPos[0][1]+j].IsUnoccupied() {
			qr.Patterns[startPos[0][0]+8][startPos[0][1]+j] = qrconst.FPFormatInfo
		}

		// Reserve one-module strip near the top-right finder pattern
		if qr.Patterns[startPos[1][0]+8][startPos[1][1]-1+j].IsUnoccupied() {
			qr.Patterns[startPos[1][0]+8][startPos[1][1]-1+j] = qrconst.FPFormatInfo
		}

		// Reserve one-module strip near the bottom-left finder pattern
		if qr.Patterns[startPos[2][0]-1+i][startPos[2][1]+8].IsUnoccupied() {
			qr.Patterns[startPos[2][0]-1+i][startPos[2][1]+8] = qrconst.FPSeparator
		}
	}
}

func (qr *QRCode) ReserveVersionInformationArea() {
	if qr.Version < 7 {
		return
	}

	startPos := [3][2]int{
		{0, qr.Size - 7},
		{qr.Size - 7, 0},
	}

	// Reserve 6x3 block to the left of the top-right finder pattern
	for i := range 6 {
		for j := range 3 {
			qr.Patterns[startPos[0][0]+i][startPos[0][1]-4+j] = qrconst.FPVersionInfo
		}
	}

	// Reserve 3x6 block above the bottom-left finder pattern
	for i := range 3 {
		for j := range 6 {
			qr.Patterns[startPos[1][0]-4+i][startPos[1][1]+j] = qrconst.FPVersionInfo
		}
	}
}

func (qr *QRCode) RenderPNG(filename string, scale int) error {
	if scale < 1 {
		scale = 1
	}

	imgSize := qr.Size * scale
	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	for y := range qr.Size {
		for x := range qr.Size {
			var c color.RGBA
			if qr.Modules[y][x] {
				c = black
			} else {
				c = white
			}

			// Fill the scale*scale block
			startX, startY := x*scale, y*scale
			for dy := range scale {
				for dx := range scale {
					img.Set(startX+dx, startY+dy, c)
				}
			}
		}
	}

	// Create output file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode PNG
	return png.Encode(f, img)
}
