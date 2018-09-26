package nfa_test

import (
	"fmt"
	"github.com/paulgriffiths/automata/nfa"
	"github.com/paulgriffiths/gods/sets"
)

func Example() {
	automaton := nfa.Nfa{
		Q: 5,
		S: sets.NewSetRune('a', 'b'),
		D: []map[rune]sets.SetInt{
			{0: sets.NewSetInt(1, 3)},
			{'a': sets.NewSetInt(2)},
			{'a': sets.NewSetInt(2)},
			{'b': sets.NewSetInt(4)},
			{'b': sets.NewSetInt(4)},
		},
		Qs: 0,
		F:  sets.NewSetInt(2, 4),
	}

	for _, s := range []string{"aaaa", "bbb", "aaabba"} {
		if automaton.Accepts(s) {
			fmt.Printf("NFA accepts string %q.\n", s)
		} else {
			fmt.Printf("NFA does not accept string %q.\n", s)
		}
	}

	dfa := automaton.ToDfa()

	for _, s := range []string{"aaabba", "baaab"} {
		if accepted, n := dfa.AcceptsPrefix(s); accepted {
			fmt.Printf("DFA accepts prefix %q of string %q.\n", s[:n], s)
		} else {
			fmt.Printf("DFA does not accept any prefix of string %q.\n", s)
		}
	}

	// Output:
	// NFA accepts string "aaaa".
	// NFA accepts string "bbb".
	// NFA does not accept string "aaabba".
	// DFA accepts prefix "aaa" of string "aaabba".
	// DFA accepts prefix "b" of string "baaab".
}
