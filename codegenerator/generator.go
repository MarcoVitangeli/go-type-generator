package codegenerator

import (
	"github.com/MarcoVitangeli/go-type-generator/types"
)

func GenerateCode(t types.TypeData, w Writer) error {
	err := w.writeHeader()
	if err != nil {
		return err
	}
	err = w.writeStructs(t.Structs)
	return err
}
