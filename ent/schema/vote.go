package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.String("status"),
		field.String("restaurant"),
		field.String("food"),
		field.String("username"),
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return nil
}
