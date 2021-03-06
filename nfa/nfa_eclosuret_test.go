/*
NFA models used for testing taken from "Introduction to the Theory of
Computation", Third Edition, Michael Sipser. (C) 2013 Cengage Learning,
and from "Compilers", Second Edition, Aho, Lam, Sethi & Ullman, (C) 2007
Pearson Education, Inc.
*/
package nfa_test

import (
	"github.com/paulgriffiths/automata/nfa"
	"github.com/paulgriffiths/gods/sets"
	"testing"
)

// NFA from Compilers, figure 3.34.
func TestNfaEclosureT1(t *testing.T) {
	n := nfa.Nfa{
		11,
		sets.NewSetRune('a', 'b'),
		[]map[rune]sets.SetInt{
			{0: sets.NewSetInt(1, 7)},
			{0: sets.NewSetInt(2, 4)},
			{'a': sets.NewSetInt(3)},
			{0: sets.NewSetInt(6)},
			{'b': sets.NewSetInt(5)},
			{0: sets.NewSetInt(6)},
			{0: sets.NewSetInt(1, 7)},
			{'a': sets.NewSetInt(8)},
			{'b': sets.NewSetInt(9)},
			{'b': sets.NewSetInt(10)},
			{},
		},
		0,
		sets.NewSetInt(10),
	}

	testCases := []struct {
		states, result sets.SetInt
	}{
		{sets.NewSetInt(0), sets.NewSetInt(0, 1, 2, 4, 7)},
		{sets.NewSetInt(2), sets.NewSetInt(2)},
		{sets.NewSetInt(5), sets.NewSetInt(1, 2, 4, 5, 6, 7)},
		{sets.NewSetInt(10), sets.NewSetInt(10)},
		{sets.NewSetInt(2, 4), sets.NewSetInt(2, 4)},
		{sets.NewSetInt(0, 1), sets.NewSetInt(0, 1, 2, 4, 7)},
		{sets.NewSetInt(3, 5), sets.NewSetInt(1, 2, 3, 4, 5, 6, 7)},
		{sets.NewSetInt(4, 6, 9), sets.NewSetInt(1, 2, 4, 6, 7, 9)},
	}

	for i, tc := range testCases {
		s := n.EclosureT(tc.states)
		if !s.Equals(tc.result) {
			t.Errorf("case %d, got %v, want %v", i+1, s, tc.result)
		}
	}
}

// NFA from Compilers, figure 3.30.
// Every state has an e-transition to every other state in this
// model, and we have to be careful we don't cause an infinite
// loop, so it's a good candidate for testing.
func TestNfaEclosureT2(t *testing.T) {
	n := nfa.Nfa{
		4,
		sets.NewSetRune('a', 'b'),
		[]map[rune]sets.SetInt{
			{0: sets.NewSetInt(3), 'a': sets.NewSetInt(1)},
			{0: sets.NewSetInt(0), 'b': sets.NewSetInt(2)},
			{0: sets.NewSetInt(1), 'b': sets.NewSetInt(3)},
			{0: sets.NewSetInt(2), 'a': sets.NewSetInt(0)},
		},
		0,
		sets.NewSetInt(3),
	}

	testCases := []struct {
		states, result sets.SetInt
	}{
		{sets.NewSetInt(0), sets.NewSetInt(0, 1, 2, 3)},
		{sets.NewSetInt(1), sets.NewSetInt(0, 1, 2, 3)},
		{sets.NewSetInt(0, 1), sets.NewSetInt(0, 1, 2, 3)},
		{sets.NewSetInt(1, 2), sets.NewSetInt(0, 1, 2, 3)},
		{sets.NewSetInt(2, 3), sets.NewSetInt(0, 1, 2, 3)},
		{sets.NewSetInt(1, 2, 3), sets.NewSetInt(0, 1, 2, 3)},
		{sets.NewSetInt(0, 1, 2, 3), sets.NewSetInt(0, 1, 2, 3)},
	}

	for i, tc := range testCases {
		s := n.EclosureT(tc.states)
		if !s.Equals(tc.result) {
			t.Errorf("case %d, got %v, want %v", i+1, s, tc.result)
		}
	}
}
