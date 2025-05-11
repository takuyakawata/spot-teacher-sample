package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type Company struct{ ent.Schema }

func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),

		field.String("name").
			NotEmpty().
			MaxLen(50),

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

		field.String("phone_number").
			NotEmpty(),

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

func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lesson_plans", LessonPlan.Type),
		edge.To("members", User.Type),
	}
}

func (Company) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
