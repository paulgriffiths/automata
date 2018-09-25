package nfa

import "github.com/paulgriffiths/gods/sets"

// Nfa implements a nondeterministic finite automaton.
type Nfa struct {
	Q      int                    // Number of states
	S      sets.SetRune           // Alphabet
	D      []map[rune]sets.SetInt // Transition function
	Start  int                    // Start state
	Accept sets.SetInt            // Set of accepting states
}

// Accepts returns true if the NFA accepts the provided string.
func (n Nfa) Accepts(input string) bool {
	current := n.EclosureS(n.Start)
	for _, letter := range input {
		current = n.EclosureT(n.Move(current, letter))
	}
	return !n.Accept.Intersection(current).IsEmpty()
}

// EclosureS returns the set of states reachable from the specified
// state on e-transitions alone. Note that a path can have zero edges,
// so any state is reachable from itself by an e-labeled path.
func (n Nfa) EclosureS(s int) sets.SetInt {
	current := sets.NewSetInt(s)
	ecl := current
	prevLength := -1

	for ecl.Length() != prevLength {
		prevLength = ecl.Length()
		next := sets.NewSetInt()
		for _, state := range current.Elements() {
			if eStates, ok := n.D[state][0]; ok {
				ecl = ecl.Union(eStates)
				next = next.Union(eStates)
			}
		}
		current = next
	}

	return ecl
}

// EclosureT returns the set of states reachable from the provided
// set of states on e-transitions alone.
func (n Nfa) EclosureT(t sets.SetInt) sets.SetInt {
	ecl := sets.NewSetInt()
	for _, state := range t.Elements() {
		ecl = ecl.Union(n.EclosureS(state))
	}
	return ecl
}

// Move returns the set of states reachable from set t on
// input symbol a.
func (n Nfa) Move(t sets.SetInt, a rune) sets.SetInt {
	trans := sets.NewSetInt()
	for _, state := range t.Elements() {
		if p, ok := n.D[state][a]; ok {
			trans = trans.Union(p)
		}
	}
	return trans
}
