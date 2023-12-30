概述

---
本文档为chatGPT镜像站的AI聊天服务接口文档，该服务整合了OpenAI和百度的聊天功能。后端系统可以通过这些接口，实现与用户的AI驱动对话。目前支持的AI模型包括OpenAI的GPT模型和百度的聊天服务。

接口说明

---
聊天接口
- URL: /chat
- 方法: POST
- 描述: 发送消息到AI服务，并获取AI的回复。

请求参数
请求体需要包含以下字段：
- service: 字符串，指定要使用的AI服务，可以是openai或baidu。
- interface_name:字符串，指定调用的接口名称，用于服务内部路由。
- messages: 消息数组，每个消息包含以下字段：
  - role: 字符串，消息的角色，可以是user或assistant。
  - content: 字符串，消息内容。

对于OpenAI服务，还需要额外的字段：
- model: 字符串，指定使用的OpenAI模型，如gpt-3.5-turbo、gpt-4-1106-preview或gpt-4。

请求示例（OpenAI）
- 单轮对话示例：
{
  "service": "OpenAI",
  "interface_name": "chatGPT_1",
  "model": "gpt-3-turbo",
  "messages": [
    {
      "role": "user",
      "content": "Your knowledge cutoff is until when?"
    }
  ]
}


- 多轮对话示例：
{
  "service": "OpenAI",
  "interface_name": "chatGPT_1",
  "model": "gpt-4-1106-preview",
  "messages": [
    {
      "role": "user",
      "content": "Hello!"
    },
    {
      "role": "assistant",
      "content": "Hello! How can I help you today?."
    },
    {
      "role": "user",
      "content": "Please repeat what you have said."
    }
  ]
}


请求示例（百度）
- 单轮对话示例：
{
  "service": "Baidu",
  "interface_name": "WenXinYiYan_1",
  "messages": [
    {
      "role": "user",
      "content": "Please introduce yourself."
    }
  ]
}


- 多轮对话示例：
{
  "service": "Baidu",
  "interface_name": "WenXinYiYan_1",
  "messages": [
    {
      "role": "user",
      "content": "Hello!"
    },
    {
      "role": "assistant",
      "content": "Hello! How can I help you today?."
    },
    {
      "role": "user",
      "content": "Please translate what you have said to Chinese."
    }
  ]
}


响应参数
响应体将包含以下字段：
- id: 字符串，聊天会话的唯一标识符。
- object: 字符串，对象类型。
- created: 整数，响应创建的时间戳。
- model: 字符串，用于响应的语言模型。
- choices: 选择数组，每个选择包含以下字段：
  - message: 包含以下字段的消息对象：
    - role: 字符串，消息角色。
    - content: 字符串，消息内容。
  - finish_reason: 字符串，模型停止生成文本的原因。

响应示例
{
  "id": "chat-xxxxxxxx",
  "object": "chat",
  "created": 1617986915,
  "model": "gpt-4",
  "choices": [
    {
      "message": {
        "role": "assistant",
        "content": "你好！有什么可以帮助你的吗？"
      },
      "finish_reason": "length"
    }
  ]
}

错误码

---
以下是可能的HTTP状态码及其含义：
- 200 OK: 请求成功。
- 400 Bad Request: 服务器无法理解请求格式。
- 401 Unauthorized: 认证失败，无效的API密钥。
- 403 Forbidden: 服务器理解请求但拒绝执行。
- 404 Not Found: 请求的资源不存在。
- 405 Method Not Allowed: 请求的方法不被允许。
- 500 Internal Server Error: 服务器内部错误，无法完成请求。

以上错误码大多为GPT假设的，请无视。

版本信息

---
- 当前接口版本：version_2.0
- 版本新增：添加了多Key管理的JSON结构基础
- 发布日期：2023年12月21日
