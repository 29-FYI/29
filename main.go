package main

import (
	"fmt"
	"os"

	"github.com/29-FYI/twentynine"
)

func main() {
	links, err := twentynine.GetLinks()
	if err != nil {
		fmt.Fprintln(os.Stderr, "29: "+err.Error())
		os.Exit(1)
	}

	for _, link := range links {
		fmt.Printf("\x1b[1;38;5;7m%s\x1b[0m %s\n", link.Headline, link.URL)
	}
}
