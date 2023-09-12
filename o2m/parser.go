package o2m

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Parser struct {
	doc *openapi3.T

	Title string
}

func NewParser(filename string) (*Parser, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(filename)
	if err != nil {
		return nil, err
	}
	title := "OpenAPI Documentation"
	if doc.Info != nil {
		title = doc.Info.Title
	}
	retv := &Parser{
		doc:   doc,
		Title: title,
	}
	return retv, nil
}

func (p *Parser) PathNames() []string {
	keys := make([]string, 0, len(p.doc.Paths))
	for k := range p.doc.Paths {
		keys = append(keys, k)
	}
	return keys
}

func (p *Parser) OperationNames(pathKey string) ([]string, error) {
	pth, ok := p.doc.Paths[pathKey]
	if !ok {
		return nil, fmt.Errorf("cannot find request path %v", pathKey)
	}
	keys := make([]string, 0, len(pth.Operations()))
	for k := range pth.Operations() {
		keys = append(keys, k)
	}
	return keys, nil
}
