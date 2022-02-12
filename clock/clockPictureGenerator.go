package clock

import (
	"fmt"
	"github.com/fogleman/gg"
	"telegramClock/settings"
	"time"
)

func GenerateClockPicture() {
	var (
		width  = 1000
		height = 1000
	)
	dc := gg.NewContext(width, height)
	dc.SetHexColor(settings.GetClockData().BackgroundColor) // choose colour
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.Fill()
	dc.SetHexColor(settings.GetClockData().TimeColor)
	err := dc.LoadFontFace("./clock/digital-7.ttf", 420)
	if err != nil {
		fmt.Println(err)
	}
	h := time.Now()
	h = h.Add(time.Duration(settings.GetClockData().OffsetTimeMinute) * time.Minute)
	h = h.Add(time.Duration(settings.GetClockData().OffsetTimeHour) * time.Hour)
	dc.DrawStringAnchored(fmt.Sprintf("%02d:%02d", h.Hour(), h.Minute()), float64(width/2), float64(height/2), 0.5, 0.5)

	err = dc.SavePNG(settings.GetPicturePath())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ok")
}
