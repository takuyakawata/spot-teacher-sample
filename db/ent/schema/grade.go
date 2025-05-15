package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Grade holds the schema definition for the Grade entity.
type Grade struct {
	ent.Schema
}

// Fields of the Grade.
func (Grade) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("code_number").Unique(), // 1
		field.String("code").
			NotEmpty().
			Unique(), // grade code
	}
}

func (Grade) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

// Edges of the Grade.
func (Grade) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lesson_plans", LessonPlan.Type).
			Ref("grades").
			Through("lesson_plan_grades", LessonPlanGrade.Type),
	}
}

// Indexes of the Grade.
func (Grade) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
		index.Fields("code_number"),
	}
}
