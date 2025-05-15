package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Mixin struct {
	mixin.Schema
}

func (Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
	}
}
