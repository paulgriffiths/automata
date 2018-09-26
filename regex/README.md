# regex 

This package implements a regular expression engine by using the
McNaughton-Yamada-Thompson algorithm to transform a regular expression
into an equivalent nondeterministic finite automaton, and then transforming
that into an equivalent deterministic finite automaton for matching.

*Note*: this package implements the regular expressions of formal language
theory (except that the empty string is not implemented), which may appear
significantly more limited than the more familiar Unix regular expressions,
as they lack, for example, wildcards, character classes, metacharacters,
and capturing groups, among other things.

Regular expressions in formal language theory are equivalent to regular
grammars, and they provide three operations over a language's alphabet:

* concatenation;
* union, or alternation; and
* Kleene star, or closure.

The Kleene star has the highest priority, followed by concatention, then
union. Parentheses may be used to alter the priority. Examples of supported
regular expressions over the alphabet {a, b} include:

* aba
* aa\*bb\*
* (a|b)\*
* aa(a|b)\*bb\*
* ((aa|bb)(aa|bb))\*

The `Compile` method converts a regular expression in string form to an
equivalent deterministic finite automata. The `Match` and `MatchPrefix`
methods of the compiled regular expression may then be used to test whether
an entire string or any prefix of a string can be matched by the regular
expression.

### Example

```go
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
```
