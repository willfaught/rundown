// Package rundown prints package documentation as Markdown.
//
// Usage:
/*
	Usage of rundown:
	  -package string
	    	Package path (default ".")
*/
// The package path is passed directly to [golang.org/x/tools/go/packages.Load].
//
// Using rundown on itself prints this to standard output:
//
//     # rundown
//
//     Package rundown prints package documentation as Markdown.
//
//     Usage:
//
//     	Usage of rundown:
//     	  -package string
//     	    	Package path (default ".")
//
//     The package path is passed directly to [golang.org/x/tools/go/packages.Load](https://pkg.go.dev/golang.org/x/tools/go/packages#Load).
//
//     Using rundown on itself prints this to standard output:
//
//     [...]
//
// Save the output to README.md:
//
//     rundown >README.md
package main

import (
	"flag"
	"fmt"
	"go/doc"
	"go/doc/comment"
	"log"
	"path"

	"github.com/willfaught/forklift"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("rundown: ")
	pkg := flag.String("package", ".", "Package path")
	flag.Parse()
	fp, err := forklift.LoadPackage(*pkg)
	if err != nil {
		log.Fatalf("cannot load package: %v", err)
	}
	if fp == nil {
		log.Fatalf("package does not exist")
	}
	dp, err := doc.NewFromFiles(fp.Fset, fp.Syntax, fp.PkgPath)
	if err != nil {
		log.Fatalf("cannot load documentation: %v", err)
	}
	printer := dp.Printer()
	printer.HeadingID = func(h *comment.Heading) string { return "" }
	printer.HeadingLevel = 2
	const base = "https://pkg.go.dev/"
	printer.DocLinkURL = func(l *comment.DocLink) string {
		importPath := l.ImportPath
		if importPath == "" {
			importPath = fp.PkgPath
		}
		switch {
		case l.Name == "":
			return base + importPath
		case l.Recv != "":
			return base + importPath + "#" + l.Recv + "." + l.Name
		default:
			return base + importPath + "#" + l.Name
		}
	}
	name := fp.Name
	if name == "main" {
		name = path.Base(fp.PkgPath)
	}
	fmt.Printf("# %s\n", name)
	parser := comment.Parser{LookupSym: func(recv, name string) (ok bool) { return true }}
	if md := string(printer.Markdown(parser.Parse(dp.Doc))); len(md) > 0 {
		fmt.Println()
		fmt.Print(md)
	}
}
