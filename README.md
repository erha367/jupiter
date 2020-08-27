# jupyter 基于Gin实现的web框架
#### 注意！要根据实际情况修改配置！！！
## 上线须知（重要）
1. 所有上线分支均为master
2. 非上线代码禁止合并到master
3. 上线时请备份好旧版本的 main 运行文件，作为版本回滚用
4. 建议放在supervisor下管理，可及时重启服务
## 功能
1. mysql\redis 支持
2. 内存缓存支持
3. 日志
4. debug追踪
5. 异步队列
6. 热重启
7. https|http 双端口支持
## 运行环境依赖
- Unix
- mysql
- redis
## 编译 & 运行
```shell script
# 版本要求 golang 14.0 +
# 1、进入到项目目录
go mod tidy
# 2、编译项目（linux平台）
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# 编译后的二进制文件（main）连同config目录一起放到线上项目目录即可（确保项目目录权限755）
# 运行
# 测试环境
./main -env=test
# 线上环境
./main -env=prod
# 平滑重启
kill -1 pid   #main进程id
```
## 其他
### 代码结构
```
jupyter/
├── application
│   ├── controller
│   ├── database
│   ├── entity
│   │   └── form
│   ├── library
│   ├── model
│   ├── service
│   └── utils
├── config
├── router
│   └── middleware
├── sql
├── static
└── test
```