import Foundation
import LoveLangCore

func getUsername() -> String? {
    ProcessInfo.processInfo.environment["USER"]
}

func getHelloUserString() -> String {
    if let username = getUsername() {
        return "Hello, \(username)!"
    }

    return "Hello!"
}

func main() {
    print("")
    print(
        """
                @@@@@@           @@@@@@
              @@@@@@@@@@       @@@@@@@@@@
            @@@@@@@@@@@@@@   @@@@@@@@@@@@@@
          @@@@@@@@@@@@@@@@@ @@@@@@@@@@@@@@@@@
         @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
        @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
        @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
        @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
         @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
          @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
           @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
            @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
              @@@@@@@@@@@@@@@@@@@@@@@@@@@
                @@@@@@@@@@@@@@@@@@@@@@@
                  @@@@@@@@@@@@@@@@@@@
                    @@@@@@@@@@@@@@@
                      @@@@@@@@@@@
                        @@@@@@@
                          @@@
                           @
        """)
    let welcomeMessage = "\(getHelloUserString()) Welcome to the LoveLang ❤ programming language!"
    print(welcomeMessage)

    while true {

        print("> ", terminator: "")

        guard let input = readLine() else {
            print("Exiting...")
            exit(0)
        }

        if input == "exit" {
            print("Goodbye!")
            break
        }

        let lexer = Lexer(input: input)

        while true {
            let token = lexer.nextToken()
            if token == .eof { break }
            print(token)
        }
    }
}

main()
