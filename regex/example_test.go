package regex_test

import (
	"fmt"
	"github.com/paulgriffiths/automata/regex"
)

func Example() {
	s := "aa(a|b)*bb"
	r := regex.Compile(s)

	for _, t := range []string{"aabb", "aaabb", "aba", "aabbbba"} {
		if r.Match(t) {
			fmt.Printf("String %q matches %s.\n", t, s)
		} else {
			fmt.Printf("String %q doesn't match %s.\n", t, s)
		}

		if matches, n := r.MatchPrefix(t); matches {
			fmt.Printf("Prefix %q of string %q matches %s.\n", t[:n], t, s)
		} else {
			fmt.Printf("No prefix of string %q matches %s.\n", t, s)
		}
	}

	// Output:
	// String "aabb" matches aa(a|b)*bb.
	// Prefix "aabb" of string "aabb" matches aa(a|b)*bb.
	// String "aaabb" matches aa(a|b)*bb.
	// Prefix "aaabb" of string "aaabb" matches aa(a|b)*bb.
	// String "aba" doesn't match aa(a|b)*bb.
	// No prefix of string "aba" matches aa(a|b)*bb.
	// String "aabbbba" doesn't match aa(a|b)*bb.
	// Prefix "aabbbb" of string "aabbbba" matches aa(a|b)*bb.
}
