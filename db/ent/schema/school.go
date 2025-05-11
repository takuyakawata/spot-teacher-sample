package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type School struct{ ent.Schema }

func (School) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),

		field.Enum("school_type").
			Values("elementary", "juniorHigh", "highSchool"),

		field.String("name").
			NotEmpty().
			MaxLen(50),

		field.String("email").
			Optional().
			MaxLen(200),

		field.String("phone_number").
			NotEmpty(),

		field.Int("prefecture").
			Min(0).
			Max(46),

		field.String("city").
			NotEmpty(),

		field.String("street").
			Optional(),

		field.String("post_code").
			NotEmpty().
			MaxLen(7),

		field.String("url").
			Optional(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (School) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teachers", User.Type),
	}
}

func (School) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("school_type"),
	}
}
