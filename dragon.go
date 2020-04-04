package linden

import "math"

// Dragon is an LSystem that creates a Dragon curve
var Dragon = LSystem{
	Seed: "FX",
	Rules: RuleSet{
		"X": "X+YF+",
		"Y": "-FX-Y",
	},
	TurnSize:     math.Pi / 2,
	InitialAngle: 0.0,
	Instructions: InstructionSet{
		"F": "draw",
		"+": "turnRight",
		"-": "turnLeft",
	},
}
