package main

import (
	"github.com/wander4747/colorog"
	"github.com/wander4747/colorog/color"
)

func main() {
	l := colorog.New()
	l.Success("success")
	l.Info("info")
	l.Warning("warning")
	l.Light("light")
	l.WithColor(color.Green, "with color")
	l.Unicorn("Mensagem do unicórnio: bom final de semana e até amanhã!")
	l.Fatal("fatal")
}
