package tcolor

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Printf("%s end\n", Colorize("test", New().ForegroundRGB(5, 0, 0).Bold().Underline().BackgroundRGB(0, 0, 5)))
	fmt.Printf("%stest%s end\n", New().ForegroundRGB(5, 0, 0), Reset)
	fmt.Println(Colorize("ForegroundRGB grayscale test", New().ForegroundGray(13).Underline()))
	fmt.Println(Colorize("Background grayscale test", New().BackgroundGray(5)))
	fmt.Println(Colorize("Standard color test", New().Foreground(Red).Background(Blue)))
	fmt.Println(Colorize("Standard bright color test", New().Foreground(BrightRed).Background(BrightBlue)))
	fmt.Println(Colorize("Negative test", New().Negative().Foreground(Red)))

	ColorOn = false
	fmt.Println(Colorize("Color is off", New().Foreground(Red).Underline()))
	ColorOn = true
	fmt.Println(Colorize("Color is on", New().Foreground(Red).Underline()))

}
