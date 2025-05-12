package schema

import (
	"entgo.io/ent/schema/index"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Inquiry struct {
	ent.Schema
}

func (Inquiry) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Positive(),

		field.Int64("lesson_schedule_id").
			Positive(),

		field.Int64("school_id").
			Positive(),

		field.Int64("user_id").
			Positive(),

		field.Enum("category").
			Values("LESSON", "RESERVATION", "CANCELLATION", "OTHER").
			Default("OTHER"),

		field.Text("inquiry_detail"),

		field.Time("deleted_at").
			Optional().
			Nillable(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
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

		edge.To("teacher", User.Type).
			Field("user_id").
			Unique().
			Required(),
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
