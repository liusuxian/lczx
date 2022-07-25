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
### CentOS 8 上下载并安装Go
```
1、wget https://dl.google.com/go/go1.18.4.linux-amd64.tar.gz
2、sudo tar -C /usr/local -xf go1.18.4.linux-amd64.tar.gz
3、vim ~/.bash_profile
4、export PATH=$PATH:/usr/local/go/bin
5、source ~/.bash_profile
6、go env -w GOPROXY="https://goproxy.cn,direct"
```
### fork gtoken
```
1、修改代码后，修改go.mod文件中的包名为自己的包名
2、go mod tidy
3、提交代码
4、git tag v1.6.0
5、git push --tags
```
