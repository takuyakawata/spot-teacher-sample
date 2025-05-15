package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LessonPlan struct{ ent.Schema }

func (LessonPlan) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("company_id").Positive(),

		field.String("title").NotEmpty().MaxLen(500),

		field.String("description").Optional().NotEmpty().MaxLen(2000),

		field.String("location").Optional().NotEmpty().MaxLen(500),

		field.Enum("lesson_type").Values("online", "offline", "online_and_offline"),

		field.Int64("annual_max_executions").Positive().Comment("年間可能実施回数"),

		field.Int64("start_month").Min(1).Max(12),

		field.Int64("start_day").Min(1).Max(31),

		field.Int64("end_month").Min(1).Max(12),

		field.Int64("end_day").Min(1).Max(31),

		field.Time("start_time"),

		field.Time("end_time"),
	}
}

func (LessonPlan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (LessonPlan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("lesson_plans").
			Field("company_id").
			Unique().
			Required(),
		edge.To("schedules", LessonSchedule.Type),

		edge.To("upload_files", UploadFile.Type).
			Through("lesson_plan_upload_files", LessonPlanUploadFile.Type),

		edge.To("subjects", Subject.Type).
			Through("lesson_plan_subjects", LessonPlanSubject.Type),

		edge.To("grades", Grade.Type).
			Through("lesson_plan_grades", LessonPlanGrade.Type),

		edge.To("education_categories", EducationCategory.Type).
			Through("lesson_plan_education_categories", LessonPlanEducationCategory.Type),
	}
}

func (LessonPlan) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("company_id"),
		index.Fields("title"),
	}
}
