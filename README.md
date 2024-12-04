# go-qwen

由于阿里百炼平台通义千问大模型没有完善的go语言示例，并且官方答复assistance是不兼容openapi sdk的。
实际使用中发现是能够支持的，所以自己写了一个demo test示例，给大家做一个参考。

## ✨ 说明
- utils 中写的是获取实例的方式
- const 中写的是baseUrl 具体的千问访问url，以及调用时候用到的apiKey （需要自己申请 sk-xxxx 格式）

- chat_test 普通的聊天连接