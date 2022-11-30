# GoFrame Template For SingleRepo

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.

### gf gen dao
```
gf gen dao -c -v
```
### Linux 上下载并安装Go
```
1、rm -rf /root/go /root/.config/go /usr/local/go（删除旧版本，如果有需要的话）
2、wget https://dl.google.com/go/go1.19.3.linux-amd64.tar.gz
3、sudo tar -C /usr/local -xf go1.19.3.linux-amd64.tar.gz
4、vim ~/.bash_profile
5、export PATH=$PATH:/usr/local/go/bin
6、source ~/.bash_profile
7、go env -w GOPROXY="https://goproxy.cn,direct"
```
### fork 第三方源码
```
1、修改代码后，修改go.mod文件中的包名为自己的包名
2、go mod tidy
3、提交代码
4、git tag v*.*.*
5、git push --tags
```
### CentOS 上安装Docker
```
1、卸载旧版本（如果有需要的话）
卸载安装包: sudo yum remove docker-ce docker-ce-cli containerd.io
删除资源: sudo rm -rf /var/lib/docker /var/lib/containerd /run/docker /etc/docker /usr/libexec/docker
2、CentOS下如果安装了Podman，需要先卸载Podman
yum erase podman buildah
3、Docker官方的镜像库比较慢，将镜像源设置为国内的站点
新建docker文件夹: mkdir /etc/docker
新建daemon.json文件: vim /etc/docker/daemon.json，输入如下内容
{
    "registry-mirrors" : [
        "https://registry.docker-cn.com",
        "https://docker.mirrors.ustc.edu.cn",
        "http://hub-mirror.c.163.com",
        "https://cr.console.aliyun.com/"
    ]
}
4、安装Docker
curl -sSL https://get.daocloud.io/docker | sh
5、Docker只能被root用户使用，可以使用下面的命令取消权限限制
sudo gpasswd -a <你的用户名> docker
6、重启Docker服务后，关闭当前的命令行，重新打开新的命令行就可以了
sudo systemctl restart docker
```
### docker 常用命令总结
```
从镜像库中查找镜像: docker search 镜像名称 (最好选择官方OFFICIAL的镜像，这样的镜像最稳定一些)
下载新的镜像: docker pull 镜像名称:版本号
启动docker: sudo systemctl start docker
停止docker: sudo systemctl stop docker
重启docker: sudo systemctl restart docker
查看docker状态: sudo systemctl status docker
开机启动: sudo systemctl enable docker
查看docker概要信息: docker info
查看docker帮助文档: docker --help
docker创建镜像: docker build -t <name>:<tag> .
docker指定Dockerfile文件创建镜像: docker build -f <Dockerfile文件> -t <name>:<tag> .
启动docker容器: docker run -d -p <hostPort>:<containerPort> <imageID>
进入docker容器: docker exec -it <containerName>/<containerID> /bin/sh
查看当前所有镜像: docker images
删除镜像，通过镜像的id来指定删除谁: docker rmi <imageID>
删除镜像id为<None>的镜像: docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
删除全部镜像: docker rmi $(docker images -q)
查看所有容器: docker ps
查看所有容器，包括关闭的: docker ps -a
停止容器: docker stop <containerID>
启动容器: docker start <containerID>
重启容器: docker restart <containerID>
停止所有的容器: docker stop $(docker ps -a -q)
删除容器，通过容器的id来指定删除谁: docker rm <containerID>
删除所有的容器: docker rm $(docker ps -a -q)
获取容器的日志: docker logs -f <containerName>/<containerID>
文件从容器中拷贝到宿主机中: docker cp <containerID>:SRC_PATH DEST_PATH
文件从宿主机中拷贝到容器中: docker cp SRC_PATH <containerID>:DEST_PATH
查看容器中运行的进程信息: docker top <containerName>/<containerID>
列出指定的容器的端口映射: docker port <containerName>/<containerID>
```
