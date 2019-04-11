package presconv

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

// ParsePres implements the Parser interface.
func (p *DropPrefixed) ParsePres(dst io.Writer, src io.Reader) error {
	// TODO: add error handling
	sc := bufio.NewScanner(src)

	for sc.Scan() {
		s := sc.Text()
		var found bool

		for _, pfx := range p.Prefixes {
			if strings.HasPrefix(s, pfx) {
				found = true
				break
			}
		}

		if found {
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
