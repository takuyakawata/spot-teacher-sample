package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LessonPlanUploadFile struct {
	ent.Schema
}

func (LessonPlanUploadFile) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("lesson_plan_id"),
		field.Int64("upload_file_id"),
	}
}

func (LessonPlanUploadFile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (LessonPlanUploadFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lesson_plan", LessonPlan.Type).
			Field("lesson_plan_id").
			Unique().
			Required(),
		edge.To("upload_file", UploadFile.Type).
			Field("upload_file_id").
			Unique().
			Required(),
	}
}

func (LessonPlanUploadFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_plan_id"),
		index.Fields("upload_file_id"),
	}
}
