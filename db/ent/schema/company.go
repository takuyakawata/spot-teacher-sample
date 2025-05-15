package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Company struct{ ent.Schema }

func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(50),

		field.Int("prefecture").
			Min(0).
			Max(50),

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
	}
}

func (Company) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
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
