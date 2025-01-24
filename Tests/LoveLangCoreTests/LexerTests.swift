import Testing

@testable import LoveLangCore

struct LexerTests {
    @Test
    func nextToken() {
        let input = "=+(){},;"
        let expectedTokens = [
            Token.operator(.assign),
            Token.operator(.plus),
            Token.punctuation(.leftParen),
            Token.punctuation(.rightParen),
            Token.punctuation(.leftBrace),
            Token.punctuation(.rightBrace),
            Token.punctuation(.comma),
            Token.punctuation(.semicolon),
        ]

        let lexer = Lexer(input: input)

        for expectedToken in expectedTokens {
            let nextToken = lexer.nextToken()

            #expect(expectedToken == nextToken)
        }
    }

    @Test
    func nextToken_withIdentifiersAndNumbers() {
        let input = """
            let numFive = 5;
            let numTen = 10;

            let add = func(x, y) {
                x + y;
            };

            let result = add(numFive, numTen);
            """
        let expectedTokens = [
            Token.keyword(.let),
            Token.identifier("numFive"),
            Token.operator(.assign),
            Token.integer(5),
            Token.punctuation(.semicolon),
            Token.keyword(.let),
            Token.identifier("numTen"),
            Token.operator(.assign),
            Token.integer(10),
            Token.punctuation(.semicolon),
            Token.keyword(.let),
            Token.identifier("add"),
            Token.operator(.assign),
            Token.keyword(.function),
            Token.punctuation(.leftParen),
            Token.identifier("x"),
            Token.punctuation(.comma),
            Token.identifier("y"),
            Token.punctuation(.rightParen),
            Token.punctuation(.leftBrace),
            Token.identifier("x"),
            Token.operator(.plus),
            Token.identifier("y"),
            Token.punctuation(.semicolon),
            Token.punctuation(.rightBrace),
            Token.punctuation(.semicolon),
            Token.keyword(.let),
            Token.identifier("result"),
            Token.operator(.assign),
            Token.identifier("add"),
            Token.punctuation(.leftParen),
            Token.identifier("numFive"),
            Token.punctuation(.comma),
            Token.identifier("numTen"),
            Token.punctuation(.rightParen),
            Token.punctuation(.semicolon),
        ]
        let lexer = Lexer(input: input)

        for expectedToken in expectedTokens {
            let token = lexer.nextToken()
            #expect(expectedToken == token)
        }
    }

    @Test
    func nextToken_withMultiCharacterOperators() {
        let input = """
              if (numResults <= 100) {
                return true;
              } else if (numResults != 200) {
                return false;
              } else if (numResults == 50) {
                return true;
              }

              return !*/->>=<=;
            """
        let lexer = Lexer(input: input)

        let expectedTokens = [
            Token.keyword(.if),
            Token.punctuation(.leftParen),
            Token.identifier("numResults"),
            Token.operator(.lessThanOrEquals),
            Token.integer(100),
            Token.punctuation(.rightParen),
            Token.punctuation(.leftBrace),
            Token.keyword(.return),
            Token.keyword(.true),
            Token.punctuation(.semicolon),
            Token.punctuation(.rightBrace),
            Token.keyword(.else),
            Token.keyword(.if),
            Token.punctuation(.leftParen),
            Token.identifier("numResults"),
            Token.operator(.notEquals),
            Token.integer(200),
            Token.punctuation(.rightParen),
            Token.punctuation(.leftBrace),
            Token.keyword(.return),
            Token.keyword(.false),
            Token.punctuation(.semicolon),
            Token.punctuation(.rightBrace),
            Token.keyword(.else),
            Token.keyword(.if),
            Token.punctuation(.leftParen),
            Token.identifier("numResults"),
            Token.operator(.equals),
            Token.integer(50),
            Token.punctuation(.rightParen),
            Token.punctuation(.leftBrace),
            Token.keyword(.return),
            Token.keyword(.true),
            Token.punctuation(.semicolon),
            Token.punctuation(.rightBrace),
            Token.keyword(.return),
            Token.operator(.not),  // '!'
            Token.operator(.asterisk),  // '*'
            Token.operator(.slash),  // '/'
            Token.operator(.minus),  // '-'
            Token.operator(.greaterThan),  // '>'
            Token.operator(.greaterThanOrEquals),  // '>='
            Token.operator(.lessThanOrEquals),  // '<='
            Token.punctuation(.semicolon),
        ]

        for expectedToken in expectedTokens {
            let token = lexer.nextToken()
            #expect(expectedToken == token)
        }
    }
}
