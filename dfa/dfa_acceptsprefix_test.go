/*
DFA models used for testing taken from "Introduction to the Theory of
Computation", Third Edition, Michael Sipser. (C) 2013 Cengage Learning.
*/
package dfa_test

import (
	"github.com/paulgriffiths/automata/dfa"
	"github.com/paulgriffiths/gods/sets"
	"testing"
)

/*
DFA M2 accepts all strings that end with a 1.
*/
func TestAcceptsPrefixM2(t *testing.T) {
	d := dfa.Dfa{
		2,
		sets.NewSetRune('0', '1'),
		[]map[rune]int{
			{'0': 0, '1': 1},
			{'0': 0, '1': 1},
		},
		0,
		sets.NewSetInt(1),
	}

	testCases := []struct {
		input   string
		matches bool
		length  int
	}{
		{"", false, 0},
		{"0", false, 0},
		{"1", true, 1},
		{"00", false, 0},
		{"01", true, 2},
		{"10", true, 1},
		{"11", true, 2},
		{"000", false, 0},
		{"001", true, 3},
		{"010", true, 2},
		{"011", true, 3},
		{"100", true, 1},
		{"101", true, 3},
		{"110", true, 2},
		{"111", true, 3},
		{"0001001010100011111000000", true, 19},
		{"0001001010100011111000001", true, 25},
		{"x11", false, 0},
	}

	for _, c := range testCases {
		m, l := d.AcceptsPrefix(c.input)
		if m != c.matches || l != c.length {
			t.Errorf("input %q, got (%t, %d), want (%t, %d)",
				c.input, m, l, c.matches, c.length)
		}
	}
}

/*
DFA M4 accepts strings that start and end with the same letter.
*/
func TestAcceptsPrefixM4(t *testing.T) {
	d := dfa.Dfa{
		5,
		sets.NewSetRune('a', 'b'),
		[]map[rune]int{
			{'a': 1, 'b': 3},
			{'a': 1, 'b': 2},
			{'a': 1, 'b': 2},
			{'a': 4, 'b': 3},
			{'a': 4, 'b': 3},
		},
		0,
		sets.NewSetInt(1, 3),
	}

	testCases := []struct {
		input   string
		matches bool
		length  int
	}{
		{"", false, 0},
		{"a", true, 1},
		{"b", true, 1},
		{"aa", true, 2},
		{"bb", true, 2},
		{"ab", true, 1},
		{"ba", true, 1},
		{"aaa", true, 3},
		{"aab", true, 2},
		{"aba", true, 3},
		{"abb", true, 1},
		{"baa", true, 1},
		{"bab", true, 3},
		{"bba", true, 2},
		{"bbb", true, 3},
		{"tt", false, 0},
	}

	for _, c := range testCases {
		m, l := d.AcceptsPrefix(c.input)
		if m != c.matches || l != c.length {
			t.Errorf("input %q, got (%t, %d), want (%t, %d)",
				c.input, m, l, c.matches, c.length)
		}
	}
}
