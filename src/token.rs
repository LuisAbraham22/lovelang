use crate::lexer;

pub trait FromLexeme: Sized {
    fn from_lexeme(lexeme: &str) -> Option<Self>;
}

#[derive(Debug, PartialEq)]
pub enum Token {
    Eof,

    Identifier(String),

    Integer(i32),

    Punctuation(PunctuationType),

    Operator(OperatorType),

    Keyword(KeywordType),
}

#[derive(Debug, PartialEq)]
pub enum PunctuationType {
    Comma,
    Semicolon,
    LeftParen,
    RightParen,
    LeftBrace,
    RightBrace,
}

impl FromLexeme for PunctuationType {
    fn from_lexeme(lexeme: &str) -> Option<Self> {
        match lexeme {
            "," => Some(Self::Comma),
            ";" => Some(Self::Semicolon),
            "(" => Some(Self::LeftParen),
            ")" => Some(Self::RightParen),
            "{" => Some(Self::LeftBrace),
            "}" => Some(Self::RightBrace),
            _ => None,
        }
    }
}

#[derive(Debug, PartialEq)]
pub enum OperatorType {
    Assign,
    Plus,
    Equals,
    NotEquals,
    Not,
    Asterisk,
    Minus,
    Slash,
    LessThan,
    LessThanOrEquals,
    GreaterThan,
    GreaterThanOrEquals,
}

impl FromLexeme for OperatorType {
    fn from_lexeme(lexeme: &str) -> Option<Self> {
        match lexeme {
            "=" => Some(Self::Assign),
            "+" => Some(Self::Plus),
            "==" => Some(Self::Equals),
            "!=" => Some(Self::NotEquals),
            "!" => Some(Self::Not),
            "*" => Some(Self::Asterisk),
            "-" => Some(Self::Minus),
            "/" => Some(Self::Slash),
            "<" => Some(Self::LessThan),
            "<=" => Some(Self::LessThanOrEquals),
            ">" => Some(Self::GreaterThan),
            ">=" => Some(Self::GreaterThanOrEquals),
            _ => None,
        }
    }
}

#[derive(Debug, PartialEq)]
pub enum KeywordType {
    Let,
    Function,
    If,
    Else,
    Return,
    True,
    False,
    For,
}

impl FromLexeme for KeywordType {
    fn from_lexeme(lexeme: &str) -> Option<Self> {
        match lexeme {
            "let" => Some(Self::Let),
            "fn" => Some(Self::Function),
            "if" => Some(Self::If),
            "else" => Some(Self::Else),
            "return" => Some(Self::Return),
            "true" => Some(Self::True),
            "false" => Some(Self::False),
            "for" => Some(Self::For),
            _ => None,
        }
    }
}
