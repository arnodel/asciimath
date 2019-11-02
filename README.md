## ASCIIMath in Go

the `scanner/symbols.go` is lifts directly the array of symbol definitions
https://raw.githubusercontent.com/asciimath/asciimathml/master/ASCIIMathML.js
with edits to make it compile in Go.

The scanner is defined in `scanner/scanner.go`.  It is extremely simple as it
builds a regexp out of the symbols array and uses is `Regexp.FindAllStringIndex`
to extract the tokens.

It seems to work as the tests in `scanner/scanner_test.go` show.  Would be
better to return token data rather than just token strings though!