package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/goldeneggg/ipcl/parser"
	"github.com/goldeneggg/ipcl/writer"
	"github.com/jessevdk/go-flags"
)

const (
	Version = "0.1.0"
)

// element names need to Uppercase
type options struct {
	Help    bool   `short:"h" long:"help" description:"Show help message"` // not "help" but "Help", because cause error using "-h" option
	File    string `short:"f" long:"file" description:"Filepath listed target CIDR"`
	IsCsv   bool   `short:"c" long:"csv" description:"Output format is csv"`
	IsTsv   bool   `short:"t" long:"tsv" description:"Output format is tsv"`
	Version bool   `short:"v" long:"version" description:"Print version"`
}

type optArgs struct {
	opts *options
	args []string
}

func main() {
	var status int
	// handler for return
	defer func() { os.Exit(status) }()

	// parse option args
	opts := &options{}
	parser := flags.NewParser(opts, flags.PrintErrors)
	args, err := parser.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		printHelp()
		status = 1
		return
	}

	// print help
	if opts.Help {
		printHelp()
		return
	}

	// print version
	if opts.Version {
		fmt.Fprintf(os.Stderr, "Ipcl: version %s (%s)\n", Version, runtime.GOARCH)
		return
	}

	// get source CIDRs
	oa := &optArgs{opts, args}
	cidrs, e := getCIDRs(oa)
	if e != nil {
		fmt.Printf("%s\n", e)
		printHelp()
		status = 1
		return
	}

	// write
	write(cidrs, oa)
}

func getCIDRs(oa *optArgs) ([]parser.CIDRInfo, error) {
	var cidrStrs []string
	var cidrs []parser.CIDRInfo
	ac := len(oa.args)

	switch {
	case ac == 1:
		cidrStrs = append(cidrStrs, oa.args[0])
	case ac > 1:
		cidrStrs = oa.args
	case oa.opts.File != "":
		var e error
		cidrStrs, e = fromFile(oa)
		if e != nil {
			return cidrs, e
		}
	default:
		return cidrs, fmt.Errorf("Target CIDR(or CIDR list file) is not assigned\n")
	}

	for i, cs := range cidrStrs {
		c, e := parser.Parse(cs)
		if e != nil {
			fmt.Fprintf(os.Stderr, "CIDR string[%d] %s validate error: %s\n", i, cs, e)
		} else {
			cidrs = append(cidrs, c)
		}
	}

	return cidrs, nil
}

func fromFile(oa *optArgs) ([]string, error) {
	cidrs := make([]string, 0, 10)

	f, err := os.Open(oa.opts.File)
	if err != nil {
		return cidrs, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cidrs = append(cidrs, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return cidrs, err
	}

	return cidrs, nil
}

func write(cidrs []parser.CIDRInfo, oa *optArgs) {
	w := writer.NewWriter(oa.opts.IsCsv, oa.opts.IsTsv)
	w.Write(cidrs)
}

func printHelp() {
	h := `
Usage:
  ipcl [OPTIONS] <CIDR TEXT | -f <FILE>>

Application Options:
  -f, --file=    Filepath listed target CIDR
  -c, --csv=     Output format is csv
  -t, --tsv=     Output format is tsv
  -v, --version  Print version

Help Options:
  -h, --help     Show this help message
`
	os.Stderr.Write([]byte(h))
}
