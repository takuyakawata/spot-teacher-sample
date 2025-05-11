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
cd ../db
```
```bash
GOWORK=off atlas migrate diff create_base_tables_and_add_columns2 --dir file://ent/migrate/migrations --to ent://ent/schema --dev-url "docker://mysql/8/ent" # 開発DBのURLを正確に指定
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
