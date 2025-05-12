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
