package dfa_test

import (
	"fmt"
	"github.com/paulgriffiths/automata/dfa"
	"github.com/paulgriffiths/gods/sets"
)

func Example() {
	automaton := dfa.Dfa{
		Q: 5,
		S: sets.NewSetRune('a', 'b'),
		D: []map[rune]int{
			{'a': 1, 'b': 3},
			{'a': 1, 'b': 2},
			{'a': 1, 'b': 2},
			{'a': 4, 'b': 3},
			{'a': 4, 'b': 3},
		},
		Qs: 0,
		F:  sets.NewSetInt(1, 3),
	}

	for _, s := range []string{"abbba", "abbb"} {
		if automaton.Accepts(s) {
			fmt.Printf("DFA accepts string %q.\n", s)
		} else {
			fmt.Printf("DFA does not accept string %q.\n", s)
		}
	}

	for _, s := range []string{"ababb", "abbbb"} {
		accepted, n := automaton.AcceptsPrefix(s)
		if accepted {
			fmt.Printf("DFA accepts prefix %q of string %q.\n", s[:n], s)
		} else {
			fmt.Printf("DFA does not accept any prefix of string %q.\n", s)
		}
	}

	// Output:
	// DFA accepts string "abbba".
	// DFA does not accept string "abbb".
	// DFA accepts prefix "aba" of string "ababb".
	// DFA accepts prefix "a" of string "abbbb".
}
