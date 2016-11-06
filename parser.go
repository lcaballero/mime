package mime

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ToExtensions is mapping from the mime-type to file extensions.
type ToExtensions map[string][]string

// ExtensionToType maps the extension to its mime-type.
type ExtensionToType map[string]string

// LoadMimeTypes reads the given file to produce a map of extensions
// to the mime-types appropriate for html content-type headers.
func LoadMimeTypes(filename string) (ExtensionToType, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("mime-type filename is directory")
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ext, err := ParseExtensionLookup(f)
	return ext, err
}

// ParseExtensionLookup parses over the given reader producing
// an extension to mime-type lookup.  An error is produced
// if the reader is nil or it find a malformed file.
func ParseExtensionLookup(reader io.Reader) (ExtensionToType, error) {
	mimes, err := Parse(reader)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	for mime, extensions := range mimes {
		for _, ext := range extensions {
			m[ext] = mime
		}
	}
	return m, nil
}

// Parse parses over the given reader producing a map of mime-types
// to file extensions for that type.  An error is produced when
// the reader is nil or the file is malformed.
func Parse(reader io.Reader) (ToExtensions, error) {
	if reader == nil {
		return nil, fmt.Errorf("cannot parse mimes given nil reader")
	}
	br := bufio.NewReader(reader)
	m := make(map[string][]string)
	err := parse(br, m)

	return m, err
}

func isNonPair(line string) bool {
	return strings.HasSuffix(line, "{") ||
		strings.HasSuffix(line, "}") ||
		strings.HasPrefix(line, "#") ||
		len(line) == 0
}

func parse(r *bufio.Reader, m ToExtensions) error {

	var err error
	var line string

	n := 0

	for err == nil {
		line, err = r.ReadString('\n')
		n++
		line = strings.TrimSpace(line)

		if isNonPair(line) {
			continue
		}

		if strings.HasSuffix(line, ";") {
			name, values, parseerr := parsePair(n, line)
			if parseerr != nil {
				err = parseerr
			} else {
				err = parseExtensions(name, values, m)
			}
		} else {
			err = fmt.Errorf("non-comment, non-pair at line %d, %s", n, line)
		}
	}
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func parseExtensions(name, exts string, m ToExtensions) error {
	if exts == "" {
		return fmt.Errorf("cannot parse extension from empty string")
	}
	if m == nil {
		return fmt.Errorf("cannot fill nil map")
	}
	m[name] = strings.Split(exts, " ")
	return nil
}

func parsePair(n int, line string) (string, string, error) {
	startWS := strings.Index(line, " ")
	name := line[0:startWS]

	// Remove left-most WS after the key
	line = strings.TrimSpace(line[startWS:])
	semiIndex := strings.LastIndex(line, ";")

	values := line[:semiIndex]

	if strings.Contains(values, ";") {
		return "", "", fmt.Errorf("improperly formed extension, @line: %d\n%s", n, line)
	}
	return name, values, nil
}
