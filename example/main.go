package main

import (
	"fmt"

	"github.com/Scorpio69t/goi18n"
)

func main() {
	// Output:
	// Goi18n.New: &{Path:main.go Language:zh}
	g := goi18n.New(&goi18n.Option{
		Path:     "i18n",
		Language: "en",
	})
	// fmt.Printf("Goi18n.New: %#v\n", g)
	// chinese
	hi := g.Translate("zh", "hi")
	world := g.Translate("zh", "world")
	fmt.Printf("%s %s!\n", hi, world)

	// english
	hi = g.T("hi")
	world = g.T("world")
	fmt.Printf("%s %s!\n", hi, world)
}
