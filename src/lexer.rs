use crate::token::KeywordType;

use super::cursor::Cursor;
use super::token::{FromLexeme, OperatorType, PunctuationType, Token};

pub struct Lexer<'a> {
    cursor: Cursor<'a>,
}

// Need to specify that implementations are valid for the same lifetime as the structs lifetimes
impl<'a> Lexer<'a> {
    pub fn new(input: &'a str) -> Self {
        Self {
            cursor: Cursor::new(input),
        }
    }

    pub fn next_token(&mut self) -> Token {
        self.skip_whitespace();

        let current_char = match self.cursor.read_char() {
            None => return Token::Eof,
            Some(ch) => ch,
        };

        let current_str = &current_char.to_string();
        // Check if single character punctuation
        if let Some(punctuation) = PunctuationType::from_lexeme(current_str) {
            return Token::Punctuation(punctuation);
        }

        // Check for multi character operator
        let next_char = self.cursor.peek_char();
        if let Some(next_ch) = next_char {
            let combined_str = Self::concat_chars(&current_char, next_ch);

            if let Some(operator) = OperatorType::from_lexeme(&combined_str) {
                _ = self.cursor.read_char();

                return Token::Operator(operator);
            }
        }

        // Fall back to checking for single character operator
        if let Some(operator) = OperatorType::from_lexeme(current_str) {
            return Token::Operator(operator);
        }

        // Check for identifiers or keywords
        let mut string_literal = String::new();
        string_literal.push(current_char);

        if current_char.is_alphabetic() {
            string_literal.push_str(&self.consume_while(char::is_alphabetic));
            if let Some(keyword) = KeywordType::from_lexeme(&string_literal) {
                return Token::Keyword(keyword);
            }

            return Token::Identifier(string_literal);
        }

        // Check for numbers
        if current_char.is_numeric() {
            string_literal.push_str(&self.consume_while(char::is_numeric));

            return match string_literal.parse::<i32>() {
                Ok(parsed_number) => Token::Integer(parsed_number),
                Err(err) => {
                    println!(
                        "Failed to parse {:#?} due to error: {:#?}",
                        string_literal, err
                    );
                    return Token::Eof;
                }
            };
        }

        Token::Eof
    }

    fn skip_whitespace(&mut self) {
        while let Some(next_char) = self.cursor.peek_char() {
            if next_char.is_whitespace() {
                self.cursor.read_char();
            } else {
                break;
            }
        }
    }

    fn concat_chars(a: &char, b: &char) -> String {
        let mut combined_str = String::with_capacity(a.len_utf8() + b.len_utf8());
        combined_str.push(*a);
        combined_str.push(*b);

        combined_str
    }

    fn consume_while<F>(&mut self, condition: F) -> String
    where
        F: Fn(char) -> bool,
    {
        let mut literal = String::new();
        while let Some(next_char) = self.cursor.peek_char() {
            if condition(*next_char) {
                // Safe to unwrap, we have already peeked the char
                let consumed_char = self.cursor.read_char().unwrap();
                literal.push(consumed_char);
            } else {
                break;
            }
        }

        literal
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::token::{KeywordType, OperatorType, PunctuationType, Token};

    // Helper function to collect tokens until EOF.
    fn collect_tokens(lexer: &mut Lexer) -> Vec<Token> {
        let mut tokens = Vec::new();
        loop {
            let tok = lexer.next_token();
            if tok == Token::Eof {
                tokens.push(tok);
                break;
            }
            tokens.push(tok);
        }
        tokens
    }

    #[test]
    fn test_keyword_token() {
        let input = "let";
        let mut lexer = Lexer::new(input);
        let token = lexer.next_token();
        assert_eq!(token, Token::Keyword(KeywordType::Let));
    }

    #[test]
    fn test_identifier_token() {
        let input = "foobar";
        let mut lexer = Lexer::new(input);
        let token = lexer.next_token();
        assert_eq!(token, Token::Identifier("foobar".to_string()));
    }

    #[test]
    fn test_operator_token_single() {
        let input = "=";
        let mut lexer = Lexer::new(input);
        let token = lexer.next_token();
        assert_eq!(token, Token::Operator(OperatorType::Assign));
    }

    #[test]
    fn test_operator_token_multi() {
        let input = "==";
        let mut lexer = Lexer::new(input);
        let token = lexer.next_token();
        assert_eq!(token, Token::Operator(OperatorType::Equals));
    }

    #[test]
    fn test_integer_token() {
        let input = "12345";
        let mut lexer = Lexer::new(input);
        let token = lexer.next_token();
        assert_eq!(token, Token::Integer(12345));
    }

    #[test]
    fn test_punctuation_token() {
        let input = ",";
        let mut lexer = Lexer::new(input);
        let token = lexer.next_token();
        assert_eq!(token, Token::Punctuation(PunctuationType::Comma));
    }

    #[test]
    fn test_complex_input() {
        // An input with several tokens.
        let input = "let x = 5; fn name(param) {}";
        let mut lexer = Lexer::new(input);
        let tokens = collect_tokens(&mut lexer);

        // Define the expected token sequence.
        let expected = vec![
            Token::Keyword(KeywordType::Let),
            Token::Identifier("x".to_string()),
            Token::Operator(OperatorType::Assign),
            Token::Integer(5),
            Token::Punctuation(PunctuationType::Semicolon),
            Token::Keyword(KeywordType::Function),
            Token::Identifier("name".to_string()),
            Token::Punctuation(PunctuationType::LeftParen),
            Token::Identifier("param".to_string()),
            Token::Punctuation(PunctuationType::RightParen),
            Token::Punctuation(PunctuationType::LeftBrace),
            Token::Punctuation(PunctuationType::RightBrace),
            Token::Eof,
        ];

        assert_eq!(tokens, expected);
    }
}
