## TODO
1. 验证抓取内容可行性
2. 从csv中读取所有url，并处理数据
3. 内容排版（作者，日期，问题，图片）
4. 内容处理（防重复，增量更新，等等）
5. 数据格式设计
6. 数据库存储（文字，图片，视频，排版）
7. 并发获取数据
8. 代理（内网服务器不能访问x）
9. 前端展示
10. 测试
11. 部署脚本



my-project/
* api/           // API 定义 (proto, api)
    * user/
        * user.api
        * user.proto
* internal/      // 内部应用代码
    * config/    // 配置
        * config.go
    * handler/   // 处理器 (http handler)
        * userhandler.go
    * logic/     // 业务逻辑
        * userlogic.go
    * middleware/ // 中间件
        * auth.go
    * model/     // 数据模型
        * usermodel.go
    * rpc        // rpc client
        * user/
            * user.go
    * svc        // 服务上下文
        * servicecontext.go
    * types/     // request, response types
        * usertypes.go
* etc/           // 配置文件
    * my-service.yaml
* scripts/       // 脚本
    * ...
* go.mod
* go.sum