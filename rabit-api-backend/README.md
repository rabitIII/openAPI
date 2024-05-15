~~~powershell
go-project/
│
├── cmd/                                # 应用程序的入口点
│   └── main.go                        # 主执行程序
│
├── internal/                           # 应用的内部逻辑和领域模型
│   ├── controllers                     # HTTP控制器
│   ├── services                        # 应用服务层
│   ├── models                          # 数据模型
│   ├── repositories                    # 数据访问层
│   └── middleware                      # 中间件
│
├── pkg/                                # 可复用的包和模块
│   ├── errors                          # 错误处理
│   ├── utils                          # 工具函数
│   └── config                          # 配置包
│
├── configs/                            # 配置文件和脚本
│   └── application.yml                 # 应用配置文件
│
├── tests/                              # 测试文件
│   ├── integration/                    # 集成测试
│   └── unit/                          # 单元测试
│
├── static/                             # 静态资源文件
│   ├── css                            # CSS样式文件
│   ├── js                             # JavaScript文件
│   └── img                            # 图片资源
│
├── templates/                          # HTML模板文件（如果后端需要渲染HTML）
│   └── index.html
│
├── Dockerfile                          # Docker配置文件
├── go.mod                              # Go模块文件
├── go.sum                              # Go模块依赖校验文件
├── .gitignore                          # Git忽略文件配置
└── README.md                            # 项目说明文件
~~~