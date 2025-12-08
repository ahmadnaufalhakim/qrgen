package main

import (
	"fmt"
	"image/color"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode/render"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

func main() {

	// text := "8675309" //Numeric
	// text := "HELLO WORLD" //Alphanumeric
	// text := "だから僕は音楽をやめた" //Kanji
	text := "joe mama\n\n- Sun Tzu,\n  The Art of War" //Byte

	qrBuilder := qrcode.NewQRBuilder(text)
	qrCode, err := qrBuilder.
		WithErrorCorrectionLevel(qrconst.L).
		Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	qrRenderer := render.NewRenderer().
		WithModuleShape(qrconst.Blob).
		WithBackgroundColor(color.RGBA{255, 255, 255, 255}).
		WithForegroundColor(color.RGBA{0, 0, 0, 255})

	err = qrRenderer.RenderPNG(*qrCode, "main.png")
	if err != nil {
		fmt.Println(err)
		return
	}
}
