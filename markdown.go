package presconv

import (
	"bufio"
	"io"
)

// Markdown ...
type Markdown struct{}

// ParsePres implements the Parser interface.
func (p *Markdown) ParsePres(dst io.Writer, src io.Reader) error {
	// TODO: add error handling
	var err error
	sc := bufio.NewScanner(src)

	for sc.Scan() {
		if _, err = dst.Write(sc.Bytes()); err != nil {
			return err
		}

		if _, err = dst.Write([]byte("DEMO")); err != nil {
			return err
		}

		if _, err = dst.Write([]byte{'\n'}); err != nil {
			return err
		}
	}

	return sc.Err()
}
