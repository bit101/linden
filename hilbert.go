package linden

import "math"

// Hilbert is an LSystem that creates a Hilbert curve.
var Hilbert = LSystem{
	Seed: "A",
	Rules: RuleSet{
		"A": "-BF+AFA+FB-",
		"B": "+AF-BFB-FA+",
	},
	TurnSize:     math.Pi / 2,
	InitialAngle: 0.0,
	Instructions: InstructionSet{
		"F": "draw",
		"+": "turnRight",
		"-": "turnLeft",
	},
}
