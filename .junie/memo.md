開発中のメモ　

gitにいれてはいけないデータに注意すること


- moduleを削除するコマンド
root dirで実行
```bash
go clean -modcache
```

- workspace modeのmoduleの最適化
```bash
go work sync
```

```bash

```
```bash
cd ./db & GOWORK=off atlas migrate diff create_inqueries_tables --dir file://ent/migrate/migrations --to ent://ent/schema --dev-url "docker://mysql/8/ent"
```

test api
- http://localhost:3000/api/spot-teacher/product/test


````
{
"build": {
"env": {
"GO_BUILD_FLAGS": "-ldflags '-s -w'"
}
},
"functions": {
"spot-teacher/api/*.go": {
"runtime": "go@1.23.2"
}
}
}
````
テーブル設計について悩んでいます

lessonPlanとlessonScheduleで
1:Nでgradeとsubjectというカテゴリーを持ちます
gradeは
subjectともに変更頻度は少ない

案1
lessonPlan_id
code (enum)
でサブテーブルで正規化

テーブル数を減らせる
gradeなどカテゴリの管理は、事務局ユーザーでできない


案2
lessonPlan_id
grade_id
のように中間テーブルで正規化

テーブルが増える
gradeなどカテゴリーの編集管理ができる


検索時にカテゴリーを条件にしてlessonplanを検索します

以上を踏まえて、意見をください
