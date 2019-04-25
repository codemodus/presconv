package convert

import "io"

type errWriter struct {
	io.Writer
	err error
}

func (w *errWriter) Write(data []byte) int {
	if w.err != nil {
		return 0
	}

	n, err := w.Writer.Write(data)
	if err != nil {
		w.err = err
	}

	return n
}

func (w *errWriter) Err() error {
	return w.err
}
