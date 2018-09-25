package nfa_test

import (
	"github.com/paulgriffiths/automata/nfa"
	"testing"
)

// Matches ab
func TestNfaConcat1(t *testing.T) {
	a := nfa.NewRuneNfa('a')
	b := nfa.NewRuneNfa('b')
	c := nfa.NewConcatNfa(a, b)

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"a", false},
		{"b", false},
		{"aa", false},
		{"ab", true},
		{"ba", false},
		{"bb", false},
		{"aba", false},
		{"abb", false},
	}

	for _, tc := range testCases {
		if r := c.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Matches abc
func TestNfaConcat2(t *testing.T) {
	a := nfa.NewRuneNfa('a')
	b := nfa.NewRuneNfa('b')
	c := nfa.NewRuneNfa('c')
	n := nfa.NewConcatNfa(a, b)
	n = nfa.NewConcatNfa(n, c)

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"a", false},
		{"b", false},
		{"aa", false},
		{"ab", false},
		{"ba", false},
		{"bb", false},
		{"aba", false},
		{"abb", false},
		{"abc", true},
		{"acb", false},
		{"bac", false},
		{"bbc", false},
		{"abca", false},
		{"abcb", false},
		{"abcc", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Matches a*
func TestNfaClosure2(t *testing.T) {
	n := nfa.NewRuneNfa('a')
	n = nfa.NewClosureNfa(n)

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aaa", true},
		{"aaaa", true},
		{"aaaaa", true},
		{"aaaaaa", true},
		{"ab", false},
		{"aab", false},
		{"aaab", false},
		{"aaaab", false},
		{"aaaaab", false},
		{"aaaaaab", false},
		{"b", false},
		{"bb", false},
		{"bbb", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Matches a*b*c*
func TestNfaClosure3(t *testing.T) {
	a := nfa.NewRuneNfa('a')
	a = nfa.NewClosureNfa(a)
	b := nfa.NewRuneNfa('b')
	b = nfa.NewClosureNfa(b)
	c := nfa.NewRuneNfa('c')
	c = nfa.NewClosureNfa(c)
	n := nfa.NewConcatNfa(a, b)
	n = nfa.NewConcatNfa(n, c)

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aaa", true},
		{"aaaa", true},
		{"aaaaa", true},
		{"b", true},
		{"bb", true},
		{"bbb", true},
		{"bbbb", true},
		{"bbbbb", true},
		{"c", true},
		{"cc", true},
		{"ccc", true},
		{"cccc", true},
		{"ccccc", true},
		{"ac", true},
		{"aacc", true},
		{"aaaccc", true},
		{"aaaacccc", true},
		{"aaaaaccccc", true},
		{"bc", true},
		{"bbcc", true},
		{"bbbccc", true},
		{"bbbbcccc", true},
		{"bbbbbccccc", true},
		{"ab", true},
		{"aabb", true},
		{"aaabbb", true},
		{"aaaabbbb", true},
		{"aaaaabbbbb", true},
		{"ba", false},
		{"ca", false},
		{"cb", false},
		{"aba", false},
		{"aca", false},
		{"bcb", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Matches a|b
func TestNfaUnion1(t *testing.T) {
	a := nfa.NewRuneNfa('a')
	b := nfa.NewRuneNfa('b')
	n := nfa.NewUnionNfa(a, b)

	testCases := []struct {
		input  string
		result bool
	}{
		{"", false},
		{"a", true},
		{"b", true},
		{"ab", false},
		{"aa", false},
		{"ba", false},
		{"bb", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Matches a*|b*
func TestNfaUnion2(t *testing.T) {
	a := nfa.NewRuneNfa('a')
	a = nfa.NewClosureNfa(a)
	b := nfa.NewRuneNfa('b')
	b = nfa.NewClosureNfa(b)
	n := nfa.NewUnionNfa(a, b)

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aaa", true},
		{"b", true},
		{"bb", true},
		{"bbb", true},
		{"ab", false},
		{"ba", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}

// Matches (a*|b*)(c*|d*)
func TestNfaAll1(t *testing.T) {
	a := nfa.NewRuneNfa('a')
	a = nfa.NewClosureNfa(a)
	b := nfa.NewRuneNfa('b')
	b = nfa.NewClosureNfa(b)
	e := nfa.NewUnionNfa(a, b)
	c := nfa.NewRuneNfa('c')
	c = nfa.NewClosureNfa(c)
	d := nfa.NewRuneNfa('d')
	d = nfa.NewClosureNfa(d)
	f := nfa.NewUnionNfa(c, d)
	n := nfa.NewConcatNfa(e, f)

	testCases := []struct {
		input  string
		result bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aaa", true},
		{"b", true},
		{"bb", true},
		{"bbb", true},
		{"ab", false},
		{"ba", false},
		{"c", true},
		{"cc", true},
		{"ccc", true},
		{"d", true},
		{"dd", true},
		{"ddd", true},
		{"cd", false},
		{"dc", false},
		{"aaaaaccccc", true},
		{"aaaaaddddd", true},
		{"bbbbbccccc", true},
		{"bbbbbddddd", true},
		{"aaaaaccccca", false},
		{"aaaaaddddda", false},
		{"bbbbbcccccb", false},
		{"bbbbbdddddb", false},
	}

	for _, tc := range testCases {
		if r := n.Accepts(tc.input); r != tc.result {
			t.Errorf("input %q, got %v, want %v", tc.input, r, tc.result)
		}
	}
}
