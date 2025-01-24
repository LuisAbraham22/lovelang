import Foundation

/// ( String in memory )
///
/// input:   [  H  |  e  |  l  |  l  |  o  ]
///    0      1      2      3      4
///      ^      ^
///      |      |
/// currentIndex ─┘      └─ nextReadIndex
///
/// currentCharacter = 'H'
public class Lexer {
    private let input: String
    private var currentIndex: String.Index
    private var nextReadIndex: String.Index
    private var currentCharacter: Character

    public init(input: String) {
        self.input = input
        self.currentIndex = input.startIndex
        self.nextReadIndex = input.startIndex
        self.currentCharacter = .eof
        advance()
    }

    public func nextToken() -> Token {
        skipWhitespace()

        // Check if single character token
        if let punctuationCharacter = PunctuationType(rawValue: String(currentCharacter)) {
            advance()
            return .punctuation(punctuationCharacter)
        }

        // Check for multi character operators
        let nextCharacter = peekCharacter()
        if nextCharacter != .eof {
            // Combine the current and next characters into a two character String
            let combinedCharacters = String(currentCharacter) + String(nextCharacter)

            if let multiCharacterOperator = OperatorType(rawValue: combinedCharacters) {
                // Consume the current character
                advance()
                // Consume the next character
                advance()
                return .operator(multiCharacterOperator)
            }
        }

        if let operatorCharacter = OperatorType(rawValue: String(currentCharacter)) {
            advance()
            return .operator(operatorCharacter)
        }

        // Check if identifier or keyword
        if currentCharacter.isLetter {
            guard let identifier = readIdentifier() else {
                return .eof
            }

            if let keyword = KeywordType(rawValue: identifier) {
                return .keyword(keyword)
            }

            return .identifier(identifier)
        }

        if currentCharacter.isWholeNumber {
            guard let number = readNumber() else {
                return .eof
            }

            guard let parsedInt = Int(number) else {
                return .eof
            }

            return .integer(parsedInt)
        }

        return .eof
    }

    private func advance() {
        if nextReadIndex >= input.endIndex {
            currentCharacter = .eof  // Update the current index to point to currently read character
            currentIndex = nextReadIndex
            // Increment next read index
            // nextReadIndex = input.index(after: nextReadIndex)
        } else {
            currentCharacter = input[nextReadIndex]
            // Update the current index to point to currently read character
            currentIndex = nextReadIndex
            // Increment next read index
            nextReadIndex = input.index(after: nextReadIndex)
        }
    }

    private func peekCharacter() -> Character {
        guard nextReadIndex < input.endIndex else {
            return .eof
        }

        return input[nextReadIndex]
    }

    // This reads letters until it finds a non letter character
    private func readIdentifier() -> String? {

        let startIndex = currentIndex
        while currentCharacter.isLetter {
            advance()
        }

        let identifierRange = startIndex..<currentIndex

        return String(input[identifierRange])
    }

    private func readNumber() -> String? {
        let startIndex = currentIndex
        while currentCharacter.isWholeNumber {
            advance()
        }

        let numberRange = startIndex..<currentIndex

        return String(input[numberRange])
    }

    private func skipWhitespace() {
        while currentCharacter.isWhitespaceOrNewline {
            advance()
        }
    }
}

extension Character {
    static var eof: Character {
        return "\0"
    }

    var isWhitespaceOrNewline: Bool {
        isWhitespace || isNewline
    }
}
