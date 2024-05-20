package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/samyouxyz/lucy/parser"
	"github.com/samyouxyz/lucy/tokenizer"
)

func run() error {
	var query string
	var url string

	flag.StringVar(&query, "query", "", "Search query")
	flag.StringVar(&url, "index", "", "Index a url")
	flag.Parse()

	text, err := parser.ParseHTMLFromURL(url)
	if err != nil {
		return err
	}

	tokenizer := tokenizer.New()
	tokens := tokenizer.Tokenize(text)
	var quotedTokens []string
	for _, token := range tokens {
		quotedTokens = append(quotedTokens, fmt.Sprintf(`"%s"`, token))
	}
	fmt.Println(strings.Join(quotedTokens, ", "))

	return nil

}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
