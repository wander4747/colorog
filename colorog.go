package colorog

import (
	"log"

	"github.com/wander4747/colorog/color"
)

func Success(text string) {
	print(color.Green, text)
}

func Info(text string) {
	print(color.Blue, text)
}

func Warning(text string) {
	print(color.Yellow, text)
}

func Danger(text string) {
	print(color.Red, text)
}

func Light(text string) {
	print(color.White, text)
}

func Fatal(text string) {
	log.Fatal(color.Red + text + color.Reset)
}

func WithColor(color, text string) {
	print(color, text)
}

func Unicorn(text string) {
	unicornText := ""

	for i, char := range text {
		currentColor := i % len(color.UnicornColors)
		unicornText = unicornText + colorize(color.UnicornColors[currentColor], string(char))
	}

	log.Print(unicornText)
}

func colorize(cl, text string) string {
	return cl + text + color.Reset
}

func print(color, text string) {
	log.Print(colorize(color, text))
}
