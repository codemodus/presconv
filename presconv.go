package presconv

import (
	"bytes"
	"io"
)

// Parser describes types intended for processing presentation slide files.
type Parser interface {
	ParsePres(io.Writer, io.Reader) error
}

// PresConv manages parsers for parsing presentation slide files.
type PresConv struct {
	ps []Parser
}

// New constructs an instance of PresConv.
func New(ps ...Parser) *PresConv {
	v := PresConv{
		ps: ps,
	}

	return &v
}

// ParsePres leverages the underlying Parser queue to process presentation slide
// files.
func (p *PresConv) ParsePres(dst io.Writer, src io.Reader) error {
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
