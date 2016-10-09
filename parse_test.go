package mime

import (
	"bytes"
	. "github.com/lcaballero/exam/assert"
	"io/ioutil"
	"testing"
)

func Test_Parse_009(t *testing.T) {
	t.Log("should find mime and all types in full file")
	bin, _ := ioutil.ReadFile(".files/mime-types.txt")

	r := bytes.NewReader(bin)
	mimes, err := Parse(r)
	IsNil(t, err)
	IsNotNil(t, mimes)

	names := []string{
		"application/javascript", "text/css", "text/html",
		"text/x-component", "font/opentype", "image/x-icon",
		"video/x-msvideo", "text/cache-manifest", "application/xml",
	}

	for _,name := range names {
		exts, ok := mimes[name]
		IsTrue(t, ok)
		if len(exts) == 0 {
			t.Logf("name: %s, exts: %v", name, exts)
		}
		GreaterThan(t, len(exts), 0)
	}
}

func Test_Parse_008(t *testing.T) {
	t.Log("should multiple extensions for audio/midi")
	bin, _ := ioutil.ReadFile(".files/mime-2.txt")

	r := bytes.NewReader(bin)
	mimes, err := Parse(r)
	IsNil(t, err)
	IsNotNil(t, mimes)

	exts, ok := mimes["audio/midi"]
	IsTrue(t, ok)

	IsEqInt(t, len(exts), 3)
	IsEqStrings(t, exts[0], "mid")
	IsEqStrings(t, exts[1], "midi")
	IsEqStrings(t, exts[2], "kar")
}

func Test_Parse_007(t *testing.T) {
	t.Log("should find mime and multiple extensions for audio/midi")
	bin, _ := ioutil.ReadFile(".files/mime-2.txt")

	r := bytes.NewReader(bin)
	mimes, err := Parse(r)
	IsNil(t, err)
	IsNotNil(t, mimes)

	exts, ok := mimes["audio/midi"]
	IsTrue(t, ok)

	IsEqInt(t, len(exts), 3)
}

func Test_Parse_006(t *testing.T) {
	t.Log("should find mime and extension for audio/midi")
	bin, _ := ioutil.ReadFile(".files/mime-2.txt")

	r := bytes.NewReader(bin)
	mimes, err := Parse(r)
	IsNil(t, err)
	IsNotNil(t, mimes)

	_, ok := mimes["audio/midi"]
	IsTrue(t, ok)
}

func Test_Parse_005(t *testing.T) {
	t.Log("should find a improperly formed extension pair")
	bin, _ := ioutil.ReadFile(".files/mime-1b-malformed.txt")

	r := bytes.NewReader(bin)
	_, err := Parse(r)
	IsNotNil(t, err)
}

func Test_Parse_004(t *testing.T) {
	t.Log("should find a malformed pair")
	bin, _ := ioutil.ReadFile(".files/mime-1a-malformed.txt")

	r := bytes.NewReader(bin)
	_, err := Parse(r)
	IsNotNil(t, err)
}

func Test_Parse_003(t *testing.T) {
	t.Log("should find 1 mime and extension in mime-1.txt")
	bin, err := ioutil.ReadFile(".files/mime-1.txt")
	IsNil(t, err)

	r := bytes.NewReader(bin)
	mimes, err := Parse(r)
	IsNil(t, err)
	t.Log(mimes)
	IsEqInt(t, len(mimes), 1)
}

func Test_Parse_002(t *testing.T) {
	t.Log("should produce empty map if file is empty")
	bin, err := ioutil.ReadFile(".files/mime-empty.txt")
	IsNil(t, err)

	r := bytes.NewReader(bin)
	mimes, err := Parse(r)
	IsNil(t, err)
	IsZero(t, len(mimes))
}

func Test_Parse_001(t *testing.T) {
	t.Log("should get error for nil reader")
	mimes, err := Parse(nil)
	IsNotNil(t, err)
	IsNil(t, mimes)
}
