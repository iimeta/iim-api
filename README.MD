<div align=center>

<img src="https://iim.ai/public/images/logo.png?t=20231207" width="128"/>

# 智元 IIM API

[![Release](https://img.shields.io/github/v/release/iimeta/iim-api?color=blue)](https://github.com/iimeta/iim-api/releases)
[![License](https://img.shields.io/static/v1?label=license&message=MIT&color=green)](https://github.com/iimeta/iim-api/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/iimeta/iim-api)](https://goreportcard.com/report/github.com/iimeta/iim-api)

</div>

`智元 IIM` 是一款开源的网页版即时聊天系统, 同时拥有AI聊天对话功能, 本项目基于[Lumen IM](https://github.com/gzydong/go-chat)使用[GoFrame](https://github.com/gogf/gf)框架开发

## 功能介绍
- 支持 OpenAI 全模型

## 项目预览

- 地址： [https://iim.ai](https://iim.ai)
- 账号： user1@iim.ai 或 user2@iim.ai
- 密码： 123456

## Docker部署
### [传送门](https://github.com/iimeta/iim-api/tree/docker)

## API接口

### 同OpenAI接口
```bash
curl --location 'https://api.iim.ai/v1/chat/completions' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer sk-IIMGD1unkIVxDbCi0ErDjnM0YwVuDUbJOs0XMkTtMJxAEg0N' \
--data '{
    "model": "gpt-3.5-turbo-16k",
    "stream": true,
    "messages": [
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "hi"}
    ]
}'
```

### 获取使用额度接口
```bash
curl --location 'https://api.iim.ai/v1/token/usage' \
--header 'Authorization: Bearer sk-IIMGD1unkIVxDbCi0ErDjnM0YwVuDUbJOs0XMkTtMJxAEg0N'
```
响应结果
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "usage_count": 1,
        "used_tokens": 11,
        "total_tokens": 10000
    }
}
```

## 开源地址
| 仓库 | 前端                              | 后端                                 | API                               |
| ---- | --------------------------------- | ------------------------------------ | --------------------------------- |
| 主库 | https://github.com/iimeta/iim-web | https://github.com/iimeta/iim-client | https://github.com/iimeta/iim-api |
| 码云 | https://gitee.com/iimeta/iim-web  | https://gitee.com/iimeta/iim-client  | https://gitee.com/iimeta/iim-api  |

## 作者
<img src="https://iim.ai/public/images/Author.png?t=20231207" width="320" alt="作者"/>

## 交流群
<img src="https://iim.ai/public/images/WeChatGroup.jpg?t=20231207" width="320" alt="交流群"/>


![Star History Chart](https://api.star-history.com/svg?repos=iimeta/iim-api&type=Date)
