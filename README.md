# mod
```sh
    go env -w GO111MODULE=auto
    cd src/demo && go mod init demo && go mod tidy && go mod vendor && cd ../../ && make run
```

db 配置格式
```yaml
db:
	main:user:pwd@tcp(host:port)/database_name?charset=utf8mb4
```

# 教程
关于本插件的更多丰富功能, 请移步[掘金专栏 -niuhe 插件](https://juejin.cn/column/7376620206338785314)。 在这里收获更多体验

