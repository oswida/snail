package o2m

import (
	"bytes"
	"embed"
	"text/template"
)

type Printer struct {
	fs embed.FS
}

func NewPrinter(fs embed.FS) *Printer {
	return &Printer{
		fs: fs,
	}
}

func (p *Printer) Print(filename string) (string, error) {
	mainTemplate, err := template.ParseFS(p.fs, "*.gotmpl")
	if err != nil {
		return "", err
	}
	parser, err := NewParser(filename)
	if err != nil {
		return "", err
	}
	var buff bytes.Buffer
	err = mainTemplate.ExecuteTemplate(&buff, "main.gotmpl", parser.Doc)
	return buff.String(), err
}
