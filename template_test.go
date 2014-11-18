package template

import (
	"errors"
	"github.com/arschles/assert"
	"testing"
	"testing/quick"
)

func createValidAssetFunc(name string, bytes []byte, notFound error) AssetFunc {
	return AssetFunc(func(n string) ([]byte, error) {
		if n == name {
			return bytes, nil
		}
		return []byte{}, notFound
	})
}

func TestName(t *testing.T) {
	err := quick.Check(func(name string, bytes []byte) bool {
		assertFunc := func(name string) ([]byte, error) {
			return bytes, nil
		}
		tmpl := New(name, assertFunc)
		return tmpl.Name() == name
	}, nil)
	assert.NoErr(t, err)
}

func TestParse(t *testing.T) {
	//TODO: use gogenerate to generate valid html templates
	//http://godoc.org/github.com/arschles/gogenerate

	tmpl := `
    <html>
      <head>
        <title>hello {{.name}}</title>
      </head>
      <body>
        {{.greeting}}
      </body>
    </html>
  `
	fileName := "mytmpl.tmpl"
	bytes := []byte(tmpl)
	expectedErr := errors.New("template not found")
	assetFunc := createValidAssetFunc(fileName, bytes, expectedErr)

	tmpl1, err1 := New("test", assetFunc).Parse(fileName)
	assert.NoErr(t, err1)
	assert.False(t, tmpl1 == nil, "tmpl1 was nil when it should not have been")
	tmpl2, err2 := New("test1", assetFunc).Parse(fileName + fileName)
	assert.Err(t, err2, expectedErr)
	assert.True(t, tmpl2 == nil, "tmpl2 was not nil when it should have been")

	//TODO: check actual template output
	// name := "Aaron"
}
