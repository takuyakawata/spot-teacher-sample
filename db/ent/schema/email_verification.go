package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type EmailVerification struct {
	ent.Schema
}

func (EmailVerification) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("token").Unique().Sensitive(),
		field.Time("expired_at"),
	}
}

func (EmailVerification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("expired_at"),
		index.Fields("token"),
		index.Fields("email"),
	}
}

func (EmailVerification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}
