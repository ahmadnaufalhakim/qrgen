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
	Function [][]bool
}

func NewQRCode(
	version int,
	ecLevel qrconst.ErrorCorrectionLevel,
	messageBitString string,
) *QRCode {
	size := (version-1)*4 + 21
	modules := make([][]bool, size)
	function := make([][]bool, size)
	for i := range size {
		modules[i] = make([]bool, size)
		function[i] = make([]bool, size)
	}

	qrcode := &QRCode{
		Version:  version,
		ECLevel:  ecLevel,
		Size:     size,
		Modules:  modules,
		Function: function,
	}

	qrcode.PlaceFinderPatterns()
	qrcode.PlaceSeparators()
	qrcode.PlaceAlignmentPattern()

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
				qr.Function[pos[0]+i][pos[1]+j] = true
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
		qr.Function[startPos[0][0]+i][startPos[0][1]+7] = true
		qr.Function[startPos[0][0]+7][startPos[0][1]+j] = true

		// Top right
		qr.Modules[startPos[1][0]+i][startPos[1][1]-1] = false
		qr.Modules[startPos[1][0]+7][startPos[1][1]-1+j] = false
		qr.Function[startPos[1][0]+i][startPos[1][1]-1] = true
		qr.Function[startPos[1][0]+7][startPos[1][1]-1+j] = true

		// Bottom left
		qr.Modules[startPos[2][0]-1+i][startPos[2][1]+7] = false
		qr.Modules[startPos[2][0]-1][startPos[2][1]+j] = false
		qr.Function[startPos[2][0]-1+i][startPos[2][1]+7] = true
		qr.Function[startPos[2][0]-1][startPos[2][1]+j] = true
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
		if qr.Function[centerPos[0]][centerPos[1]] {
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
				qr.Function[centerPos[0]-2+i][centerPos[1]-2+j] = true
			}
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
