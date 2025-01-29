my-project/
├── api/           // API 定义 (proto, api)
│   └── user/
│       ├── user.api
│       └── user.proto
├── internal/      // 内部应用代码
│   ├── config/    // 配置
│   │   └── config.go
│   ├── handler/   // 处理器 (http handler)
│   │   └── userhandler.go
│   ├── logic/     // 业务逻辑
│   │   └── userlogic.go
│   ├── middleware/ // 中间件
│   │   └── auth.go
│   ├── model/     // 数据模型
│   │   └── usermodel.go
│   ├── rpc        // rpc client
│   │   └── user/
│   │       └── user.go
│   ├── svc        // 服务上下文
│   │   └── servicecontext.go
│   └── types/     // request, response types
│       └── usertypes.go
├── etc/           // 配置文件
│   └── my-service.yaml
├── scripts/       // 脚本
│   └── ...
├── go.mod
└── go.sum