# Colorog

[![Test](https://github.com/wander4747/colorog/actions/workflows/test.yaml/badge.svg)](https://github.com/wander4747/colorog/actions/workflows/test.yaml)
[![Lint](https://github.com/wander4747/colorog/actions/workflows/lint.yaml/badge.svg)](https://github.com/wander4747/colorog/actions/workflows/lint.yaml)

Simple package for coloring text in terminals

## Installation
```sh
go get github.com/wander4747/colorog
```

## How to use

```go
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
```
