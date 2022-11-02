package main

import (
	"fmt"
	"os"

	"github.com/MarcoVitangeli/go-type-generator/json"
)

func main() {
	d, err := json.ParseJson("test_files/types.json")

	if err != nil {
		fmt.Printf("error parsing json: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(d.Interfaces)
}
