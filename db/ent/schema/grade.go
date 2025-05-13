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
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("code").
			NotEmpty().
			Unique(),
	}
}

func (Grade) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Grade.
func (Grade) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("lesson_plans", LessonPlan.Type).
			Ref("grades"),
	}
}

// Indexes of the Grade.
func (Grade) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("code"),
	}
}
