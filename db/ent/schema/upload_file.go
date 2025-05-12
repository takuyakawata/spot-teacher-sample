package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type UploadFile struct{ ent.Schema }

func (UploadFile) Fields() []ent.Field {
	return []ent.Field{
		field.String("photo_key").NotEmpty().Comment("ファイルのユニークキー"),
		field.Int64("user_id").Positive().Comment("写真をアップロードしたユーザーのID"),
	}
}

func (UploadFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("photo_key"),
	}
}

func (UploadFile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
