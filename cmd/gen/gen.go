package main

import (
	"fmt"
	"os"

	tmbl "github.com/ppenguin/temblate-go"
)

func main() {
	if err := tmbl.Generate("./messages.go", "./templates/"); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
	}
}
