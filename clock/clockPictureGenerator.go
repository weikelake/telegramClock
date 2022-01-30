package clock

import (
	"fmt"
	"github.com/fogleman/gg"
	"time"
)

func GenerateClockPicture() {
	var (
		width  = 1000
		height = 1000
	)
	dc := gg.NewContext(width, height)
	dc.SetHexColor("#000000") // choose colour
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.Fill()
	dc.SetRGB(0, 100, 0)
	dc.LoadFontFace("./clock/digital-7.ttf", 380)
	h := time.Now()
	dc.DrawStringAnchored(fmt.Sprintf("%02d:%02d", h.Hour(), h.Minute()), float64(width/2), float64(height/2), 0.5, 0.5)

	dc.SavePNG("clock/out.png")
}
