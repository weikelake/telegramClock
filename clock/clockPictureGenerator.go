package clock

import (
	"fmt"
	"github.com/fogleman/gg"
	"telegramClock/settings"
	"time"
)

func GenerateClockPicture(offsetMinute int) {
	var (
		width  = 1000
		height = 1000
	)
	dc := gg.NewContext(width, height)
	dc.SetHexColor("#000000") // choose colour
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.Fill()
	dc.SetRGB(0, 100, 0)
	err := dc.LoadFontFace("./clock/digital-7.ttf", 380)
	if err != nil {
		fmt.Println(err)
	}
	h := time.Now()
	h = h.Add(time.Duration(offsetMinute) * time.Minute)
	dc.DrawStringAnchored(fmt.Sprintf("%02d:%02d", h.Hour(), h.Minute()), float64(width/2), float64(height/2), 0.5, 0.5)

	err = dc.SavePNG(settings.GetPicturePath())
	if err != nil {
		fmt.Println(err)
	}
}
