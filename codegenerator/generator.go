package codegenerator

import (
	"github.com/MarcoVitangeli/go-type-generator/types"
)

func GenerateCode(t types.TypeData, w writer) {
	writeHeader(w)
	w.writeStructs(t.Structs)
}

func writeHeader(w writer) {
	w.Write([]byte("package types"))
}
