
# API開發管理系統


### 文档结构（1.0）

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

### 文檔結構（test）
~~~powershell
├── api                     # api接口（暫定）
├── cmd
│   └── server
│   |   └── app.go          # 可執行文件生成（暫定）
│
├── config
│   └── settings.yaml       # 配置文件(settings.yaml)
│
├── core                    # 初始化 x
├── flags                   # 命令参数(cmd指令) ？
│
├── global
│   └── global.go           # 全局变量 （暫定）
│
├── internal
│   ├── conf                # 配置（入參）
│   ├── middleware          # 中间件
│   ├── models              # orm数据表
│   ├── router              # 路由
│   └── service             # 服務層
│
├── test                    # 測試
├── utils                   # 工具包
└── main.go                 # 項目主入口（暫定）
~~~