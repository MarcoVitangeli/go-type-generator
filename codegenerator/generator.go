package codegenerator

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/MarcoVitangeli/go-type-generator/types"
)

func GenerateCode(t types.TypeData, w Writer) error {
	err := w.writeHeader()
	if err != nil {
		return err
	}

	err = w.writeInterfaces(t.Interfaces)

	if err != nil {
		return err
	}

	err = w.writeStructs(t.Structs)
	return err
}

func GenerateDirAndFile() (io.Writer, error) {
	err := os.Mkdir(DirName, 0777)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error creating directory: %s\n", err))
	}

	file, err := os.Create(FileName)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error creating file: %s\n", err))
	}

	return file, nil
}
