package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal    TokenType = "ILLEGAL"
	EOF                  = "EOF"
	Identifier           = "IDENTIFIER"

	Integer = "INT"
	String  = "STRING"

	// Punctuation
	Comma      TokenType = "COMMA"
	Semicolon  TokenType = "SEMICOLON"
	LeftParen  TokenType = "LEFT_PAREN"
	RightParen TokenType = "RIGHT_PAREN"
	LeftBrace  TokenType = "LEFT_BRACE"
	RightBrace TokenType = "RIGHT_BRACE"

	// Operators
	Assign              TokenType = "ASSIGN"
	Plus                TokenType = "PLUS"
	Equals              TokenType = "EQUALS"
	NotEquals           TokenType = "NOT_EQUALS"
	Not                 TokenType = "NOT"
	Asterisk            TokenType = "ASTERISK"
	Minus               TokenType = "MINUS"
	Slash               TokenType = "SLASH"
	LessThan            TokenType = "LESS_THAN"
	LessThanOrEquals    TokenType = "LESS_THAN_OR_EQUALS"
	GreaterThan         TokenType = "GREATER_THAN"
	GreaterThanOrEquals TokenType = "GREATER_THAN_OR_EQUALS"

	// Keywords
	Function TokenType = "FUNCTION"
	Let      TokenType = "LET"
	True     TokenType = "TRUE"
	False    TokenType = "FALSE"
	If       TokenType = "IF"
	Else     TokenType = "ELSE"
	Return   TokenType = "RETURN"
	For      TokenType = "FOR"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
	"for":    For,
}

var operators = map[string]TokenType{
	"=":  Assign,
	"+":  Plus,
	"==": Equals,
	"!=": NotEquals,
	"!":  Not,
	"*":  Asterisk,
	"-":  Minus,
	"/":  Slash,
	"<":  LessThan,
	"<=": LessThanOrEquals,
	">":  GreaterThan,
	">=": GreaterThanOrEquals,
}

var punctuation = map[string]TokenType{
	",": Comma,
	";": Semicolon,
	"(": LeftParen,
	")": RightParen,
	"{": LeftBrace,
	"}": RightBrace,
}

func LookupIdentifier(literal string) TokenType {
	if token, ok := keywords[literal]; ok {
		return token
	}

	return Identifier
}

func LookupOperator(literal string) (TokenType, bool) {
	t, ok := operators[literal]
	return t, ok
}

func LookupPunctuation(literal string) (TokenType, bool) {
	t, ok := punctuation[literal]
	return t, ok
}
