# 构建 使用golang:1.16版本
FROM golang:1.15 as build

# 容器环境变量添加
ENV GOPROXY=https://goproxy.cn,direct

RUN git clone https://github.com/wangcheng1018/doServer.git

ADD . $GOPATH/src/doServer
RUN go build.

EXPOSE 9091

ENTRYPOINT ["/doServer"]