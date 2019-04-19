package convert

import (
	"bufio"
	"io"
)

// Markdown ...
type Markdown struct{}

// ConvertPres implements the presconv.Converter interface.
func (p *Markdown) ConvertPres(dst io.Writer, src io.Reader) error {
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
