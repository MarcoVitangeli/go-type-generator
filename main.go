package main

import (
	"fmt"
	"os"

	"github.com/MarcoVitangeli/go-type-generator/codegenerator"
	"github.com/MarcoVitangeli/go-type-generator/json"
)

func main() {
	d, err := json.ParseJson("test_files/types.json")

	if err != nil {
		fmt.Printf("error parsing json: %s\n", err)
		os.Exit(1)
	}

	err = os.Mkdir("typesgen", 0777)

	if err != nil {
		fmt.Printf("error creating directory: %s\n", err)
		os.Exit(1)
	}

	file, err := os.Create("typesgen/types.go")

	if err != nil {
		fmt.Printf("error creating file: %s\n", err)
		os.Exit(1)
	}

	err = codegenerator.GenerateCode(d, codegenerator.Writer{Writer: file})

	if err != nil {
		fmt.Printf("error generating code: %s\n", err)
		os.Exit(1)
	}
}
