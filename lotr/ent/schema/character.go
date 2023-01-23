package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Character holds the schema definition for the character entity.
type Character struct {
    ent.Schema
}

// Fields of the Character struct.
func (User) Fields() []ent.Field {
    return []ent.Field{
		field.String("name"),
		field.String("culture")
		field.Int("magic")
	}
}

// Edges of the Character struct.
func (User) Edges() []ent.Edge {
    return nil
}