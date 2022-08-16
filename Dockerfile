FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

#拷贝mod并下载
COPY go.mod .
COPY go.sum .
RUN go mod download

# 拷贝代码
COPY . .

# 拷贝配置文件
COPY ./apps/app/api/etc /app/api/etc
COPY ./apps/sms/etc /app/sms/etc
COPY ./apps/user/rpc/etc /app/user/etc

RUN go build -ldflags="-s -w" -o /app/api/api apps/app/api/api.go
RUN go build -ldflags="-s -w" -o /app/sms/sms apps/sms/sms.go
RUN go build -ldflags="-s -w" -o /app/user/user apps/user/rpc/user.go



FROM scratch as chat-api

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/api/ /app/

CMD ["./api", "-f", "etc/api-api.yaml"]



FROM scratch as sms-rpc

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/sms/ /app/

CMD ["./sms", "-f", "etc/sms.yaml"]


FROM scratch as user-rpc

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user/ /app/

CMD ["./user", "-f", "etc/user.yaml"]
