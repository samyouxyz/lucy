package parser

import "testing"

func TestParseHTMLFromURL(t *testing.T) {

	url := "https://example.com/"

	expected := `Example Domain Example Domain This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission. More information...`

	text, err := ParseHTMLFromURL(url)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if text != expected {
		t.Errorf("Expected\n%v, got\n%v", expected, text)
	}

}
