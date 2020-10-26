package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("password"),
		field.String("email"),
		field.String("token"),
		field.String("city"),
		field.String("role"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
