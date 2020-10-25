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
		field.String("role"),
		field.String("city"),
		field.String("token"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
