package linden

import "math"

type operation int

const (
	move operation = iota
	draw
)

type point struct {
	x, y float64
	op   operation
}

func newPoint(x, y float64, op operation) point {
	return point{x, y, op}
}

type path struct {
	points []point
}

func newPath() *path {
	return &path{}
}

func (p *path) push(pos point) {
	p.points = append(p.points, pos)
}

// Stack is the stack of angles and positions
type stack struct {
	angleStack []float64
	pointStack []point
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(angle float64, p point) {
	s.angleStack = append(s.angleStack, angle)
	s.pointStack = append(s.pointStack, p)
}

func (s *stack) pop() (float64, point) {
	angle := s.angleStack[len(s.angleStack)-1]
	s.angleStack = s.angleStack[:len(s.angleStack)-1]

	p := s.pointStack[len(s.pointStack)-1]
	s.pointStack = s.pointStack[:len(s.pointStack)-1]

	return angle, p
}

type sizer struct {
	minX, maxX, minY, maxY float64
}

func newSizer() *sizer {
	return &sizer{0.0, 0.0, 0.0, 0.0}
}

func (s *sizer) adjust(p point) {
	s.minX = math.Min(p.x, s.minX)
	s.maxX = math.Max(p.x, s.maxX)
	s.minY = math.Min(p.y, s.minY)
	s.maxY = math.Max(p.y, s.maxY)
}

func (s *sizer) xrange() float64 {
	return s.maxX - s.minX
}

func (s *sizer) yrange() float64 {
	return s.maxY - s.minY
}
