// Echo4 prints its command-line arguments.
// > go run Flag.go -s % joe blow
// > joe%blow
//
// > go run Flag.go -help
//Usage of /var/folders/kv/dp8glrlx5n3512fmwznfyj5m0000gn/T/go-build668593350/b001/exe/Flag:
//-n	omit trailing newline
//-s string
//separator (default " ")
//exit status 2

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
