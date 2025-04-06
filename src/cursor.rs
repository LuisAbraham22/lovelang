use std::{iter::Peekable, str::Chars};

pub struct Cursor<'a> {
    input: Peekable<Chars<'a>>,
}

impl<'a> Cursor<'a> {
    pub fn new(input: &'a str) -> Self {
        Self {
            input: input.chars().peekable(),
        }
    }

    pub fn read_char(&mut self) -> Option<char> {
        self.input.next()
    }

    pub fn peek_char(&mut self) -> Option<&char> {
        self.input.peek()
    }
}
