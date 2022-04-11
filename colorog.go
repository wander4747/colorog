package colorog

import (
	"log"

	"github.com/wander4747/colorog/color"
)

type Colorog struct {
}

func New() *Colorog {
	return &Colorog{}
}

func (c Colorog) Success(text string) {
	c.print(color.Green, text)
}

func (c Colorog) Info(text string) {
	c.print(color.Blue, text)
}

func (c Colorog) Warning(text string) {
	c.print(color.Yellow, text)
}

func (c Colorog) Danger(text string) {
	c.print(color.Red, text)
}

func (c Colorog) Light(text string) {
	c.print(color.White, text)
}

func (c Colorog) Fatal(text string) {
	log.Fatal(color.Red + text + color.Reset)
}

func (c Colorog) WithColor(color, text string) {
	c.print(color, text)
}

func (c Colorog) Unicorn(text string) {
	unicornText := ""

	for i, char := range text {
		currentColor := i % len(color.UnicornColors)
		unicornText = unicornText + c.colorize(color.UnicornColors[currentColor], string(char))
	}

	log.Print(unicornText)
}

func (c Colorog) colorize(cl, text string) string {
	return cl + text + color.Reset
}

func (c Colorog) print(color, text string) {
	log.Print(c.colorize(color, text))
}
