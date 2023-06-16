package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Hello holds the schema definition for the Hello entity.
type Hello struct {
	ent.Schema
}

// Fields of the Hello.
func (Hello) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			Default(""),
		field.String("body").
			NotEmpty().
			Default(""),
		field.String("description").
			NotEmpty().
			Default(""),
		field.String("slug").
			NotEmpty().
			Unique(),
		field.Int("user_id").
			Optional(),
	}
}

// Edges of the Hello.
func (Hello) Edges() []ent.Edge {
	return nil
}
