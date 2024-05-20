package tokenizer

import (
	"strings"
)

var stopWords = map[string]struct{}{
	"a": {}, "an": {}, "and": {}, "are": {}, "as": {}, "at": {}, "be": {}, "but": {}, "by": {},
	"for": {}, "if": {}, "in": {}, "into": {}, "is": {}, "it": {},
	"no": {}, "not": {}, "of": {}, "on": {}, "or": {}, "s": {}, "such": {},
	"t": {}, "that": {}, "the": {}, "their": {}, "then": {}, "there": {}, "these": {},
	"they": {}, "this": {}, "to": {}, "was": {}, "will": {}, "with": {},
	"from": {}, "were": {}, "have": {}, "has": {}, "some": {}, "any": {}, "each": {},
	"every": {}, "either": {}, "neither": {},
}

type tokenizer struct {
	// contains filtered or unexported fields
}

func New() *tokenizer {
	return &tokenizer{}
}

func (t *tokenizer) Tokenize(text string) []string {
	splits := strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !('a' <= r && r <= 'z' || '0' <= r && r <= '9')
	})
	tokens := []string{}
	for _, s := range splits {
		if _, found := stopWords[s]; !found {
			tokens = append(tokens, s)
		}
	}
	return tokens
}
