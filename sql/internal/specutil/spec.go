package specutil

import (
	"ariga.io/atlas/schema/schemaspec"
	"ariga.io/atlas/sql/sqlspec"
)

// NewCol is a helper method for constructing *sqlspec.Column instances.
func NewCol(name, coltype string, attrs ...*schemaspec.Attr) *sqlspec.Column {
	return &sqlspec.Column{
		Name:     name,
		TypeName: coltype,
		DefaultExtension: schemaspec.DefaultExtension{
			Extra: schemaspec.Resource{Attrs: attrs},
		},
	}
}

// LitAttr is a helper method for constructing *schemaspec.Attr instances that contain literal values.
func LitAttr(k, v string) *schemaspec.Attr {
	return &schemaspec.Attr{
		K: k,
		V: &schemaspec.LiteralValue{V: v},
	}
}

// ListAttr is a helper method for constructing *schemaspec.Attr instances that contain list values.
func ListAttr(k string, litValues ...string) *schemaspec.Attr {
	lv := &schemaspec.ListValue{}
	for _, v := range litValues {
		lv.V = append(lv.V, &schemaspec.LiteralValue{V: v})
	}
	return &schemaspec.Attr{
		K: k,
		V: lv,
	}
}