package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("code").
			NotEmpty().
			Unique(),
	}
}

func (Subject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lesson_plans", LessonPlan.Type).
			Ref("subjects").
			Through("lesson_plan_subjects", LessonPlanSubject.Type),
	}
}

// Indexes of the Subject.
func (Subject) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("code"),
	}
}
