package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF = "EOF"

	// Punctuation
	Comma      TokenType = ","
	Semicolon  TokenType = ";"
	LeftParen  TokenType = "("
	RightParen TokenType = ")"
	LeftBrace  TokenType = "{"
	RightBrace TokenType = "}"
)

func FromLexeme(ch string) Token {
	var tokenType TokenType = EOF

	switch ch {
	case ",":
		tokenType = Comma
	case ";":
		tokenType = Semicolon
	case "(":
		tokenType = LeftParen
	case ")":
		tokenType = RightParen
	case "{":
		tokenType = LeftBrace
	case "}":
		tokenType = RightBrace
	}

	return Token{
		tokenType,
		ch,
	}
}
