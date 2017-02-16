package tcolor

import "log"
import "strconv"

const Reset Color = "0"

const (
	Black byte = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

// ColorOn enables colorization.
var ColorOn = true

// Colorize returns the given string into the specified color.
// Colorize prepends the color sequence to the string and appends the reset color sequence.
func Colorize(s string, color Color) string {
	return color.String() + s + Reset.String()
}

type Color string

func New() Color {
	return Color("")
}

func (c Color) String() string {
	if len(c) == 0 || !ColorOn {
		return ""
	}
	return "\033[" + string(c) + "m"
}

func colorCode(r, g, b byte) string {
	if r > 5 || g > 5 || b > 5 {
		log.Panicln("RGB values must be between 0 and 5, but was:", r, g, b)
	}
	return strconv.Itoa(int(0x10 + 36*r + 6*g + b))
}

func grayscaleCode(x byte) string {
	if x > 23 {
		log.Panicln("Grayscale value must be between 0 and 23, but was:", x)
	}
	return strconv.Itoa(int(0xe8 + x))
}

func (c Color) append(s string) Color {
	if c == "" {
		return Color(s)
	} else {
		return c + Color(";"+s)
	}
}

func (c Color) Foreground(standardColor byte) Color {
	return c.append("38;5;" + strconv.Itoa(int(standardColor)))
}

func (c Color) ForegroundRGB(r, g, b byte) Color {
	return c.append("38;5;" + colorCode(r, g, b))
}

func (c Color) ForegroundGray(x byte) Color {
	return c.append("38;5;" + grayscaleCode(x))
}

func (c Color) Background(standardColor byte) Color {
	return c.append("48;5;" + strconv.Itoa(int(standardColor)))
}

func (c Color) BackgroundRGB(r, g, b byte) Color {
	return c.append("48;5;" + colorCode(r, g, b))
}

func (c Color) BackgroundGray(x byte) Color {
	return c.append("48;5;" + grayscaleCode(x))
}

func (c Color) Bold() Color {
	return c.append("1")
}

func (c Color) Faint() Color {
	return c.append("2")
}

func (c Color) Italic() Color {
	return c.append("3")
}

func (c Color) Underline() Color {
	return c.append("4")
}

func (c Color) Negative() Color {
	return c.append("7")
}
