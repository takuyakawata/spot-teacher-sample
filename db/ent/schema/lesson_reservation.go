package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type LessonReservation struct {
	ent.Schema
}

func (LessonReservation) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Positive(),

		field.Int64("lesson_schedule_id").
			Positive(),

		field.Int64("school_id").
			Positive(),

		field.Int64("user_id").
			Positive(),

		field.Enum("reservation_status").
			Values("PENDING", "APPROVED", "CANCELED"),

		field.String("count_student"),

		field.String("graduate"),

		field.String("subject"),

		field.String("remarks").
			Optional().
			Nillable(),

		field.Time("reservation_confirm_at").
			Optional().
			Nillable(),
	}
}

func (LessonReservation) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (LessonReservation) Edge() []ent.Edge {
	return []ent.Edge{
		edge.To("lesson_schedule", LessonSchedule.Type).
			Field("lesson_schedule_id").
			Unique().
			Required(),

		edge.To("school", School.Type).
			Field("school_id").
			Unique().
			Required(),

		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required(),

		edge.To("preferred_dates", LessonReservationPreferredDate.Type),

		//edge.To("confirmation", LessonConfirmation.Type),
	}
}

func (LessonReservation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_schedule_id"),
		index.Fields("school_id"),
		index.Fields("user_id"),
		index.Fields("reservation_status"),
		index.Fields("created_at"),
	}
}
