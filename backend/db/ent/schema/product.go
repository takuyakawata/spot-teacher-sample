package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type Product struct{ ent.Schema }

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(100),
		field.Int("price").
			Positive(),
		field.String("description").
			Optional().
			MaxLen(500),
		field.Time("created_at"). // カラム名: created_at, 型: TIMESTAMP NOT NULL
						Default(time.Now). // Go側でのデフォルト値（レコード作成時に現在時刻が入る）
						Immutable(),       // 作成後に変更できないようにする

		field.Time("updated_at"). // カラム名: updated_at, 型: TIMESTAMP NOT NULL
						Default(time.Now). // Go側でのデフォルト値
						UpdateDefault(time.Now),
	}
}
