package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

const EX_USAGE = 64

func usage() {
	fmt.Fprintf(os.Stderr, "usage: htmlquery expression [document]\n")
	flag.PrintDefaults()
	os.Exit(EX_USAGE)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("")
	flag.Usage = usage
	flag.Parse()

	var expr string
	r := os.Stdin
	switch flag.NArg() {
	case 2:
		var err error
		r, err = os.Open(flag.Arg(1))
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		fallthrough
	case 1:
		expr = flag.Arg(0)
	default:
		flag.Usage()
	}

	sel, err := css.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}
	node, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	for _, el := range sel.Select(node) {
		buf := new(bytes.Buffer)
		html.Render(buf, el)
		s := strings.ReplaceAll(buf.String(), "\n", " ")
		fmt.Println(s)
	}
}
