package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LessonConfirmation struct{ ent.Schema }

func (LessonConfirmation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("lesson_reservation_id").Positive(),

		field.Time("matching_date"),

		field.Time("start_time").
			Comment("確定した授業の開始時間"),

		field.Time("finish_time").
			Comment("確定した授業の終了時間"),

		field.String("remarks").
			Optional().
			Comment("備考"),
	}
}

func (LessonConfirmation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lesson_reservation", LessonReservation.Type).
			Ref("lesson_confirmation").
			Field("lesson_reservation_id").
			Unique().
			Required(),
	}
}

func (LessonConfirmation) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
