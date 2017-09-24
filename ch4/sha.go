// Hashes the first argument string with sha256.
// Optionally uses sha384 or sha512 if passed in as a second argument.

package main

import "os"
import "fmt"
import "crypto/sha256"
import "crypto/sha512"

func main() {
	arg := os.Args[1]
	fmt.Println(arg)

	if (len(os.Args) > 2 && os.Args[2] == "sha384") {
		argSha := sha512.Sum384([]byte(arg))
		fmt.Printf("%x", argSha)
	} else if (len(os.Args) > 2 && os.Args[2] == "sha512") {
		argSha := sha512.Sum512([]byte(arg))
		fmt.Printf("%x", argSha)
	} else {
		argSha := sha256.Sum256([]byte(arg))
		fmt.Printf("%x", argSha)
	}
}
