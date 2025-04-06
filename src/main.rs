use lexer::Lexer;
use token::Token;

mod cursor;
mod lexer;
mod token;

fn main() {
    let input = String::from("=!+!=,;(){}");

    let mut lexer = Lexer::new(&input);

    loop {
        let token = lexer.next_token();

        if token == Token::Eof {
            println!("EOF");
            break;
        }

        println!("{:#?}", token);
    }
}
