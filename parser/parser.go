package parser

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func ExtractVisibleText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style" || n.Data == "noscript") {
		return ""
	}

	var result strings.Builder
	var prevEndsWithSpace bool
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text := ExtractVisibleText(c)
		if len(text) > 0 {
			if prevEndsWithSpace || result.Len() == 0 {
				result.WriteString(text)
			} else {
				result.WriteByte(' ')
				result.WriteString(text)
			}
			prevEndsWithSpace = text[len(text)-1] == ' '
		}
	}
	return result.String()
}

func ParseHTMLFromURL(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	visibleText := ExtractVisibleText(doc)
	return visibleText, nil
}
