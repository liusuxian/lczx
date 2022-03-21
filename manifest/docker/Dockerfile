FROM centos:7.9.2009
MAINTAINER liusuxian "382185882@qq.com"

###############################################################################
#                                INSTALLATION
###############################################################################
# 设置固定的项目路径
ENV WORKDIR /lczx
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

# 添加应用可执行文件，并设置执行权限
ADD ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main

# ------------------------------------本地打包镜像---------------------
# gf docker main.go -t lczx:v1.0
# docker run -d -p 8089:80 <imageID>
