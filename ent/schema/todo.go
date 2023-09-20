package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

func (Todo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name"),
		field.Bool("completed").
			Default(false).
			Optional().
			Nillable(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
