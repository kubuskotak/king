package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
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

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
