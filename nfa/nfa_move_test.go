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
func TestNfaMove1(t *testing.T) {
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
		a              rune
		states, result sets.SetInt
	}{
		{'a', sets.NewSetInt(0), sets.NewSetInt()},
		{'b', sets.NewSetInt(0), sets.NewSetInt()},
		{'a', sets.NewSetInt(1), sets.NewSetInt()},
		{'b', sets.NewSetInt(1), sets.NewSetInt()},
		{'a', sets.NewSetInt(2), sets.NewSetInt(3)},
		{'b', sets.NewSetInt(2), sets.NewSetInt()},
		{'a', sets.NewSetInt(3), sets.NewSetInt()},
		{'b', sets.NewSetInt(3), sets.NewSetInt()},
		{'a', sets.NewSetInt(4), sets.NewSetInt()},
		{'b', sets.NewSetInt(4), sets.NewSetInt(5)},
		{'a', sets.NewSetInt(5), sets.NewSetInt()},
		{'b', sets.NewSetInt(5), sets.NewSetInt()},
		{'a', sets.NewSetInt(6), sets.NewSetInt()},
		{'b', sets.NewSetInt(6), sets.NewSetInt()},
		{'a', sets.NewSetInt(7), sets.NewSetInt(8)},
		{'b', sets.NewSetInt(7), sets.NewSetInt()},
		{'a', sets.NewSetInt(8), sets.NewSetInt()},
		{'b', sets.NewSetInt(8), sets.NewSetInt(9)},
		{'a', sets.NewSetInt(9), sets.NewSetInt()},
		{'b', sets.NewSetInt(9), sets.NewSetInt(10)},
		{'a', sets.NewSetInt(10), sets.NewSetInt()},
		{'b', sets.NewSetInt(10), sets.NewSetInt()},
		{'a', sets.NewSetInt(1, 2, 4), sets.NewSetInt(3)},
		{'b', sets.NewSetInt(1, 2, 4), sets.NewSetInt(5)},
		{'a', sets.NewSetInt(2, 7, 8), sets.NewSetInt(3, 8)},
		{'b', sets.NewSetInt(2, 7, 8), sets.NewSetInt(9)},
		{'a', sets.NewSetInt(3, 5, 9, 10), sets.NewSetInt()},
		{'b', sets.NewSetInt(3, 5, 9, 10), sets.NewSetInt(10)},
		{'a', sets.NewSetInt(0, 1, 2, 3, 4, 5), sets.NewSetInt(3)},
		{'b', sets.NewSetInt(0, 1, 2, 3, 4, 5), sets.NewSetInt(5)},
	}

	for i, tc := range testCases {
		s := n.Move(tc.states, tc.a)
		if !s.Equals(tc.result) {
			t.Errorf("case %d, got %v, want %v", i+1, s, tc.result)
		}
	}
}

// NFA from Compilers, figure 3.30.
// Every state has an e-transition to every other state in this
// model, and we have to be careful we don't cause an infinite
// loop, so it's a good candidate for testing.
func TestNfaMove2(t *testing.T) {
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
		a              rune
		states, result sets.SetInt
	}{
		{'a', sets.NewSetInt(0), sets.NewSetInt(1)},
		{'b', sets.NewSetInt(0), sets.NewSetInt()},
		{'a', sets.NewSetInt(1), sets.NewSetInt()},
		{'b', sets.NewSetInt(1), sets.NewSetInt(2)},
		{'a', sets.NewSetInt(2), sets.NewSetInt()},
		{'b', sets.NewSetInt(2), sets.NewSetInt(3)},
		{'a', sets.NewSetInt(3), sets.NewSetInt(0)},
		{'b', sets.NewSetInt(3), sets.NewSetInt()},
		{'a', sets.NewSetInt(0, 1, 2), sets.NewSetInt(1)},
		{'b', sets.NewSetInt(0, 1, 2), sets.NewSetInt(2, 3)},
		{'a', sets.NewSetInt(0, 1, 2, 3), sets.NewSetInt(0, 1)},
		{'b', sets.NewSetInt(0, 1, 2, 3), sets.NewSetInt(2, 3)},
	}

	for i, tc := range testCases {
		s := n.Move(tc.states, tc.a)
		if !s.Equals(tc.result) {
			t.Errorf("case %d, got %v, want %v", i+1, s, tc.result)
		}
	}
}
