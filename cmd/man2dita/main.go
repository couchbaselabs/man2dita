package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/couchbaselabs/man2dita/parser"
	"github.com/couchbaselabs/man2dita/tokenizer"
)

func main() {
	infile := flag.String("man-file", "", "The man input file")
	outdir := flag.String("output-dir", "", "The output directory")
	title := flag.String("title", "", "The title of the output file")

	flag.Parse()

	if *infile == "" {
		fmt.Printf("No man file specified for conversion\n")
		os.Exit(1)
	}

	if *outdir == "" {
		fmt.Printf("No output directory specified\n")
		os.Exit(1)
	}

	if *title == "" {
		fmt.Printf("No title specified\n")
		os.Exit(1)
	}

	_, fname := filepath.Split(*infile)
	outfile := strings.Split(fname, ".")[0] + ".dita"

	dita, err := CreateDitaFile(filepath.Join(*outdir, outfile))
	checkError(err)

	tokens, err := tokenizer.CreateTokenizer(*infile)
	checkError(err)
	manParser := parser.CreateParser(tokens)
	root := manParser.Parse(*title)

	root.Execute(dita)
}

func checkError(err error) {
	if err != nil {
		fatalError(err)
	}
}

func fatalError(err error) {
	pc, _, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Printf("%s\n", err.Error())
	} else {
		fnname := ""
		if fn := runtime.FuncForPC(pc); fn != nil {
			fnname = fn.Name()
		}
		fmt.Printf("%s - %d: %s\n", fnname, line, err.Error())
	}

	os.Exit(1)
}

type DitaFile struct {
	outfile *os.File
	indent  int
}

func CreateDitaFile(outfile string) (*DitaFile, error) {
	out, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &DitaFile{
		outfile: out,
		indent:  0,
	}, nil
}

func (d *DitaFile) Write(data string, newline bool) {
	prefix := ""
	if newline {
		prefix = "\n" + strings.Repeat(" ", d.indent)
	}

	_, err := d.outfile.WriteString(prefix + data)
	checkError(err)
}

func (d *DitaFile) Indent() {
	d.indent += 2
}

func (d *DitaFile) Unindent() {
	d.indent -= 2
}
