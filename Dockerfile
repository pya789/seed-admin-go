# 编译阶段
FROM golang:alpine AS go-builder

# 进入工作目录
WORKDIR /www
COPY . .
# 配置模块代理 国外服务器可以不配
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
# 打包 linux/AMD64 架构
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server

# 运行阶段
FROM alpine AS admin-runner

# 进入工作目录
WORKDIR /www

# 复制打包的Go文件到系统用户可执行程序目录下
COPY --from=go-builder /www/server /www
# 复制配置文件到系统用户可执行程序目录下
COPY --from=go-builder /www/config.toml /www
# 复制sql到系统用户可执行程序目录下
COPY --from=go-builder /www/sql /www
# 复制静态目录到系统用户可执行程序目录下
COPY --from=go-builder /www/wwwroot /www
# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata

# 暴露服务端口
EXPOSE 8080
# 启动服务
ENTRYPOINT ["/www/server"]