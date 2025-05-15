package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LessonReservation struct {
	ent.Schema
}

func (LessonReservation) Fields() []ent.Field {
	return []ent.Field{
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

func (LessonReservation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lesson_schedule", LessonSchedule.Type).
			Ref("lesson_reservations").
			Field("lesson_schedule_id").
			Unique().
			Required(),

		edge.From("school", School.Type).
			Ref("lesson_reservations").
			Field("school_id").
			Unique().
			Required(),

		edge.From("user", User.Type).
			Ref("lesson_reservations").
			Field("user_id").
			Unique().
			Required(),

		edge.To("lesson_reservation_preferred_dates", LessonReservationPreferredDate.Type),
		edge.To("lesson_confirmation", LessonConfirmation.Type),
	}
}

func (LessonReservation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (LessonReservation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_schedule_id"),
		index.Fields("school_id"),
		index.Fields("user_id"),
		index.Fields("reservation_status"),
	}
}
