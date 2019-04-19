package presconv

import (
	"bytes"
	"io"
)

// Converter describes types intended for processing presentation slide
// files.
type Converter interface {
	ConvertPres(io.Writer, io.Reader) error
}

// PresConv manages converters for parsing presentation slide files.
type PresConv struct {
	cs []Converter
}

// New constructs an instance of PresConv.
func New(cs ...Converter) *PresConv {
	c := PresConv{
		cs: cs,
	}

	return &c
}

// ConvertPres leverages the underlying converters queue to process presentation
// slide files.
func (c *PresConv) ConvertPres(dst io.Writer, src io.Reader) error {
	var buf *bytes.Buffer
	i := len(c.cs) - 1

	for _, cc := range c.cs[:i] {
		buf = &bytes.Buffer{}

		if err := cc.ConvertPres(buf, src); err != nil {
			return err
		}

		src = buf
	}

	return c.cs[i].ConvertPres(dst, src)
}
