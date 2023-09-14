package main

import (
	"fmt"
	"snail/internal/tpl"
	"snail/o2m"
)

func main() {
	printer := o2m.NewPrinter(tpl.TemplateFS)
	data, err := printer.Print("test.yml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", data)
}
