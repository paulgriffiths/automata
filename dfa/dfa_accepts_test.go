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
DFA M1 accepts strings that contain at least one 1
with an even number of 0s following the last 1.
*/
func TestM1(t *testing.T) {
	d := dfa.Dfa{
		3,
		sets.NewSetRune('0', '1'),
		[]map[rune]int{
			{'0': 0, '1': 1},
			{'0': 2, '1': 1},
			{'0': 1, '1': 1},
		},
		0,
		sets.NewSetInt(1),
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"0", false},
		{"1", true},
		{"00", false},
		{"01", true},
		{"10", false},
		{"11", true},
		{"000", false},
		{"001", true},
		{"010", false},
		{"011", true},
		{"100", true},
		{"101", true},
		{"110", false},
		{"111", true},
		{"0001001010100011111000000", true},
		{"1y1", false},
	}

	for _, c := range testCases {
		a := d.Accepts(c.input)
		if a != c.result {
			t.Errorf("input %q, got %v, want %v", c.input, a, c.result)
		}
	}
}

/*
DFA M2 accepts all strings that end with a 1.
*/
func TestM2(t *testing.T) {
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
		input  string
		result bool
	}{
		{"", false},
		{"0", false},
		{"1", true},
		{"00", false},
		{"01", true},
		{"10", false},
		{"11", true},
		{"000", false},
		{"001", true},
		{"010", false},
		{"011", true},
		{"100", false},
		{"101", true},
		{"110", false},
		{"111", true},
		{"0001001010100011111000000", false},
		{"0001001010100011111000001", true},
		{"x11", false},
	}

	for _, c := range testCases {
		a := d.Accepts(c.input)
		if a != c.result {
			t.Errorf("input %q, got %v, want %v", c.input, a, c.result)
		}
	}
}

/*
DFA M3 accepts the empty string and all strings that end in 0.
*/
func TestM3(t *testing.T) {
	d := dfa.Dfa{
		2,
		sets.NewSetRune('0', '1'),
		[]map[rune]int{
			{'0': 0, '1': 1},
			{'0': 0, '1': 1},
		},
		0,
		sets.NewSetInt(0),
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{"0", true},
		{"1", false},
		{"00", true},
		{"01", false},
		{"10", true},
		{"11", false},
		{"000", true},
		{"001", false},
		{"010", true},
		{"011", false},
		{"100", true},
		{"101", false},
		{"110", true},
		{"111", false},
		{"0001001010100011111000000", true},
		{"0001001010100011111000001", false},
		{"z10", false},
	}

	for _, c := range testCases {
		a := d.Accepts(c.input)
		if a != c.result {
			t.Errorf("input %q, got %v, want %v", c.input, a, c.result)
		}
	}
}

/*
DFA M4 accepts strings that start and end with the same letter.
*/
func TestM4(t *testing.T) {
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
		input  string
		result bool
	}{
		{"", false},
		{"a", true},
		{"b", true},
		{"aa", true},
		{"bb", true},
		{"ab", false},
		{"ba", false},
		{"aaa", true},
		{"aab", false},
		{"aba", true},
		{"abb", false},
		{"baa", false},
		{"bab", true},
		{"bba", false},
		{"bbb", true},
		{"tt", false},
	}

	for _, c := range testCases {
		a := d.Accepts(c.input)
		if a != c.result {
			t.Errorf("input %q, got %v, want %v", c.input, a, c.result)
		}
	}
}

/*
DFA M5 accepts strings where the sum of the numeric input symbols
is divisible by 3. The reset symbol '.' resets the count to zero.
*/
func TestM5(t *testing.T) {
	d := dfa.Dfa{
		5,
		sets.NewSetRune('.', '0', '1', '2'),
		[]map[rune]int{
			{'.': 0, '0': 0, '1': 1, '2': 2},
			{'.': 0, '0': 1, '1': 2, '2': 0},
			{'.': 0, '0': 2, '1': 0, '2': 1},
		},
		0,
		sets.NewSetInt(0),
	}

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{".", true},
		{"0", true},
		{"1", false},
		{"2", false},
		{"01", false},
		{"02", false},
		{"10", false},
		{"12", true},
		{"21", true},
		{"20", false},
		{"2.2", false},
		{"2.201", true},
		{"2.2.", true},
		{"2210121012021.012010.22.012", true},
		{"1212120021210120212101110210222012", true},
		{"1212120021210120212101110210222011", false},
		{"12+12", false},
	}

	for _, c := range testCases {
		a := d.Accepts(c.input)
		if a != c.result {
			t.Errorf("input %q, got %v, want %v", c.input, a, c.result)
		}
	}
}
