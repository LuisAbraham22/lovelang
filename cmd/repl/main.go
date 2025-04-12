package main

import (
	"fmt"
	"luisabraham22/lovelang/core/lexer"
	"luisabraham22/lovelang/core/token"
)

func main() {

	l := lexer.New("let x = fn(str, foo) {}")

	for {
		tok := l.NextToken()

		fmt.Printf("%s\n", tok)

		if tok.Type == token.EOF {
			break

		}
	}

}
