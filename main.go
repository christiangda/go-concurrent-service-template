package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/christiangda/go-concurrent-service-template/internal/single_file"
)

var (
	fs *flag.FlagSet

	examples = map[string]func(){
		"single_file": single_file.Run,
	}

	target string
)

func init() {
	fs = flag.NewFlagSet("go-concurrent-service-template", flag.ExitOnError)

	fs.Bool("h", false, "This help")
	fs.Bool("list", false, "List available examples")
	fs.StringVar(&target, "run", "", "Name of the example to run, use -list to list available examples")
}

func main() {
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fs.Lookup("h").Value.String() == "true" {
		fs.PrintDefaults()
		os.Exit(0)
	}

	if fs.Lookup("list").Value.String() == "true" {
		fmt.Printf("Available examples name:\n")
		for k := range examples {
			fmt.Printf("  - %s\n", k)
		}
		os.Exit(0)
	}

	if fs.Lookup("run").Value.String() != "" {
		if name, ok := examples[target]; ok {
			name()
			os.Exit(0)
		} else {
			fmt.Printf("Unknown example: %s\n", target)
			fmt.Println("Use -list to list available examples")
			os.Exit(1)
		}
	}

	// catch the case where no flag is provided
	if fs.NFlag() == 0 {
		fs.PrintDefaults()
		os.Exit(0)
	}
}
