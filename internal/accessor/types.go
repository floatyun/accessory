package accessor

import (
	"go/types"
	"unicode"

	"golang.org/x/tools/go/packages"
)

// ParsedSource contains the parsed source of a package.
type ParsedSource struct {
	Package *packages.Package
	Dir     string
	Imports []*Import
	Structs []*Struct
}

// Import contains the information of an import.
type Import struct {
	Name    string
	Path    string
	IsNamed bool
}

// Struct contains the information of a struct.
type Struct struct {
	Name   string
	Fields []*Field
}

// Field contains the information of a field in a struct.
type Field struct {
	Name string
	Type types.Type
	Tag  *Tag
}

func (f *Field) IsExported() bool {
	if f == nil {
		return false
	}
	for _, r := range f.Name {
		return unicode.IsUpper(r)
	}
	return false
}

// Tag contains the information of a struct field's tag.
type Tag struct {
	Getter    *string
	Setter    *string
	NoDefault bool
}
