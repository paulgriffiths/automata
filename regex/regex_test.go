package regex_test

import (
	"github.com/paulgriffiths/automata/regex"
	"testing"
)

func TestRegexAccepts(t *testing.T) {
	testCases := []struct {
		rx, s  string
		result bool
	}{
		{"a", "a", true},
		{"a", "a", true},
		{"(a)", "a", true},
		{"(a)", "b", false},
		{"((((a))))", "a", true},
		{"a*", "", true},
		{"a*", "a", true},
		{"a*", "aa", true},
		{"a*", "aaa", true},
		{"a*", "b", false},
		{"ab", "ab", true},
		{"ab", "abc", false},
		{"abc", "abc", true},
		{"a*b*", "", true},
		{"a*b*", "aaa", true},
		{"a*b*", "bbbb", true},
		{"a*b*", "aaaabbb", true},
		{"a*(b*)", "aaaabbb", true},
		{"(a*)b*", "aaaabbb", true},
		{"(a*)(b*)", "aaaabbb", true},
		{"(a*b*)", "aaaabbb", true},
		{"a*b*", "bbbaaaa", false},
		{"a|b", "a", true},
		{"(a|b)", "a", true},
		{"(a)|(b)", "a", true},
		{"a|b", "b", true},
		{"a|b", "ab", false},
		{"a|b", "ba", false},
		{"a|b", "aa", false},
		{"a|b", "bb", false},
		{"a*|b*", "a", true},
		{"a*|b*", "aa", true},
		{"a*|b*", "aaa", true},
		{"a*|b*", "b", true},
		{"a*|b*", "bb", true},
		{"a*|b*", "bbb", true},
		{"a*|b*", "ab", false},
		{"th(i|e|o)se*", "this", true},
		{"th(i|e|o)se*", "those", true},
		{"th(i|e|o)se*", "these", true},
		{"th(i|e|o)se*", "thesa", false},
		{"f(a|e|u)*h", "fh", true},
		{"f(a|e|u)*h", "faah", true},
		{"f(a|e|u)*h", "feeeh", true},
		{"f(a|e|u)*h", "fuuuuh", true},
		{"f(a|e|u)*h", "foh", false},
		{"(a|b)(c|d)", "ab", false},
		{"(a|b)(c|d)", "ac", true},
		{"(a|b)(c|d)", "ad", true},
		{"(a|b)(c|d)", "ba", false},
		{"(a|b)(c|d)", "bc", true},
		{"(a|b)(c|d)", "bd", true},
		{"(a|b)(c|d)", "cd", false},
		{"((a|b)(c|d))*", "", true},
		{"((a|b)(c|d))*", "ac", true},
		{"((a|b)(c|d))*", "acbd", true},
		{"((a|b)(c|d))*", "acbdadbc", true},
		{"((a|b)(1|2|3)*(c|d))*", "a3cb21da131dbc", true},
		{"((a|b)(1|2|3)*(c|d))*", "a3ab21da131dbc", false},
	}

	for n, tc := range testCases {
		r := regex.Compile(tc.rx)
		if r == nil {
			t.Errorf("case %d, couldn't compile regex", n+1)
			continue
		}
		if result := r.Match(tc.s); result != tc.result {
			t.Errorf("case %d, got %t, want %t", n+1, result, tc.result)
		}
	}
}

func TestRegexAcceptsPrefix(t *testing.T) {
	testCases := []struct {
		rx, s   string
		matches bool
		length  int
	}{
		{"a*", "", true, 0},
		{"a*", "a", true, 1},
		{"a*", "ab", true, 1},
		{"a*", "aa", true, 2},
		{"a*", "aab", true, 2},
		{"(a|b)*", "aababababa", true, 10},
		{"(a|b)*", "aababababacdefg", true, 10},
		{"(a|b)*", "cdefgaababababa", false, 0},
		{"a*b*", "aaabbb", true, 6},
		{"a*b*", "aaabbbaaa", true, 6},
		{"a*b*", "bbbaaa", true, 3},
		{"a*b*", "bbbaaabbb", true, 3},
		{"a*b*", "cccbbbaaabbb", false, 0},
	}

	for n, tc := range testCases {
		r := regex.Compile(tc.rx)
		if r == nil {
			t.Errorf("case %d, couldn't compile regex", n+1)
			continue
		}
		if m, l := r.MatchPrefix(tc.s); m != tc.matches || l != tc.length {
			t.Errorf("case %d, got (%t, %d), want (%t, %d)",
				n+1, m, l, tc.matches, tc.length)
		}
	}
}

func TestBadRegex(t *testing.T) {
	testCases := []string{
		"",
		"^",
		"()",
		"(())",
		"(",
		")",
		"(a",
		"a)",
		"*",
		"*b",
		"|",
		"|*",
		"*|",
		"a|",
		"|a",
	}

	for n, tc := range testCases {
		r := regex.Compile(tc)
		if r != nil {
			t.Errorf("case %d, unexpectedly compiled regex %q", n+1, tc)
		}
	}
}
