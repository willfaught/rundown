# rundown

Package rundown prints package documentation as Markdown.

Usage:

	Usage of rundown:
	  -package string
	    	Package path (default ".")

The package path is passed directly to [golang.org/x/tools/go/packages.Load](https://pkg.go.dev/golang.org/x/tools/go/packages#Load).

Using rundown on itself prints this to standard output:

	# rundown

	Package rundown prints package documentation as Markdown.

	Usage:

		Usage of rundown:
		  -package string
		    	Package path (default ".")

	The package path is passed directly to [golang.org/x/tools/go/packages.Load](https://pkg.go.dev/golang.org/x/tools/go/packages#Load).

	Using rundown on itself prints this to standard output:

	[...]

Save the output to README.md:

	rundown >README.md
