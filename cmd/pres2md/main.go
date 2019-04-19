package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/codemodus/presconv"
	"github.com/codemodus/presconv/convert"
)

func main() {
	if err := run(); err != nil {
		cmd := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "%s: %s\n", cmd, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		file string
	)

	flag.StringVar(&file, "src", file, "source to process")
	flag.Parse()

	if file == "" {
		return fmt.Errorf("must set 'src' flag")
	}

	src, err := os.Open(file)
	if err != nil {
		return err
	}
	defer src.Close()

	p := presconv.New(
		&convert.DropPrefixed{
			Prefixes: []string{"*", "##"},
		},
		&convert.Markdown{},
	)

	return p.ConvertPres(os.Stdout, src)
}
