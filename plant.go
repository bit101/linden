package linden

import "math"

// Plant is an LSystem that creates a plant
var Plant = LSystem{
	Seed: "X",
	Rules: RuleSet{
		"X": "F+[[X]-X]-F[-FX]+X",
		"F": "FF",
	},
	TurnSize:     25.0 * math.Pi / 180.0,
	InitialAngle: -math.Pi / 2,
	Instructions: InstructionSet{
		"F": "draw",
		"+": "turnRight",
		"-": "turnLeft",
		"[": "push",
		"]": "pop",
	},
}
