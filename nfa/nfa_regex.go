package nfa

import "github.com/paulgriffiths/gods/sets"

// NewRuneNfa creates a new NFA of two states connected by a
// single transition on the specified input character.
func NewRuneNfa(r rune) Nfa {
	return Nfa{
		Q:      2,
		S:      sets.NewSetRune(r),
		D:      []map[rune]sets.SetInt{{r: sets.NewSetInt(1)}, {}},
		Start:  0,
		Accept: sets.NewSetInt(1),
	}
}

// NewConcatNfa creates an Nfa representing the concatenation
// of the two provided Nfas.
// Note: this function assumes that the final state is the single
// accepting state for both provided Nfas. All the New...Nfa functions
// in this file create Nfas that satisfy that condition, so it's
// safe to rely on the assumption if these and only these functions
// are used together.
func NewConcatNfa(a, b Nfa) Nfa {
	return Nfa{
		Q:      a.Q + b.Q - 1,
		S:      a.S.Union(b.S),
		D:      append(a.D[:a.Q-1], advanceD(b.D, a.Q-1)...),
		Start:  0,
		Accept: advanceSet(b.Accept, a.Q-1),
	}
}

// NewUnionNfa creates an Nfa representing the union
// of the two provided Nfas.
// Note: this function assumes that the final state is the single
// accepting state for both provided Nfas. All the New...Nfa functions
// in this file create Nfas that satisfy that condition, so it's
// safe to rely on the assumption if these and only these functions
// are used together.
func NewUnionNfa(a, b Nfa) Nfa {
	dA := advanceD(a.D, 1)
	dB := advanceD(b.D, 1+a.Q)
	dA[a.Q-1][0] = sets.NewSetInt(a.Q + b.Q + 1)
	dB[b.Q-1][0] = sets.NewSetInt(a.Q + b.Q + 1)
	dJ := []map[rune]sets.SetInt{{0: sets.NewSetInt(1, a.Q+1)}}
	dJ = append(dJ, dA...)
	dJ = append(dJ, dB...)
	dJ = append(dJ, map[rune]sets.SetInt{})

	return Nfa{
		Q:      a.Q + b.Q + 2,
		S:      a.S.Union(b.S),
		D:      dJ,
		Start:  0,
		Accept: sets.NewSetInt(a.Q + b.Q + 1),
	}
}

// NewClosureNfa creates an Nfa representing the closure or Kleene
// star operation of the provided Nfa.
// Note: this function assumes that the final state is the single
// accepting state for the provided Nfa. All the New...Nfa functions
// in this file create Nfas that satisfy that condition, so it's
// safe to rely on the assumption if these and only these functions
// are used together.
func NewClosureNfa(n Nfa) Nfa {
	d := []map[rune]sets.SetInt{{0: sets.NewSetInt(1, n.Q+1)}}
	d = append(d, advanceD(n.D, 1)...)
	d = append(d, map[rune]sets.SetInt{})
	d[n.Q][0] = sets.NewSetInt(1, n.Q+1)
	return Nfa{
		Q:      n.Q + 2,
		S:      n.S,
		D:      d,
		Start:  0,
		Accept: sets.NewSetInt(n.Q + 1),
	}
}

// advanceSet returns a new set of integers representing the
// provided set of integers where all the elements have been
// increased in value by n. This is necessary for joining two
// Nfas together whose states are both initially assumed to be
// 0...Q-1.
func advanceSet(s sets.SetInt, n int) sets.SetInt {
	newSet := sets.NewSetInt()
	for _, elem := range s.Elements() {
		newSet.Insert(elem + n)
	}
	return newSet
}

// advanceD modifies in place a transition function to increase
// in value by n all the advanced-to states, and returns the
// modified transition function as a convenience. This is necessary
// for joining two Nfas together whose states are both initially
// assumed to be 0...Q-1.
func advanceD(d []map[rune]sets.SetInt, n int) []map[rune]sets.SetInt {
	for i := range d {
		for key, value := range d[i] {
			if !value.IsEmpty() {
				d[i][key] = advanceSet(value, n)
			}
		}
	}
	return d
}
