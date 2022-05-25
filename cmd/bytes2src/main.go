package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dgmcdona/bytes2src"
)

func main() {
	x := flag.Bool("x", false, "interpret input data as hex dump")
	f := flag.String("f", "", "input file from which to read bytes")
	lang := flag.String("l", "", "programming language for source code output")
	w := flag.Int("w", 8, "width of the source code array output")
	o := flag.String("o", "", "output to file <path>")

	flag.Parse()

	var output *string
	var l bytes2src.Language

	switch *lang {
	case "go", "golang":
		l = bytes2src.Go
	case "javascript", "js":
		l = bytes2src.JavaScript
	case "csharp", "c#", "java":
		l = bytes2src.CSharp
	default:
		fmt.Fprintf(os.Stderr, "supply a valid output language:\n")
		fmt.Fprintf(os.Stderr, "\tjava\n")
		fmt.Fprintf(os.Stderr, "\tcsharp\n")
		fmt.Fprintf(os.Stderr, "\tgo\n")
		fmt.Fprintf(os.Stderr, "\tjavascript\n")
		os.Exit(-1)
	}

	fi, _ := os.Stdin.Stat()
	var d []byte

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading from pipe: %v\n", err)
			os.Exit(-1)
		}
		d = data
	} else if *f != "" {
		file, err := os.Open(*f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open file %s: %v\n", *f, err)
			os.Exit(-1)
		}

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read data from file: %v\n", err)
			os.Exit(-1)
		}
		d = data
	} else {
		fmt.Fprintf(os.Stderr, "must supply input from file with -f flag or from pipe\n")
		os.Exit(-1)
	}

	if *x {
		br := bytes.NewReader(d)
		var err error
		_, d, err = bytes2src.ReadHexDump(br)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading hexdump: %v\n", err)
			os.Exit(-1)
		}
	}

	br := bytes.NewReader(d)
	var err error
	output, err = bytes2src.DumpString(br, l, *w)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to convert bytes to source code: %v\n", err)
		os.Exit(-1)
	}

	if *o != "" {
		file, err := os.Open(*o)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open file for writing: %v\n", err)
		}

		if _, err = file.WriteString(*output); err != nil {
			fmt.Fprintf(os.Stderr, "error writing output to file: %v\n", err)
			os.Exit(-1)
		}
	}

	fmt.Printf("%s\n", *output)
}
