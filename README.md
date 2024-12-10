# go-qwen

由于阿里百炼平台通义千问大模型没有完善的go语言示例，并且官方答复assistant是不兼容openapi sdk的。
实际使用中发现是能够支持的，所以自己写了一个demo示例，给大家做一个参考。

# apikey 生成官方教程
https://help.aliyun.com/zh/model-studio/getting-started/first-api-call-to-qwen
# apifox 共享链接
https://apifox.com/apidoc/shared-78492d5a-83f6-479c-99cb-e7d18a80fcf6


## internal 结构
```
go-qwen/
└── internal/
  ├── config/   配置文件操作
  ├── doce/     文档:数据库sql,  接口文档json-可导入apifox
  ├── handle/   接收层
  ├── logic/    逻辑层
  ├── model/    模型层
  ├── repo/     持久层
  ├── types/    请求-返回体
  ├── utils/    工具包
  └── root.go   路由组
```


## ✨  test 目录下 - 说明
- utils 中写的是获取实例的方式
- const 中写的是baseUrl 具体的千问访问url，以及调用时候用到的apiKey （需要自己申请 sk-xxxx 格式）
- -----------------------------------------------------------------------------------------------------------------
- full_step_test 完整流程
- chat_test 普通的聊天示例
- chat_stream_test 普通的聊天示例-流式输出
- assistant_test 助手增删改查
- tread_test 会话示例
- message_test 消息示例
- run_test 运行任务示例
- -----------------------------------------------------------------------------------------------------------------
## ✨ assistant 流程
    1. 创建助手 
    2. 创建会话
    3. 创建消息(在会话上创建)
    4. 创建执行任务（执行会话）
    5. 获取最后一次执行结果
