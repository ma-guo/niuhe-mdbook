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
关于本插件的更多丰富功能, 请移步[官方教程](http://niuhe.zuxing.net)

