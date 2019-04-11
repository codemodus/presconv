package presm

import (
	"bytes"
	"io"
)

// Parser describes types intended for processing presentation slide files.
type Parser interface {
	ParsePres(io.Writer, io.Reader) error
}

// Presm manages parsers for parsing presentation slide files.
type Presm struct {
	ps []Parser
}

// New constructs an instance of Presm.
func New(ps ...Parser) *Presm {
	v := Presm{
		ps: ps,
	}

	return &v
}

// ParsePres leverages the underlying Parser queue to process presentation slide
// files.
func (p *Presm) ParsePres(dst io.Writer, src io.Reader) error {
	var buf *bytes.Buffer
	i := len(p.ps) - 1

	for _, cp := range p.ps[:i] {
		buf = &bytes.Buffer{}

		if err := cp.ParsePres(buf, src); err != nil {
			return err
		}

		src = buf
	}

	return p.ps[i].ParsePres(dst, src)
}
