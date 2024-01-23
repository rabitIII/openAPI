
### 文档结构

~~~powershell
├── api                     # api接口
├── config                  # 配置(settings.yaml)
├── core                    # 初始化
├── flags                   # 命令参数(cmd指令)
├── global
│   └── global.go           # 全局变量 
├── middleware              # 中间件
├── models                  # orm数据表
├── router                  # 路由
├── service                 # 服务
├── test                    # 测试
├── utils                   # 工具包
├── settings.yaml           # 配置文件
~~~

### 项目启动顺序
1. 环境配置：`settings.yaml`
2. 初始化数据库，在根目录下执行以下命令：
~~~powershell
go run main.go -db
~~~
