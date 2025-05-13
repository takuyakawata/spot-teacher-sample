package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Inquiry struct {
	ent.Schema
}

func (Inquiry) Fields() []ent.Field {
	return []ent.Field{
		field.Int("lesson_schedule_id").
			Positive(),

		field.Int("school_id").
			Positive(),

		field.Int("user_id").
			Positive(),

		field.Enum("category").
			Values("LESSON", "RESERVATION", "CANCELLATION", "OTHER").
			Default("OTHER"),

		field.Text("inquiry_detail"),
	}
}

func (Inquiry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lesson", LessonPlan.Type).
			Field("lesson_schedule_id").
			Unique().
			Required(),

		edge.To("school", School.Type).
			Field("school_id").
			Unique().
			Required(),

		edge.From("teacher", User.Type).
			Ref("inquiries").
			Field("user_id").
			Unique().
			Required(),
	}
}

func (Inquiry) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Inquiry) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_schedule_id"),
		index.Fields("school_id"),
		index.Fields("user_id"),
		index.Fields("category"),
		index.Fields("created_at"),
	}
}
