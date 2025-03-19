package field_type

import "path"

type Custom struct {
	Name string

	// PkgName if empty — not from another pkg
	PkgName string
	// PkgPath if empty — not from another pkg
	PkgPath string
}

func (c Custom) Kind() Kind {
	return KindCustom
}

func (c Custom) Unwraps() []FieldTyper {
	return nil
}

func (c Custom) GoTypeString() string {
	typeName := c.Name
	if c.PkgName != "" {
		typeName = c.PkgName + "." + typeName
	}

	if c.PkgName == "" && c.PkgPath != "" {
		typeName = path.Base(c.PkgPath) + "." + typeName
	}

	return typeName
}

func (c Custom) GoTypeStringWithAlias(alias string) string {
	typeName := c.Name
	if c.PkgName != "" {
		typeName = c.PkgName + "." + typeName
	}

	if c.PkgName == "" && c.PkgPath != "" {
		typeName = path.Base(c.PkgPath) + "." + typeName
	}

	if c.PkgName == "" && c.PkgPath == "" {
		typeName = alias + "." + typeName
	}

	return typeName
}

func (c Custom) String() string {
	return c.GoTypeString()
}

func (c Custom) Accept(visitor Visitor) {
	visitor.VisitCustom(c)
}

func CustomFiled(name, pkgName, pkgPath string) Custom {
	return Custom{
		Name:    name,
		PkgName: pkgName,
		PkgPath: pkgPath,
	}
}
