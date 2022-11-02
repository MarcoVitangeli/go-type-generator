package codegenerator

import (
	"fmt"
	"io"
	"strings"

	"github.com/MarcoVitangeli/go-type-generator/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Writer struct {
	io.Writer
}

func (w Writer) writeHeader() error {
	_, err := w.Write([]byte("package " + PackageName + "\n"))
	return err
}

func (w Writer) buildStruct(st types.Struct) string {
	name := strings.ToLower(st.Name)
	if st.Public {
		name = toTitle(name)
	}
	// they start with '\n' because of the space that the
	// 'package types' line will ocupate
	finalStr := fmt.Sprintf("\ntype %s struct {\n", name)

	for _, attr := range st.Attrs {
		name := strings.ToLower(attr.Name)
		if attr.Public {
			name = toTitle(name)
		}
		finalStr += fmt.Sprintf("\t%s %s\n", name, attr.Type)
	}
	finalStr += "}\n"

	return finalStr
}

func toTitle(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func (w Writer) writeStructs(t []types.Struct) error {
	var finalStr string
	for _, st := range t {
		finalStr += w.buildStruct(st)
	}
	_, err := w.Write([]byte(finalStr))

	return err
}

func (w Writer) buildInterface(it types.Interface) string {
	name := strings.ToLower(it.Name)
	if it.Public {
		name = toTitle(name)
	}

	finalStr := fmt.Sprintf("\ntype %s interface {\n", name)

	for _, attr := range it.Attrs {
		name := strings.ToLower(attr.Name)
		if it.Public {
			name = toTitle(name)
		}
		finalStr += fmt.Sprintf("\t%s(", name)
		var parArr []string
		for _, fp := range attr.Params {
			parArr = append(parArr, fmt.Sprintf("%s %s", strings.ToLower(fp.Name), fp.Type))
		}
		finalStr += strings.Join(parArr, ", ") + ")"

		rLen := len(attr.Returns)

		switch rLen {
		case 0:
			finalStr += "\n"
		case 1:
			finalStr += " " + attr.Returns[0] + "\n"
		default:
			finalStr += fmt.Sprintf(
				" (%s)\n",
				strings.Join(attr.Returns, ", "),
			)
		}
	}

	finalStr += "}\n"
	return finalStr
}

func (w Writer) writeInterfaces(its []types.Interface) error {
	var finalStr string
	for _, i := range its {
		finalStr += w.buildInterface(i)
	}
	_, err := w.Write([]byte(finalStr))

	return err
}
