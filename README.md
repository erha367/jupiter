# jupyter 基于Gin实现的web框架
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
5. 异步队列（后续完善）
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
jupyter
├── Makefile
├── application				/*- 应用目录 -*/
│   ├── bootstrap.go
│   ├── controller			//控制器
│   ├── database			//数据库相关
│   ├── entity				//参数验证
│   ├── library				//通用类库（通常不依赖业务）
│   ├── model               //模型
│   ├── service             //服务层
│   ├── utils				//通用类库（依赖业务）
│   └── ytasks				//异步队列
├── commands				/*- 脚本相关 -*/
│   ├── listener.go         //监听定时任务
│   └── publish.go          //监听Redis
├── config                  /*- 配置相关 -*/
│   ├── app.go
│   ├── app_dev.json
│   ├── app_prod.json
│   └── app_test.json
├── go.mod
├── go.sum
├── log                      	/*- 日志 -*/
├── main.go    				 	/*- 入库文件（重要！） -*/	
├── readme.md
├── router                   	/*- 路由 -*/
│   ├── api.go
│   └── middleware           	/*- 中间件 -*/
├── sql							/*- 修改表结构sql -*/
│   └── init.sql
├── static						/*- 静态资源 -*/
│   └── x.html
├── test 						/*- 测试用例 -*/
│   ├── curl_test.go
│   └── model_test.go
└── vendor 						/*- 扩展 -*/
    ├── github.com
    ├── go.mongodb.org
    ├── go.uber.org
    ├── golang.org
    ├── gopkg.in
    └── modules.txt
```