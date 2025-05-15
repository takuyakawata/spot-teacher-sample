package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Product struct{ ent.Schema }

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(100),
		field.Int("price").
			Positive(),
		field.String("description").
			Optional().
			MaxLen(500),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (Product) Edges() []ent.Edge {
	return nil
}
