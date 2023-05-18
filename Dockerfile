FROM golang:1.19
ENV env1=env1value
ENV env2=env2value

MAINTAINER blackfish
# 仅指定镜像元数据内容
LABEL tomatoList=1.0.0
# RUN git clone git@github.com:cauliflower-beep/tomatoList.git 克隆远程仓库作为镜像的源代码和依赖项
# 添加本地仓库作为镜像的源代码和依赖项 目标路径需要使用绝对路径
COPY . /tomatoList
WORKDIR /tomatoList

# 配置代理 从代理拉取go依赖
RUN export GOPROXY=https://mirrors.aliyun.com/goproxy/ && go mod tidy

# 后面那个 tomatoList 指的是go mod中的模块名
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tomatoList tomatoList
EXPOSE 9090
CMD ["./tomatoList","conf/config.ini"]
# 执行 docker build -f Dockerfile 创建镜像
# 或者执行 docker build -t tomatoList:1.0.0 .  使用当前目录作为上下文，构建一个名为tomatoList,标签为1.0.0的镜像。最后一个点表示使用当前目录作为上下文，可替换为其他目录路径
# 上下文路径传 . docker在构建镜像的时候，会自动去当前目录下找 Dockerfile文件 来构建镜像
# 上下文指docker引擎在构建镜像时使用的文件和目录的集合，上下文中的所有文件和目录都会被打包发送到docker引擎，然后在其中构建镜像
# 如果镜像或容器已存在 先docker rm 容器id 再docker rmi 镜像 ，删除容器及镜像之后再构建新镜像