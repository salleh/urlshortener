package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Url holds the schema definition for the Url entity.
type Url struct {
	ent.Schema
}

func (Url) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Url.
func (Url) Fields() []ent.Field {
	return []ent.Field{
		field.String("short_path").Unique(),
		field.String("long_path"),
	}
}

// Edges of the Url.
func (Url) Edges() []ent.Edge {
	return nil
}
