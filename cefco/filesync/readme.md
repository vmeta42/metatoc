**文件同步镜像生成方法：**

首先，在主机/etc/hosts中添加上一级目录中 hosts.md中的内容

执行docker-build.sh脚本，需要三个参数：

**参数1** ：镜像的名字

**参数2** ：镜像的版本（TAG）号，跟产品包的版本一致

**参数3**：产品包的名字，产品包放置在pkg目录

**参数4**：基础镜像类型（centos7.5 或 alpine）

**参数5**：CPU架构 (amd64 or arm64)


**（1）文件采集（FileReader）**

执行：docker-build.sh  filereader 8.1.1-arm FileReader-8.1.1-2021081600.arm64.tar.gz debian arm64
docker build --build-arg BASE_OS=debian:1.0.0 --build-arg PACKAGENAME=FileReader-8.1.1-2021081600.arm64.tar.gz --build-arg NAME=filereader --build-arg PROCESS_NAME=FileReader -f dockerfile/Dockerfile-debian  -t filereader:8.1.1-arm /root/cefco/filesync


**（2）文件写入（FileWriter）**

执行：docker-build.sh  filewriter 8.1.0 FileWriter-8.1.0-2021052700.x86_64.el7.tar.gz alpine amd64

