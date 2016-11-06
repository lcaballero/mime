package mime

import (
	"bytes"
	. "github.com/lcaballero/exam/assert"
	"io/ioutil"
	"testing"
)

func Test_ParseExtensionLookup_001(t *testing.T) {
	t.Log("should mime-types for given extensions")
	bin, _ := ioutil.ReadFile(".files/mime-types.txt")

	r := bytes.NewReader(bin)
	lookup, _ := ParseExtensionLookup(r)

	extensions := []string{
		"atom", "rss", "rdf", "xml", "appcache",
		"avi", "ogg", "ico", "woff", "jar", "torrent",
		"rar", "zip", "wml", "shtml", "htc", "pl",
		"css", "js", "png",
	}

	set := make(map[string]struct{})

	for _, name := range extensions {
		mime, ok := lookup[name]
		IsTrue(t, ok)
		GreaterThan(t, len(mime), 0)
		_, ok = set[name]
		IsFalse(t, ok)
		set[name] = struct{}{}
	}
}
