package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox( "./assets")
	jokes, err := LoadFacts(&box, "jokes.json")

	if err != nil {
		fmt.Println(err)
	}

	a := App{
		Facts: jokes,
	}

	a.Initialize()
	a.Run(":8080")
}
