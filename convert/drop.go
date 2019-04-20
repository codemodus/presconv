package convert

import (
	"bufio"
	"io"
	"strings"
)

// DropPrefixed stores prefixes used to determine which lines to drop from a
// presentation slide file.
type DropPrefixed struct {
	Prefixes []string
}

// ConvertPres implements the presconv.Converter interface.
func (p *DropPrefixed) ConvertPres(dst io.Writer, src io.Reader) error {
	// TODO: add error handling
	sc := bufio.NewScanner(src)

	for sc.Scan() {
		s := sc.Text()

		if hasAnyPrefix(s, p.Prefixes) {
			continue
		}

		if _, err := dst.Write([]byte(s)); err != nil {
			return err
		}

		if _, err := dst.Write([]byte{'\n'}); err != nil {
			return err
		}
	}

	return sc.Err()
}

func hasAnyPrefix(s string, prefixes []string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}

	return false
}
