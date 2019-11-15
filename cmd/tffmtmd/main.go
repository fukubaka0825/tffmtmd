package main

import (
	"flag"
	"fmt"
	"github.com/po3rin/mdfile"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	ExitCodeOK = iota
	ExitCodeInvalidArgsError
	ExitCodeParseFlagsError
)

const (
	NAME    = "tffmtmd"
	USAGE   = `Usage: tffmtmd [options...] filePath

tffmtmd is a command line tool to format HCL code in Markdown.

OPTIONS:
  --replace value, -r value  replace HCL code with formated code
  --write value, -w value    write result to file instead of stdout
  --help, -h              prints out help
`
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var replace bool
	var writeFile string

	flags := flag.NewFlagSet(NAME, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(cli.outStream, USAGE)
	}

	flags.BoolVar(&replace, "replace", false, "")
	flags.BoolVar(&replace, "r", false, "")

	flags.StringVar(&writeFile, "write", "", "")
	flags.StringVar(&writeFile, "w", "", "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	nonFlagArgs := flags.Args()
	if len(nonFlagArgs) != 1 {
		fmt.Fprintf(cli.errStream, "Failed to set up tffmtmd: invalid argument\n"+
			"Please specify the exact one path to a file\n\n")
		return ExitCodeInvalidArgsError
	}

	filePath := nonFlagArgs[0]
	md, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read bytes from %v: ", filePath)
	}
	mdFile := mdfile.NewMdFile(&md, filePath)
	output, err := mdFile.FmtHclCodeInMd()
	if err != nil {
		log.Fatal(err)
	}

	if replace {
		err = ioutil.WriteFile(filePath, output, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to writes to %v: ", filePath)
		}
	}
	if writeFile != "" {
		err = ioutil.WriteFile(writeFile, output, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to writes to %v: ", filePath)
		}
	}
	if !replace && writeFile == "" {
		fmt.Fprint(cli.outStream, string(output))
	}

	return ExitCodeOK
}


func main() {
	cli := &CLI{
		outStream: os.Stdout, errStream: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))
}
