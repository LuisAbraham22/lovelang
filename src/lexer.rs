use std::{iter::Peekable, str::Chars};

use crate::token::{FromLexeme, OperatorType, PunctuationType, Token};

pub struct Lexer<'a> {
    input: Peekable<Chars<'a>>,
}

// Need to specify that implementations are valid for the same lifetime as the structs lifetimes
impl<'a> Lexer<'a> {
    pub fn new(input: &'a str) -> Self {
        Self {
            input: input.chars().peekable(),
        }
    }

    fn read_char(&mut self) -> Option<char> {
        self.input.next()
    }

    fn peek_char(&mut self) -> Option<&char> {
        self.input.peek()
    }

    pub fn next_token(&mut self) -> Token {
        let current_char = match self.read_char() {
            None => return Token::Eof,
            Some(ch) => ch,
        };

        let current_str = &current_char.to_string();
        // Check if single character punctuation
        if let Some(punctuation) = PunctuationType::from_lexeme(current_str) {
            return Token::Punctuation(punctuation);
        }

        // Check for multi character operator
        let next_char = self.peek_char();
        if let Some(next_ch) = next_char {
            let mut combined_str =
                String::with_capacity(current_char.len_utf8() + next_ch.len_utf8());
            combined_str.push(current_char);
            combined_str.push(*next_ch);

            if let Some(operator) = OperatorType::from_lexeme(&combined_str) {
                _ = self.read_char();

                return Token::Operator(operator);
            }
        }

        // Fall back to checking for single character operator
        if let Some(operator) = OperatorType::from_lexeme(current_str) {
            return Token::Operator(operator);
        }

        Token::Eof
    }
}
