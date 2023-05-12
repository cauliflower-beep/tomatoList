# tomatoList清单

一个基于gin+gorm开发的练手小项目，通过该项目可初识go web开发该有的姿势。

前端页面基于vue和ElementUI开发，对前端不熟悉的童鞋可直接下载`templates`和`static`文件夹下的内容使用。

## 使用指南
### 下载
```bash
git clone https://github.com/Q1mi/tomatoList.git
```
### 配置MySQL
1. 在你的数据库中执行以下命令，创建本项目所用的数据库：
```sql
CREATE DATABASE tomatoList DEFAULT CHARSET=utf8mb4;
```
2. 在`tomatoList/conf/config.ini`文件中按如下提示配置数据库连接信息。

```ini
port = 9090
release = false

[mysql]
user = 你的数据库用户名
password = 你的数据库密码
host = 你的数据库host地址
port = 你的数据库端口
db = tomatoList
```

### 编译
```bash
go build
```

### 执行

Mac/Unix：
```bash
./tomatoList conf/config.ini
```
Windows:
```bash
tomatoList.exe conf/config.ini
```

启动之后，使用浏览器打开`http://127.0.0.1:9000/`即可。

![example.png](example.png)

### docker运行

将本项目打包成docker运行一般分为4步：

1. 设置dockerfile

   ```go
   FROM scratch
   ADD /tomatoList // go build生成的可执行文件 
   ADD /config.json // 根目录下的配置文件
   ENTRYPOINT ["/tomatoList"]
   ```

2. 进入main所在根目录

3. 设置环境变量

   ```go
   SET CGO_ENABLED=0
   SET GOOS=linux
   SET GOARCH=amd64
   go env //查看环境变量是否设置成功
   ```

   我没有配置全局的环境变量，而是使用了goland自带的编译配置，同样配置了上面三个编译选项。但是当我点击debug按钮之后，发现linux下的可执行文件是生成了，但是ide并没有进入运行调试状态。

   当时还一脸懵，现在想想不就是因为编译的是linux可执行文件，但又在windows上跑了吗？

   真蠢呀。。

4. 编译

   ```go
   go build
   ```

5. 打包成镜像

   ```go
   docker build . -t tomatoList
   docker push tomatoList // 推送到远程镜像服务器
   ```

6. 运行docker镜像

   ```go
   docker run -p 9090:9090 tomatoList
   ```

   

