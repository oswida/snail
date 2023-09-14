package tpl

import "embed"

//go:embed *.gotmpl
var TemplateFS embed.FS
