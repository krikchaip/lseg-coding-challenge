package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var args = Args{
	Filepath: "data/logs.log",
}

type Args struct {
	Filepath string
}

func init() {
	flag.Usage = func() {
		program := path.Base(os.Args[0])
		output := "Usage of %s:\n" +
			"  %[1]s [flags] /path/to/logfile.log\n" +
			"\n"

		fmt.Fprintf(flag.CommandLine.Output(), output, program)
		flag.PrintDefaults()
	}

	flag.Parse()

	if args.Filepath = flag.Arg(0); args.Filepath == "" {
		flag.Usage()
		os.Exit(0)
	}
}
