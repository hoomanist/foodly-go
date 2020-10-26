package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field {
		field.String("food"),
		field.String("restaurant"),
		field.String("msg"),
		field.String("username"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
