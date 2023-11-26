<div align=center>

<img src="https://iim.ai/public/images/logo.png?t=2023112623" width="128"/>

# 智元 IIM API

[![Release](https://img.shields.io/github/v/release/iimeta/iim-api?color=blue)](https://github.com/iimeta/iim-api/releases)
[![License](https://img.shields.io/static/v1?label=license&message=MIT&color=green)](https://github.com/iimeta/iim-api/blob/main/LICENSE)
[![Docker Pull](https://img.shields.io/docker/pulls/iimeta/iim-api?color=brightgreen)](https://hub.docker.com/r/iimeta/iim-api)

</div>

`智元 IIM` 是一个网页版即时聊天系统, 同时拥有AI聊天对话功能, 本项目基于[Lumen IM](https://github.com/gzydong/go-chat)使用[GoFrame](https://github.com/gogf/gf)框架开发

## 手动Docker部署
0. 需提前安装好Redis和MongoDB以及Docker服务
1. 拉取镜像: docker pull iimeta/iim-api:1.1.0
2. 创建配置文件目录: mkdir -p /data/iim-api/manifest/config
5. 进入配置文件目录: cd /data/iim-api/manifest/config
6. 下载配置文件: wget https://github.com/iimeta/iim-api/raw/docker/manifest/config/config.yaml
7. 修改配置文件: vi config.yaml 按 config.yaml 中配置项的注释, 修改对应参数
8. 启动服务: 
```shell
docker run -d \
  --network host \
  --restart=always \
  -p 11000:11000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/iim-api/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  --name iim-api \
  iimeta/iim-api:1.1.0
```

## 自动Docker部署
0. 需提前安装好Redis和MongoDB以及Docker服务
1. 下载一键部署脚本: wget https://github.com/iimeta/iim-api/raw/docker/bin/install.sh
2. 执行安装: chmod +777 install.sh && ./install.sh
3. 修改配置文件: vi /data/iim-api/manifest/config/config.yaml 按 config.yaml 中配置项的注释, 修改对应参数
4. 启动服务: chmod +777 start.sh && ./start.sh

## 备用国内镜像地址
```shell
docker pull registry.cn-hangzhou.aliyuncs.com/iimeta/iim-api:1.1.0
```
```shell
docker run -d \
  --network host \
  --restart=always \
  -p 11000:11000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/iim-api/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  --name iim-api \
  registry.cn-hangzhou.aliyuncs.com/iimeta/iim-api:1.1.0
```


![Star History Chart](https://api.star-history.com/svg?repos=iimeta/iim-api&type=Date)
