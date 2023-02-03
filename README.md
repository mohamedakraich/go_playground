## Notes

- Every Go program is made up of functions contained in packages.
- To print text on the screen, use the fmt package provided by the standard library.
- Go supports characters of every language. Print text in Chinese, Japanese, Russian, or Spanish...
- With Printf and the %v format verb, values can be placed anywhere in the displayed text.
- Constants are declared with the const keyword and can’t be changed.
- Variables are declared with var and can be assigned new values while a program is running.
- The math/rand import path refers to the rand package.
- The Intn function in the rand package generates pseudorandom numbers.
- Booleans are the only values that can be used in conditions.
- Go provides branching and repetition with if, switch, and for.
- An opening curly brace { introduces a new scope that ends with a closing brace }.
- The `case` and `default` keywords also introduce a new scope even though no curly braces are involved.
- The location where a variable is declared determines which scope it’s in.
- Not only is shortcut declaration shorter, you can take it places where `var` can’t go.
- Variables declared on the same line as a `for`, `if`, or `switch` are in scope until the end of that statement.
- A wide scope is better than a narrow scope in some situations—and vice versa.

- UDP server => `func ListenPacket(network, address string) (PacketConn, error)`

## Useful Commands

- go doc net.IP
- go doc net.ParseIP
- go test -list ".*ParseIP.*" $(go env GOROOT)/src/net
- go test -run ParseIP -bench ParseIP -count=1 $(go env GOROOT)/src/net
- grep -l TestParseIP -nr $(go env GOROOT)/src/net
- go doc net.Dial
- go doc -u net.protocols
- go doc net | grep Lookup
- go test -test.v -run "Timeout|KeepAlive" -count=1 $(go env GOROOT)/src/net
- go test -list "Timeout|KeepAlive" $(go env GOROOT)/src/net
