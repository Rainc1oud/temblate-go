package main

import (
	"fmt"
	"os"

	tmbl "github.com/Rainc1oud/temblate-go"
)

func main() {
	var dtmpl, gotmpl string

	if len(os.Args) > 1 {
		dtmpl = os.Args[1]
	} else {
		dtmpl = "./templates/"
	}

	if len(os.Args) > 2 {
		gotmpl = os.Args[2]
	} else {
		gotmpl = "./messages.go"
	}

	if err := tmbl.Generate(gotmpl, dtmpl); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
	}
}
