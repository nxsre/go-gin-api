## 关于

`go-gin-api` 是基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，约束项目组开发成员，规避混乱无序及自由随意的编码。

供参考学习，线上使用请谨慎！

集成组件：

1. 支持 [rate](https://golang.org/x/time/rate) 接口限流 
1. 支持 panic 异常时邮件通知 
1. 支持 [cors](https://github.com/rs/cors) 接口跨域 
1. 支持 [Prometheus](https://github.com/prometheus/client_golang) 指标记录 
1. 支持 [Swagger](https://github.com/swaggo/gin-swagger) 接口文档生成 
1. 支持 [GraphQL](https://github.com/99designs/gqlgen) 查询语言 
1. 支持 trace 项目内部链路追踪 
1. 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析
1. 支持 errno 统一定义错误码 
1. 支持 [zap](https://go.uber.org/zap) 日志收集 
1. 支持 [viper](https://github.com/spf13/viper) 配置文件解析
1. 支持 [gorm](https://gorm.io/gorm) 数据库组件
1. 支持 [go-redis](https://github.com/go-redis/redis/v7) 组件
1. 支持 RESTful API 返回值规范
1. 支持 生成数据表 CURD、控制器方法 等代码生成器
1. 支持 [cron](https://github.com/jakecoffman/cron) 定时任务，在后台可界面配置
1. 支持 [websocket](https://github.com/gorilla/websocket) 实时通讯，在后台有界面演示
1. 支持 web 界面，使用的 [Light Year Admin 模板](https://gitee.com/yinqi/Light-Year-Admin-Using-Iframe)


## 文档索引（可加入交流群）

- 中文文档：[go-gin-api - 语雀](https://www.yuque.com/xinliangnote/go-gin-api/ngc3x5)
- English Document：[en.md](https://github.com/nxsre/go-gin-api/blob/master/en.md)



## 快速启动

### MySQL 初始化
```mysql
create database go_gin_api;

create user 'root'@'127.0.0.1' identified by '123456';

grant create,select,insert,update,delete on go_gin_api.* to 'root'@'127.0.0.1';
```

### 生成 graphql 代码
```shell
printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
go run github.com/99designs/gqlgen init
go mod tidy
go run github.com/99designs/gqlgen
```
#### 错误处理
```text
generating core failed: unable to load github.com/nxsre/go-gin-api/internal/graph/model - make sure you're using an import path to a package that exists
```
#### 修改 gqlgen.yml 重新运行 gqlgen
```shell
vi gqlgen.yml
```
```text
...
# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.
#autobind:
#   - "github.com/nxsre/go-gin-api/internal/graph/model"
...
```
```go
go run github.com/99designs/gqlgen
```

## 目录结构
```shell
.
├── assets                ## 静态资源
│   ├── bootstrap
│   └── templates
├── cmd                   ## 可独立执行程序
│   ├── ...
├── configs               ## 配置文件
├── deploy                ## 部署相关配置
│   ├── loki            
│   └── prometheus
├── docs
├── internal              ## 业务目录
│   ├── api               ## api 接口
│   ├── code              ## 错误码定义
│   ├── graph             ## graph 接口
│   ├── pkg               ## 内部使用的 package
│   ├── proposal          ## 提案
│   ├── render            ## 渲染 HTML
│   ├── repository        ## 资源
│   ├── router            ## 路由
│   ├── services          ## 服务
│   └── websocket         ## websocket 接口
├── logs                  ## 日志目录
├── pkg                   ## 可供外部使用的 package
│   ├── ...
└── scripts               ## 脚本程序

```
## 其他

查看 Jaeger 链路追踪 Demo 代码，请查看 [v1.0 版](https://github.com/nxsre/go-gin-api/releases/tag/v1.0) ，链接地址：http://127.0.0.1:9999/jaeger_test

调用的其他服务端 Demo 代码为 [https://github.com/nxsre/go-jaeger-demo](https://github.com/nxsre/go-jaeger-demo)

## 联系作者

![联系作者](https://i.loli.net/2021/07/02/cwiLQ13CRgJIS86.jpg)

