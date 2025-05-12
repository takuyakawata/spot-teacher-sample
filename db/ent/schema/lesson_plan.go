package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type LessonPlan struct{ ent.Schema }

func (LessonPlan) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive(),

		field.Int64("company_id").Positive(),

		field.String("title").NotEmpty().MaxLen(500),

		field.String("description").Optional().NotEmpty().MaxLen(2000),

		field.String("location").Optional().NotEmpty().MaxLen(500),

		field.Enum("lesson_type").Values("online", "offline", "online_and_offline"),

		field.Int("annual_max_executions").Positive().Comment("年間可能実施回数"),

		field.Int("start_month").Min(1).Max(12),

		field.Int("start_day").Min(1).Max(31),

		field.Int("end_month").Min(1).Max(12),

		field.Int("end_day").Min(1).Max(31),

		field.Time("start_time"),

		field.Time("end_time"),

		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (LessonPlan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("schedules", LessonSchedule.Type),
		edge.From("company", Company.Type).
			Ref("lesson_plans").
			Field("company_id").
			Unique().
			Required(),
		edge.To("grades", Grade.Type),
		edge.To("subjects", Subject.Type),
		edge.To("education_categories", EducationCategory.Type),

		//edge.From("uploadFile", UploadFile.Type).
		//	Ref("lessonPlans").
		//	Unique(),
	}
}

func (LessonPlan) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("company_id"),
		index.Fields("title"),
	}
}
