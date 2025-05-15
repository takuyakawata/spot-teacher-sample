package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LessonPlanEducationCategory struct {
	ent.Schema
}

func (LessonPlanEducationCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("lesson_plan_id"),
		field.Int64("education_category_id"),
	}
}

func (LessonPlanEducationCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (LessonPlanEducationCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lesson_plan", LessonPlan.Type).
			Field("lesson_plan_id").
			Unique().
			Required(),

		edge.To("education_category", EducationCategory.Type).
			Field("education_category_id").
			Unique().
			Required(),
	}
}

func (LessonPlanEducationCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_plan_id"),
		index.Fields("education_category_id"),
	}
}
