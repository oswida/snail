package o2m_test

import (
	"snail/internal/tpl"
	"snail/o2m"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPrintBasic(t *testing.T) {
	printer := o2m.NewPrinter(tpl.TemplateFS)
	_, err := printer.Print("testdata/basic.yml")
	assert.NilError(t, err)
	// assert.Equal()
}
