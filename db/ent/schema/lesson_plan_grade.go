package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LessonPlanGrade struct {
	ent.Schema
}

func (LessonPlanGrade) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("lesson_plan_id"),
		field.Int64("grade_id"),
	}
}

func (LessonPlanGrade) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (LessonPlanGrade) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lesson_plan", LessonPlan.Type).
			Field("lesson_plan_id").
			Unique().
			Required(),
		edge.To("grade", Grade.Type).
			Field("grade_id").
			Unique().
			Required(),
	}
}

func (LessonPlanGrade) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_plan_id"),
		index.Fields("grade_id"),
	}
}
