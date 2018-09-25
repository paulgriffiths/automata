/*
Package regex implements a simple regular expression compiler.

Regular expressions of the form ab*(c|d|e)*f are accepted. Letters
and digits only may be used as the language alphabet (no wildcards
are accepted).

The Kleene star or closure operator has the highest precedence, and
is right-associative. Concatenation has the next highest precedence,
and the union operator has the lowest precedence. Arbitary parentheses
may be used to group terms or alter the standard operator precedence.

The matching function attempts to match the entire string to the
regular expression, e.g. the string "ha" will match the regular
expresion "ha", but the string "that" will not.
*/
package regex
