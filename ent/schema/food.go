package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Food holds the schema definition for the Food entity.
type Food struct {
	ent.Schema
}

// Fields of the Food.
func (Food) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("restaurant"),
		field.String("desc"),
		field.String("price"),
		field.String("image_name"),

	}
}

// Edges of the Food.
func (Food) Edges() []ent.Edge {
	return nil
}
