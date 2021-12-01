# jupiter 基于Gin实现的web框架
作者：杨森 sen.yang@eeoa.com
#### 注意！要根据实际情况修改配置！！！
1. 修改mysql配置
2. 修改redis配置
## 上线须知（重要）
1. 所有上线分支均为master
2. 非上线代码禁止合并到master
3. 上线时请备份好旧版本的 main 运行文件，作为版本回滚用
4. 建议放在supervisor下管理，可及时重启服务
## 功能
1. mysql\redis 集群支持
2. 内存缓存支持
3. zap日志
4. debug追踪
5. ~~异步队列~~
6. ~~endless热重启~~
7. ~~https|http 双端口支持~~
8. grpc+consul 支持
## 运行环境依赖
- Centos
- mysql
- redis集群模式
## 编译 & 运行
```shell script
# 版本要求 golang 14.0 +
# go配置了代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
# 1、进入到项目目录
go mod tidy
# 2、编译项目（mac开发环境-->linux测试-->生产环境）
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
# 编译后的二进制文件（main）连同config目录一起放到线上项目目录即可（确保项目目录权限755）
# 运行
# 开发环境
./main -env=dev
# 线上环境
./main -env=prod
# 平滑重启
kill -1 pid   #main进程id
```
## 其他
### 代码结构
```
jupiter
├── application //业务相关
│   ├── controller  //控制器
│   ├── domain      //领域
│   ├── entity      //实体
│   ├── model       //模型
│   ├── service     //服务
│   └── validator   //验证器
├── config  //配置
├── library //公共组件
├── logs    //日志（可以定义位置）
└── router  //路由表
    └── middleware  //中间件   
```