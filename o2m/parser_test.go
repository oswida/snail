package o2m_test

import (
	"snail/o2m"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPaths(t *testing.T) {
	parser, err := o2m.NewParser("testdata/basic.yml")
	assert.NilError(t, err)
	assert.Equal(t, 3, len(parser.PathNames()))
}

func TestPathNotFound(t *testing.T) {
	parser, err := o2m.NewParser("testdata/basic.yml")
	assert.NilError(t, err)
	_, err = parser.OperationNames("not_existing")
	assert.ErrorContains(t, err, "cannot find request path")
}

func TestOperations(t *testing.T) {
	type testData struct {
		Name        string
		Path        string
		ExpectedOps []string
	}
	tests := []testData{
		{
			Name:        "1",
			Path:        "/test1",
			ExpectedOps: []string{"GET"},
		},
		{
			Name:        "2",
			Path:        "/test2",
			ExpectedOps: []string{"POST"},
		},
		{
			Name:        "3",
			Path:        "/test/test3",
			ExpectedOps: []string{"DELETE"},
		},
	}
	parser, err := o2m.NewParser("testdata/basic.yml")
	assert.NilError(t, err)
	for _, tst := range tests {
		names, err := parser.OperationNames(tst.Path)
		assert.NilError(t, err)
		assert.DeepEqual(t, tst.ExpectedOps, names)
	}

}
