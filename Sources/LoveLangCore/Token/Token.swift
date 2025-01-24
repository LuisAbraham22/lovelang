import Foundation

public enum Token: Equatable {
    case eof

    case identifier(String)
    case integer(Int)

    case punctuation(PunctuationType)
    case `operator`(OperatorType)

    case keyword(KeywordType)

}

public enum PunctuationType: String {
    case comma = ","
    case semicolon = ";"
    case leftParen = "("
    case rightParen = ")"
    case leftBrace = "{"
    case rightBrace = "}"
}

public enum OperatorType: String {
    case assign = "="
    case plus = "+"
    case equals = "=="
    case notEquals = "!="
    case not = "!"
    case asterisk = "*"
    case minus = "-"
    case slash = "/"
    case lessThan = "<"
    case lessThanOrEquals = "<="
    case greaterThan = ">"
    case greaterThanOrEquals = ">="
}

public enum KeywordType: String {
    case `let`
    case function = "func"
    case `if`
    case `else`
    case `return`
    case `true`
    case `false`
    case `throws`
}
