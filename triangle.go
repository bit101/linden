package linden

import "math"

// Triangle is an LSystem that creates a Sierpinski triangle
var Triangle = LSystem{
	Seed: "A",
	Rules: RuleSet{
		"A": "B-A-B",
		"B": "A+B+A",
	},
	TurnSize:     60.0 * math.Pi / 180.0,
	InitialAngle: 0.0,
	Instructions: InstructionSet{
		"A": "draw",
		"B": "draw",
		"+": "turnRight",
		"-": "turnLeft",
	},
}
