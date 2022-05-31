FROM centos:7.9.2009
MAINTAINER liusuxian "382185882@qq.com"
COPY utility/fonts/* /usr/share/fonts/ChineseFonts/

###############################################################################
#                                INSTALLATION
###############################################################################
# 设置固定的项目路径
ENV WORKDIR /lczx
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

RUN yum install vim -y
RUN yum install java -y
RUN yum install deltarpm -y &&\
    yum install libreoffice -y &&\
    yum install libreoffice-headless -y &&\
    yum install libreoffice-writer -y &&\
    yum install ImageMagick -y &&\
    export DISPLAY=:0.0

# 添加应用可执行文件，并设置执行权限
ADD ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD manifest/config $WORKDIR/config

# 创建 cache 目录以及子目录
RUN mkdir -p $WORKDIR/cache/convert/
RUN mkdir -p $WORKDIR/cache/download/
RUN mkdir -p $WORKDIR/cache/local/
RUN mkdir -p $WORKDIR/cache/pdf/

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main

# ------------------------------------本地打包镜像---------------------
# gf docker main.go -t lczx:v1.0
# docker run -d -p 8000:8000 <imageID>
