package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrcode/render"
	"github.com/ahmadnaufalhakim/qrgen/internal/qrconst"
)

var (
	BLACK  = color.RGBA{0, 0, 0, 255}
	WHITE  = color.RGBA{255, 255, 255, 255}
	PINK   = color.RGBA{224, 134, 200, 255}
	YELLOW = color.RGBA{251, 208, 67, 255}
)

func main() {

	// text := "8675309" //Numeric
	// text := "HELLO WORLD" //Alphanumeric
	// text := "だから僕は音楽をやめた" //Kanji
	text := "stop making quotes i never said\n\n- Sun Tzu,\n  The Art of War" //Byte
	// text := "What the fuck did you just fucking say about me, you little bitch? I’ll have you know I graduated top of my class in the Navy Seals, and I’ve been involved in numerous secret raids on Al-Quaeda, and I have over 300 confirmed kills. I am trained in gorilla warfare and I’m the top sniper in the entire US armed forces. You are nothing to me but just another target. I will wipe you the fuck out with precision the likes of which has never been seen before on this Earth, mark my fucking words. You think you can get away with saying that shit to me over the Internet? Think again, fucker. As we speak I am contacting my secret network of spies across the USA and your IP is being traced right now so you better prepare for the storm, maggot. The storm that wipes out the pathetic little thing you call your “life”. You’re fucking dead, kid. I can be anywhere, anytime, and I can kill you in over seven hundred ways, and that’s just with my bare hands. Not only am I extensively trained in unarmed combat, but I have access to the entire arsenal of the United States Marine Corps and I will use it to its full extent to wipe your miserable ass off the face of the continent, you little shit. If only you could have known what unholy retribution your little “clever” comment was about to bring down upon you, maybe you would have held your fucking tongue. But you couldn’t, you didn’t, and now you’re paying the price, you goddamn idiot. I will shit fury all over you and you will drown in it. You’re fucking dead, kiddo." //Byte

	qrBuilder := qrcode.NewQRBuilder(text)
	qrCode, err := qrBuilder.
		WithErrorCorrectionLevel(qrconst.L).
		Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	bg := WHITE
	fg := BLACK

	start := time.Now()
	qrRenderer := render.NewRenderer().
		WithModuleShape(qrconst.Pointillism).
		WithBackgroundColor(bg).
		WithForegroundColor(fg).
		WithKernelType(qrconst.KernelLanczos2)

	err = qrRenderer.RenderPNG(*qrCode, "main.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Rendering took %s\n", elapsed)
}
