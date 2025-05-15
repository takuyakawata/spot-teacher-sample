package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UploadFile struct{ ent.Schema }

func (UploadFile) Fields() []ent.Field {
	return []ent.Field{
		field.String("photo_key").NotEmpty().Comment("ファイルのユニークキー"),
		field.Int64("user_id").Positive().Comment("写真をアップロードしたユーザーのID"),
	}
}

func (UploadFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("LessonPlan", LessonPlan.Type).
			Ref("upload_files").
			Through("lesson_plan_upload_files", LessonPlanUploadFile.Type),
	}
}

func (UploadFile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
		TimeMixin{},
	}
}

func (UploadFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("photo_key"),
	}
}
