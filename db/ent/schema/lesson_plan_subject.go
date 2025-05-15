package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LessonPlanSubject struct {
	ent.Schema
}

func (LessonPlanSubject) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("lesson_plan_id"),
		field.Int64("subject_id"),
	}
}

func (LessonPlanSubject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (LessonPlanSubject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("lesson_plan", LessonPlan.Type).
			Field("lesson_plan_id").
			Unique().
			Required(),
		edge.To("subject", Subject.Type).
			Field("subject_id").
			Unique().
			Required(),
	}
}

func (LessonPlanSubject) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("lesson_plan_id"),
		index.Fields("subject_id"),
	}
}
