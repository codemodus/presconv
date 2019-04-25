package convert

import (
	"bufio"
	"io"
	"strings"

	"github.com/codemodus/presconv"
)

// Markdown ...
type Markdown struct {
	*presconv.PresConv
}

func NewMarkdown() *Markdown {
	return &Markdown{
		PresConv: presconv.New(
			Conv(toMdH1H2),
			Conv(toMdH5H6),
			Conv(toMdH3H4),
		),
	}
}

const (
	Title       = "* "
	Bullet      = "- "
	TitleSub    = "** "
	TitleSubSub = "*** "
	Preform     = "  "
	Code        = ".code"
	Play        = ".play"
	Image       = ".image"
)

const (
	H1 = "# "      // pres title
	H2 = "## "     // pres subtitle
	H3 = "### "    // section title
	H4 = "#### "   // slide title
	H5 = "##### "  // subsection title
	H6 = "###### " // subsubsection title
)

type ConvFunc func(w *errWriter, sc *bufio.Scanner) error

func Conv(fn ConvFunc) presconv.Converter {
	return presconv.ConverterFunc(func(dst io.Writer, src io.Reader) error {
		sc := bufio.NewScanner(src)
		w := &errWriter{Writer: dst}

		return fn(w, sc)
	})
}

func passThrough(w *errWriter, sc *bufio.Scanner) error {
	for sc.Scan() {
		w.Write(sc.Bytes())
		w.Write([]byte{'\n'})
		if w.Err() != nil {
			return w.Err()
		}
	}

	return sc.Err()
}

func toMdH1H2(w *errWriter, sc *bufio.Scanner) error {
	var ct int
	for sc.Scan() {
		l := sc.Text()
		switch ct {
		case 0:
			if len(l) > 0 && l[0] != '#' {
				l = H1 + l
				ct++
			}
		case 1:
			if len(l) > 0 && l[0] != '#' {
				l = H2 + l
				ct++
			}
		default:
		}

		w.Write([]byte(l + "\n"))
		if w.Err() != nil {
			return w.Err()
		}
	}

	return sc.Err()
}

func toMdH3H4(w *errWriter, sc *bufio.Scanner) error {
	var cur string
	for sc.Scan() {
		l := sc.Text()
		if len(l) == 0 {
			if cur != "" {
				cur += "\n"
				continue
			}
		}

		if strings.HasPrefix(l, Title) {
			if cur == "" {
				cur = strings.TrimPrefix(l, Title) + "\n"
				continue
			}

			l = H3 + cur + swapPrefix(l, Title, H4)
			cur = ""
		}

		if cur != "" {
			l = H4 + cur + l
			cur = ""
		}

		w.Write([]byte(l + "\n"))
		if w.Err() != nil {
			return w.Err()
		}
	}

	return sc.Err()
}

func toMdH5H6(w *errWriter, sc *bufio.Scanner) error {
	for sc.Scan() {
		l := sc.Text()
		l = swapPrefix(l, TitleSub, H5)
		l = swapPrefix(l, TitleSubSub, H6)

		w.Write([]byte(l + "\n"))
		if w.Err() != nil {
			return w.Err()
		}
	}

	return sc.Err()
}

func swapPrefix(s, a, b string) string {
	if a == "" {
		return s
	}

	if strings.HasPrefix(s, a) {
		s = b + strings.TrimPrefix(s, a)
	}

	return s
}

func linePrefix(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			return s[0 : i+1]
		}
	}
	return ""
}
