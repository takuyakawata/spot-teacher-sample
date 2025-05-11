package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type LessonSchedule struct{ ent.Schema }

func (LessonSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive(),

		field.Int64("lesson_plan_id").Positive(),

		field.String("title").NotEmpty().MaxLen(100),

		field.String("description").Optional().MaxLen(500),

		field.String("location").Optional().MaxLen(500),

		field.Enum("lesson_type").Values("online", "offline", "online_and_offline"),

		field.Int("annual_max_executions").Positive().Comment("年間可能実施回数"),

		field.Time("start_date"), // 2024.04.04 00 00 00

		field.Time("end_date"), // 2024.09.09 00 00 00

		field.Time("start_time"), // 12:00

		field.Time("end_time"), // 16:00

		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (LessonSchedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plan", LessonPlan.Type).
			Ref("schedules").
			Field("lesson_plan_id").
			Unique().
			Required(),
	}
}

func (LessonSchedule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_plan_id"),
		index.Fields("title"),
	}
}
