package o2m

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type Parser struct {
	Doc *openapi3.T
}

func NewParser(filename string) (*Parser, error) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	doc, err := loader.LoadFromFile(filename)
	if err != nil {
		return nil, err
	}
	retv := &Parser{
		Doc: doc,
	}
	return retv, nil
}
