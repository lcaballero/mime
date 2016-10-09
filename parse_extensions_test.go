package mime

import (
	. "github.com/lcaballero/exam/assert"
	"testing"
)

func Test_ParseExtensions_002(t *testing.T) {
	t.Log("nil map shuld produce error")
	err := parseExtensions("image/jpg", "jpg", nil)
	IsNotNil(t, err)
}

func Test_ParseExtensions_001(t *testing.T) {
	t.Log("empty extensions string should produce error")
	err := parseExtensions("image/jpg", "", nil)
	IsNotNil(t, err)
}
