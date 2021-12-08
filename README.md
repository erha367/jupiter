# jupiter 基于Gin实现的web框架【采用领域驱动设计】
作者：杨森 sen.yang@eeoa.com
#### 注意！要根据实际情况修改配置！！！
1. 修改mysql配置
2. 修改redis配置
## golang领域驱动介绍
https://qw41hb3siw.feishu.cn/docs/doccnBS4JYVal8WYMD1Sfj96SDb
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
# env的枚举值为 `dev`, `13`, `14`, `im`, `prod` 等
# 平滑重启
kill -1 pid   #main进程id
```
## 其他
### 代码结构
```
jupiter/
├── application    //应用层
├── config         //配置文件
├── domain        //领域层
│   ├── entity    //实体
│   ├── repository    //仓库
│   ├── service    //服务
│   └── valobj    //值对象（约定常量）
├── interfaces    //接口层
│   ├── commands    //脚本入口
│   ├── controller    //控制器
│   ├── rpc            //rpc接口
│   └── validator    //验证器
├── library    //基础层
├── main.go    //主程序入口
└── router    //路由
    └── middleware    //中间件  
```