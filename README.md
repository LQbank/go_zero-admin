# go_zero-admin

安装go和goctl工具和框架
go mod tidy
etcd // 启动etcd
1.使用Gen工具创建Qury操作类和数据库结构
进到models目录操作

go run main.go
2.创建api
进到api/doc/目录执行

goctl api -o admin.api
goctl api go -api admin.api -dir ../ --home ../../.template
3.创建rpc
进到rpc/sys/目录操作

goctl rpc template -o sys.proto
goctl rpc protoc sys.proto --go_out=./ --go-grpc_out=./ --zrpc_out=. --home ../../.template
4.运行测试
进到rpc/models目录导入数据库

go run main.go
进到api/目录操作

go run admin.go -f etc/admin-api.yaml
进到rpc/sys/目录操作

go run sys.go -f etc/sys.yaml
5.测试运行
$ curl -i -X POST \
  http://127.0.0.1:8888/api/sys/user/login \
  -H 'content-type: application/json' \
  -d '{"userName":"admin", "password":"123456"}'
