package lexer

import (
	"luisabraham22/lovelang/core/cursor"
	"luisabraham22/lovelang/core/token"
	"unicode"
)

const EofIdentifier = rune(0)

type Lexer struct {
	cursor *cursor.Cursor
}

func New(input string) *Lexer {
	cur := cursor.New(input)
	return &Lexer{
		cur,
	}
}
func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	currChar := l.cursor.ReadChar()
	currStr := string(currChar)

	// Check if punctuation
	if tt, ok := token.LookupPunctuation(currStr); ok {
		return token.Token{
			Type:    tt,
			Literal: currStr,
		}
	}

	// Check if multi character operator
	if nextChar := l.cursor.Peek(); nextChar != EofIdentifier {
		combinedStr := currStr + string(nextChar)
		if tt, ok := token.LookupOperator(combinedStr); ok {
			// Consume next char
			_ = l.cursor.ReadChar()
			return token.Token{
				Type:    tt,
				Literal: combinedStr,
			}
		}
	}

	// Check if single character operator
	if tt, ok := token.LookupOperator(currStr); ok {
		return token.Token{
			Type:    tt,
			Literal: currStr,
		}
	}

	// check if number
	if unicode.IsDigit(currChar) {
		numberLiteral := currStr + l.consumeWhile(unicode.IsDigit)
		return token.Token{
			Type:    token.Integer,
			Literal: numberLiteral,
		}
	}

	// Check if identifier or keyword
	if unicode.IsLetter(currChar) {
		stringLiteral := currStr + l.consumeWhile(unicode.IsLetter)

		tt := token.LookupIdentifier(stringLiteral)

		return token.Token{
			Type:    tt,
			Literal: stringLiteral,
		}
	}

	return token.Token{
		Type:    token.EOF,
		Literal: "",
	}
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.cursor.Peek()) {
		_ = l.cursor.ReadChar()
	}
}

func (l *Lexer) consumeWhile(cond func(rune) bool) string {
	var runes []rune

	nextChar := l.cursor.Peek()

	for nextChar != EofIdentifier && cond(nextChar) {
		// Consume next char
		runes = append(runes, l.cursor.ReadChar())
		// peek next char
		nextChar = l.cursor.Peek()
	}

	return string(runes)
}
