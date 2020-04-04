package linden

// RuleSet defines a set of Lindenmeyer string transform rules.
type RuleSet map[string]string

// InstructionSet defines a set of Lindenmeyer string rendering rules.
type InstructionSet map[string]string

// LSystem defines the elements needed for a full Lindenmeyer system
type LSystem struct {
	Seed         string
	Rules        RuleSet
	TurnSize     float64
	InitialAngle float64
	Instructions InstructionSet
}

// ExpandSystem expands a specific LSystem by a specified number of iterations.
func ExpandSystem(system LSystem, count int) string {
	return Expand(system.Seed, system.Rules, count)
}

// Expand expands a Lindenmeyer seed string using a RuleSet a number of times
func Expand(seed string, rules RuleSet, count int) string {
	result := ""
	for _, c := range seed {
		char := string(c)
		rule, ok := rules[char]
		if ok {
			result += rule
		} else {
			result += char
		}
	}
	if count > 0 {
		return Expand(result, rules, count-1)
	}
	return result
}
