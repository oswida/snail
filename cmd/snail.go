package main

import (
	"fmt"
	"snail/internal/tpl"
	"snail/o2m"
)

func main() {
	printer := o2m.NewPrinter(tpl.TemplateFS)
	data, err := printer.Print("o2m/testdata/basic.yml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", data)
}
