# go-zero-chat

go-zero微服务初尝试、websocket服务

#### docker-compose启动

1. 自建网络

```shell
docker network create --driver bridge --subnet 192.167.1.0/16 --gateway 192.167.1.1 chat
```

2. 启动

```shell
docker-compose up -d
```