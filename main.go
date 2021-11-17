package main

import (
	"fmt"

	"github.com/RyoJerryYu/gogoengine/engine"
)

func main() {
	var c engine.Engine = engine.NewEngine()
	fmt.Printf("c: %v\n", c)
}
