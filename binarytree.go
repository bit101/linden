package linden

import "math"

// BinaryTree is an LSystem that creates a binary tree.
var BinaryTree = LSystem{
	Seed: "A",
	Rules: RuleSet{
		"A": "B[-A]+A",
		"B": "BB",
	},
	TurnSize:     math.Pi / 4,
	InitialAngle: -math.Pi / 2,
	Instructions: InstructionSet{
		"A": "draw",
		"B": "draw",
		"+": "turnRight",
		"-": "turnLeft",
		"[": "push",
		"]": "pop",
	},
}
