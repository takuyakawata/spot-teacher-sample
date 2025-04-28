package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type User struct{ ent.Schema }

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").NotEmpty().MaxLen(100),

		field.String("family_name").NotEmpty().MaxLen(100),

		field.String("email").NotEmpty().MaxLen(100).Unique(),

		field.String("password").NotEmpty().Sensitive(),

		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		//name
		index.Fields("first_name", "family_name"),
	}
}
