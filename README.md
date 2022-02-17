# Colorog

[![Test](https://github.com/wander4747/colorog/actions/workflows/test.yml/badge.svg)](https://github.com/wander4747/colorog/actions/workflows/test.yml)
[![Lint](https://github.com/wander4747/colorog/actions/workflows/lint.yml/badge.svg)](https://github.com/wander4747/colorog/actions/workflows/lint.yml)

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
	l := colorog.New()
	l.Success("success")
	l.Info("info")
	l.Warning("warning")
	l.Light("light")
	l.WithColor(color.Green, "with color")
	l.Fatal("fatal")
}

```
