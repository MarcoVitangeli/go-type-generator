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
	_, err := w.Write([]byte("package " + DirName + "\n"))
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
		name := strings.ToLower(st.Name)
		if st.Public {
			name = toTitle(name)
		}
		// they start with '\n' because of the space that the
		// 'package types' line will ocupate
		finalStr += fmt.Sprintf("\ntype %s struct {\n", name)

		for _, attr := range st.Attrs {
			name := strings.ToLower(attr.Name)
			if attr.Public {
				name = toTitle(name)
			}
			finalStr += fmt.Sprintf("\t%s %s\n", name, attr.Type)
		}
		finalStr += "}\n"
	}

	_, err := w.Write([]byte(finalStr))

	return err
}
