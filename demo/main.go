package main

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
	"github.com/bit101/linden"
)

const (
	filename = "out.png"
	width    = 800.0
	height   = 800.0
)

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.85)

	// linden.DrawSystem(surface, linden.BinaryTree, 9)
	// linden.DrawSystem(surface, linden.Hilbert, 5)
	// linden.DrawSystem(surface, linden.Dragon, 14)
	// linden.DrawSystem(surface, linden.Plant, 3)
	linden.DrawSystem(surface, linden.Triangle, 9)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}
