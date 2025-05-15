package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type User struct{ ent.Schema }

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("user_type").
			Values("teacher", "company_member", "admin").
			Default("teacher"),

		field.Int64("school_id").
			Positive().
			Optional().
			Nillable(),

		field.Int64("company_id").
			Positive().
			Optional().
			Nillable(),

		field.String("first_name").NotEmpty().MaxLen(50),

		field.String("family_name").NotEmpty().MaxLen(50),

		field.String("email").NotEmpty().MaxLen(100).Unique(),

		field.String("phone_number").MaxLen(20),

		field.String("password").
			Optional().
			Nillable().
			Sensitive(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("school", School.Type).
			Ref("teachers").
			Field("school_id").
			Unique(),

		edge.From("company", Company.Type).
			Ref("members").
			Field("company_id").
			Unique(),

		edge.To("inquiries", Inquiry.Type),

		edge.To("lesson_reservations", LessonReservation.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		//name
		index.Fields("first_name", "family_name"),
		index.Fields("user_type"),
		index.Fields("school_id"),
		index.Fields("company_id"),
		index.Fields("email"),
	}
}
