package main

import (
	"fmt"
	"os"

	"github.com/taganaka/go-cracklib"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <password>\n", os.Args[0])
		os.Exit(-1)
	}
	if message, ok := cracklib.FascistCheck(os.Args[1]); !ok {
		fmt.Println(message)
		os.Exit(-1)
	} else {
		fmt.Println("Your password looks good!")
	}
}
