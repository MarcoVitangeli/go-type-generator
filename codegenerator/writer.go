package codegenerator

import (
	"fmt"
	"io"
	"strings"

	"github.com/MarcoVitangeli/go-type-generator/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type writer struct {
	io.Writer
}

func (w writer) buildStruct(st types.Struct) string {
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

func (w writer) writeStructs(t []types.Struct) {
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

	w.Write([]byte(finalStr))
}
