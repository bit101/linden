package linden

import (
	"math"

	"github.com/bit101/blgo"
)

// DrawSystem renders a specified LSystem
func DrawSystem(surface *blgo.Surface, system LSystem, count int) {
	stack := newStack()
	path := newPath()

	a := system.InitialAngle
	p := newPoint(0.0, 0.0, move)

	sizer := newSizer()

	for _, c := range ExpandSystem(system, count) {
		switch system.Instructions[string(c)] {
		case "draw":
			p = newPoint(p.x+math.Cos(a), p.y+math.Sin(a), draw)
			sizer.adjust(p)
			path.push(p)
		case "turnLeft":
			a += system.TurnSize
		case "turnRight":
			a -= system.TurnSize
		case "push":
			stack.push(a, p)
		case "pop":
			a, p = stack.pop()
			path.push(newPoint(p.x, p.y, move))
		}
	}

	render(surface, path, sizer)
}

func render(surface *blgo.Surface, path *path, sizer *sizer) {
	w := float64(surface.GetWidth())
	h := float64(surface.GetHeight())
	xs := math.Abs((w - 20) / sizer.xrange())
	ys := math.Abs((h - 20) / sizer.yrange())
	scale := math.Min(xs, ys)

	surface.Save()
	surface.Translate(
		(w/2 - scale*(sizer.minX+sizer.xrange()/2)),
		(h/2 - scale*(sizer.minY+sizer.yrange()/2)),
	)
	surface.Scale(scale, scale)
	for _, p := range path.points {
		if p.op == draw {
			surface.LineTo(p.x, p.y)
		} else {
			surface.MoveTo(p.x, p.y)
		}
	}
	surface.Restore()
	surface.Stroke()
}
