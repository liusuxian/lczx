FROM centos:7.9.2009
MAINTAINER liusuxian "382185882@qq.com"
COPY utility/fonts/* /usr/share/fonts/

###############################################################################
#                                INSTALLATION
###############################################################################
# 设置固定的项目路径
ENV WORKDIR /lczx
RUN yum install kde-l10n-Chinese -y
RUN yum install glibc-common -y
RUN localedef -c -f UTF-8 -i zh_CN zh_CN.utf8
ENV LANG zh_CN.UTF-8
ENV LC_ALL zh_CN.UTF-8
ENV TZ=Asia/Shanghai

RUN yum install wget -y && wget https://mirrors.nju.edu.cn/tdf/libreoffice/stable/7.3.3/rpm/x86_64/LibreOffice_7.3.3_Linux_x86-64_rpm.tar.gz
RUN tar -zxvf LibreOffice_7.3.3_Linux_x86-64_rpm.tar.gz
RUN yum install LibreOffice_7.3.3_Linux_x86-64_rpm/RPMS/* -y
ENV PATH /opt/libreoffice7.3/program:$PATH

RUN yum install vim -y
RUN yum install java -y
RUN yum install deltarpm -y &&\
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
