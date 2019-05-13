package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	majBranch "github.com/go-modules-by-example/goinfo-maj-branch/designers"
	majBranch2 "github.com/go-modules-by-example/goinfo-maj-branch/v2/designers"

	majSubdir "github.com/go-modules-by-example/goinfo-maj-subdir/designers"
	majSubdir2 "github.com/go-modules-by-example/goinfo-maj-subdir/v2/designers"
)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	w := func(format string, args ...interface{}) {
		fmt.Fprintf(tw, format, args...)
	}

	w(".../goinfo-maj-branch/designers.Names():\t%v\n", majBranch.Names())
	w(".../goinfo-maj-branch/v2/designers.FullNames():\t%v\n", majBranch2.FullNames())

	w(".../goinfo-maj-subdir/designers.Names():\t%v\n", majSubdir.Names())
	w(".../goinfo-maj-subdir/v2/designers.FullNames():\t%v\n", majSubdir2.FullNames())

	tw.Flush()
}
