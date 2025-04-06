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

        Token::Eof
    }

    fn concat_chars(a: &char, b: &char) -> String {
        let mut combined_str = String::with_capacity(a.len_utf8() + b.len_utf8());
        combined_str.push(*a);
        combined_str.push(*b);

        combined_str
    }
}
