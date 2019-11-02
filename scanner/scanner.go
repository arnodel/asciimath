package scanner

import (
	"errors"
	"regexp"
	"strings"
)

type Scanner struct {
	ptn *regexp.Regexp
}

func (s *Scanner) Tokenise(input string) ([]string, error) {
	var tokens []string
	start := 0
	end := 0
	for _, match := range s.ptn.FindAllStringIndex(input, len(input)) {
		start = match[0]
		if end != start {
			return nil, errors.New("cannot tokenise")
		}
		end = match[1]
		firstByte := input[start]
		if firstByte == ' ' || firstByte == '\t' || firstByte == '\n' || firstByte == '\r' {
			continue
		}
		tokens = append(tokens, input[start:end])
	}
	if end != len(input) {
		return nil, errors.New("cannot tokenise")
	}
	return tokens, nil
}

func newScanner() *Scanner {
	options := []string{
		`\s+`,
		`.`,
		`[0-9]+(?:\.[0-9]*)?`,
	}
	for _, s := range AMsymbols {
		options = append(options, regexp.QuoteMeta(s.input))
	}
	ptn := regexp.MustCompile(strings.Join(options, "|"))
	ptn.Longest()
	return &Scanner{
		ptn: ptn,
	}
}
