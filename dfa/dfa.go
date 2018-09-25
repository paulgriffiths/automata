package dfa

import "github.com/paulgriffiths/gods/sets"

// Dfa implements a deterministic finite automaton.
type Dfa struct {
	Q  int            // Number of states
	S  sets.SetRune   // Alphabet
	D  []map[rune]int // Transition function
	Qs int            // Start state
	F  sets.SetInt    // Set of accepting states
}

// Accepts returns true if the DFA accepts the provided string.
func (d Dfa) Accepts(input string) bool {
	currentState := d.Qs
	ok := false

	for _, letter := range input {
		currentState, ok = d.D[currentState][letter]
		if !ok {
			return false
		}
	}

	return d.F.Contains(currentState)
}

// AcceptsPrefix checks if there is a prefix of the provided string
// which is accepted by the DFA. If it is, the function returns true
// and the length of the prefix. Otherwise, it returns false and zero.
func (d Dfa) AcceptsPrefix(input string) (bool, int) {
	currentState := d.Qs
	ok := false
	matches := false
	longest := 0

	if len(input) == 0 && d.F.Contains(currentState) {
		return true, 0
	}

	for n, letter := range input {
		currentState, ok = d.D[currentState][letter]
		if !ok {
			break
		}
		if d.F.Contains(currentState) {
			matches = true
			longest = n + 1
		}
	}

	return matches, longest
}
