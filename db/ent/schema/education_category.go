package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// EducationCategory holds the schema definition for the EducationCategory entity.
type EducationCategory struct {
	ent.Schema
}

// Fields of the EducationCategory.
func (EducationCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("code").
			NotEmpty().
			Unique(),
	}
}

func (EducationCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

// Edges of the EducationCategory.
func (EducationCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lesson_plans", LessonPlan.Type).
			Ref("education_categories").
			Through("lesson_plan_education_categories", LessonPlanEducationCategory.Type),
	}
}

// Indexes of the EducationCategory.
func (EducationCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("code"),
	}
}
