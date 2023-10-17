package ast

import (
	"fmt"
	"strings"
)

type Scope struct {
	Outer   *Scope
	Objects map[string]*Object
}

func NewScope(outer *Scope) *Scope {
	return &Scope{
		Outer:   outer,
		Objects: make(map[string]*Object),
	}
}

func (s *Scope) Lookup(name string) *Object {
	return s.Objects[name]
}

func (s *Scope) Insert(obj *Object) {
	s.Objects[obj.Name] = obj
}

// Debugging support
func (s *Scope) String() string {
	var buf strings.Builder
	buf.WriteString("scope %p {")
	buf.WriteString(s.String())
	if s != nil && len(s.Objects) > 0 {
		buf.WriteString("\n")
		for _, obj := range s.Objects {
			buf.WriteString(fmt.Sprintf("\t%s %s\n", obj.Kind, obj.Name))
		}
	}
	buf.WriteString("}\n")
	return buf.String()
}

type Object struct {
	Kind ObjectKind  // Тип объекта
	Name string      // Declared name
	Data interface{} // даннные специфичные для объекта
}

func NewObject(kind ObjectKind, name string) *Object {
	return &Object{
		Kind: kind,
		Name: name,
		Data: nil,
	}
}

type ObjectKind int

const (
	OBJECT_KIND_BAD ObjectKind = iota // для обработки ошибок
	OBJECT_KIND_PACKAGE
	OBJECT_KIND_CONST
	OBJECT_KIND_VARIABLE
	OBJECT_KIND_FUNCTION
)
