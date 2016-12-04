package tcolor

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Printf("%s end\n", Colorize("test", DefaultColor().ForegroundRGB(5, 0, 0).Bold().Underline().BackgroundRGB(0, 0, 5)))
	fmt.Printf("%stest%s end\n", DefaultColor().ForegroundRGB(5, 0, 0), Reset)
	fmt.Println(Colorize("ForegroundRGB grayscale test", DefaultColor().ForegroundGray(13).Underline()))
	fmt.Println(Colorize("Background grayscale test", DefaultColor().BackgroundGray(5)))
	fmt.Println(Colorize("Standard color test", DefaultColor().Foreground(Red).Background(Blue)))
	fmt.Println(Colorize("Standard bright color test", DefaultColor().Foreground(BrightRed).Background(BrightBlue)))

	ColorOn = false
	fmt.Println(Colorize("Color is off", DefaultColor().Foreground(Red).Underline()))
	ColorOn = true
	fmt.Println(Colorize("Color is on", DefaultColor().Foreground(Red).Underline()))
}
