package tokenizer

import "testing"

func TestTokenizeText(t *testing.T) {
	text := `Example Domain Example Domain This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission. More information...`

	expected := []string{"example", "domain", "example", "domain", "domain", "use", "illustrative", "examples", "documents", "you", "may", "use", "domain", "literature", "without", "prior", "coordination", "asking", "permission", "more", "information"}

	tokenizer := New()
	tokens := tokenizer.Tokenize(text)

	if len(tokens) != len(expected) {
		t.Errorf("Expected %v tokens, got %v", len(expected), len(tokens))
	}

	for i, token := range tokens {
		if token != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], token)
		}
	}

}
