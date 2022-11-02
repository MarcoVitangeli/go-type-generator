package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/MarcoVitangeli/go-type-generator/codegenerator"
	"github.com/MarcoVitangeli/go-type-generator/json"
)

const (
	ColorGreen = "\u001b[32m"
	ColorRed   = "\u001b[31m"
)

func usage() {
	fmt.Println(ColorGreen, "Golang CLI code generator")
	fmt.Println("  with this CLI, you can load a json file, and convert it to a golang types file")
	fmt.Printf("  Usage:\n")
	fmt.Printf("    %s <json_file_path>\n", os.Args[0])
}

func printErr(msg string) {
	fmt.Println(ColorRed, msg)
}

func run() error {
	flag.Usage = usage
	flag.Parse()
	if len(os.Args) < 2 {
		return errors.New("please, provide a path to the json file that you want to parse")
	}

	t, err := json.ParseJson(os.Args[1])

	if err != nil {
		return err
	}

	w, err := codegenerator.GenerateDirAndFile()

	if err != nil {
		return err
	}

	err = codegenerator.GenerateCode(t, codegenerator.Writer{Writer: w})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		printErr(err.Error())
		os.Exit(1)
	}
}
