package validation

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	"github.com/MarcoVitangeli/go-type-generator/codegenerator"
)

func Validate() error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, codegenerator.FileName, nil, 0)
	if err != nil {
		return err
	}

	// Run type checker
	info := types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}

	_, err = (&types.Config{}).Check(codegenerator.PackageName, fset, []*ast.File{f}, &info)
	if err != nil {
		errStr := fmt.Sprintf("Error: check your type definitions in your json file, some checks failed")
		return errors.New(errStr + "\n" + err.Error())
	}
	return nil
}
