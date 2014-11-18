package template

import (
	"github.com/arschles/assert"
	"testing"
  "testing/quick"
)

func TestName(t *testing.T) {
  err := quick.Check(func(name string, bytes []byte) bool {
    assetFunc := func(name string) ([]byte, error) {
      return bytes, nil
    }
    tmpl := New(name, assetFunc)
    return tmpl.Name() == name
  }, nil)
  assert.NoErr(t, err)
}
