# 得到最新的 golang docker 镜像
FROM golang:latest
# 在容器内部创建一个目录来存储我们的 web 应用，接着使它成为工作目录。
#RUN mkdir -p /go/src/github.com/xufwind95/go-web-base
RUN mkdir -p /go/src/github.com/xufwind95/go-web-base
WORKDIR $GOPATH/src/github.com/xufwind95/go-web-base
# 复制 web-app 目录到容器中
ADD . $GOPATH/src/github.com/xufwind95/go-web-base
#调整容器时间
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
#RUN echo "Asia/Shanghai" > /etc/timezone
#RUN dpkg-reconfigure -f noninteractive tzdata

# 下载并安装第三方依赖到容器中
#RUN go get github.com/go-sql-driver/mysql && go get github.com/astaxie/beego && go get github.com/astaxie/beego/logs && go get gopkg.in/gomail.v2 && go get github.com/tealeg/xlsx && go get  github.com/shopspring/decimal
#RUN go get -u github.com/kardianos/govendor
#RUN govendor sync
RUN go build .
# 告诉 Docker 启动容器运行的命令
ENTRYPOINT  ["./go-web-base"]
