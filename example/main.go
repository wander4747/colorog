package main

import (
	"github.com/wander4747/colorog"
	"github.com/wander4747/colorog/color"
)

func main() {
	colorog.Success("success")
	colorog.Info("info")
	colorog.Warning("warning")
	colorog.Light("light")
	colorog.WithColor(color.Green, "with color")
	colorog.Unicorn("Bom final de semana e até amanhã!")
	colorog.Fatal("fatal")
}
