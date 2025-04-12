package lexer

import (
	"fmt"
	"testing"

	"luisabraham22/lovelang/core/token"
)

func TestNextToken_PunctuationAndSingleCharOperators(t *testing.T) {
	input := ",;(){}=+-!*/<>"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Comma, ","},
		{token.Semicolon, ";"},
		{token.LeftParen, "("},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.RightBrace, "}"},
		{token.Assign, "="},
		{token.Plus, "+"},
		{token.Minus, "-"},
		{token.Not, "!"},
		{token.Asterisk, "*"},
		{token.Slash, "/"},
		{token.LessThan, "<"},
		{token.GreaterThan, ">"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Printf("tok: %s\n", tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] wrong token type: expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] wrong literal: expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}

func TestNextToken_MultiCharOperators(t *testing.T) {
	input := "== != <= >="
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Equals, "=="},
		{token.NotEquals, "!="},
		{token.LessThanOrEquals, "<="},
		{token.GreaterThanOrEquals, ">="},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] wrong token type: expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] wrong literal: expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

	eof := l.NextToken()
	if eof.Type != token.EOF {
		t.Fatalf("expected EOF, got %q", eof.Type)
	}
}
func TestNextToken_Numbers(t *testing.T) {
	input := "123 45"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Integer, "123"},
		{token.Integer, "45"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] wrong token type: expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] wrong literal: expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}

func TestNextToken_IdentifiersAndKeywords(t *testing.T) {
	input := "let five foobar if else return for true false"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Identifier, "foobar"},
		{token.If, "if"},
		{token.Else, "else"},
		{token.Return, "return"},
		{token.For, "for"},
		{token.True, "true"},
		{token.False, "false"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] wrong token type: expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] wrong literal: expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}

func TestNextToken_MixedInput(t *testing.T) {
	input := `let x = 10 + 20; if x >= 30 { return true; } else { return false; }`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Identifier, "x"},
		{token.Assign, "="},
		{token.Integer, "10"},
		{token.Plus, "+"},
		{token.Integer, "20"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.Identifier, "x"},
		{token.GreaterThanOrEquals, ">="},
		{token.Integer, "30"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Else, "else"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] wrong token type: expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] wrong literal: expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
