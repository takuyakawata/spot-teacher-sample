package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type LessonReservationPreferredDate struct {
	ent.Schema
}

func (LessonReservationPreferredDate) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("priority").
			Values("FIRST", "SECOND", "THIRD", "FOURTH", "FIFTH").
			Comment("希望日時の優先順位"),

		field.Time("date").
			Comment("希望日"),

		// TODO timeのみで表す方法がわかり次第変更
		field.Time("start_time").
			Comment("希望開始時間"),

		field.Time("end_time").
			Comment("希望終了時間"),
	}
}

func (LessonReservationPreferredDate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (LessonReservationPreferredDate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lessonReservations", LessonReservation.Type).
			Ref("preferredDates").
			Unique().
			Required(),
	}
}

func (LessonReservationPreferredDate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_reservation_id"),
		index.Fields("priority"),
		index.Fields("date", "start_time"),
	}
}
