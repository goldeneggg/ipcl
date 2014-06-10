package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jpshadowapps/ipcl/parser"
	"github.com/jpshadowapps/ipcl/writer"
	"os"
	"runtime"
)

const (
	Version = "0.1.0"
)

var file = flag.String("f", "", "Filepath listed target CIDR")
var isCsv = flag.Bool("csv", false, "Output format is csv")
var isTsv = flag.Bool("tsv", false, "Output format is tsv")
var printVersion = flag.Bool("version", false, "Print version")

func main() {
	// option parse
	flag.Parse()

	if *printVersion {
		fmt.Printf("Ipcl: version %s (%s)\n", Version, runtime.GOARCH)
		os.Exit(2)
	}

	// get source CIDRs
	cidrs, e := getCIDRs()
	if e != nil {
		fmt.Printf("%s\n", e)
		os.Exit(1)
	}

	// write
	write(cidrs)
}

func getCIDRs() ([]parser.CIDRInfo, error) {
	var cidrStrs []string
	var cidrs []parser.CIDRInfo
	ac := flag.NArg()

	switch {
	case ac == 1:
		cidrStrs = append(cidrStrs, flag.Arg(0))
	case ac > 1:
		cidrStrs = flag.Args()
	case *file != "":
		var e error
		cidrStrs, e = fromFile()
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

func fromFile() ([]string, error) {
	cidrs := make([]string, 0, 10)

	f, err := os.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
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

func write(cidrs []parser.CIDRInfo) {
	w := writer.NewWriter(*isCsv, *isTsv)
	w.Write(cidrs)
}
