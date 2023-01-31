## Notes

- Every Go program is made up of functions contained in packages.
- To print text on the screen, use the fmt package provided by the standard library.

## Useful Commands

- go doc net.IP
- go doc net.ParseIP
- go test -list ".*ParseIP.*" $(go env GOROOT)/src/net
- go test -run ParseIP -bench ParseIP -count=1 $(go env GOROOT)/src/net
- grep -l TestParseIP -nr $(go env GOROOT)/src/net
- go doc net.Dial
- go doc -u net.protocols
- go doc net | grep Lookup
