package colorog

import (
	"log"
	"strings"

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

func (c Colorog) colorize(cl, text string) string {
	return cl + text + color.Reset
}

func (c Colorog) print(color, text string) {
	log.Print(c.colorize(color, text))
}

func (c Colorog) Unicorn(text string) {
	s := strings.Split(text, "")

	t := ""
	size := len(color.Unicorn)
	j := 0

	for i, s2 := range s {
		if i >= size && j == size {
			j = 0
		}
		ci := color.Unicorn[j]

		t += ci + s2
		j++

	}
	t += color.Reset
	log.Print(c.colorize("", t))
}
