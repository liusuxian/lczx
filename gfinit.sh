#!/bin/bash

# 关闭所有关键字进程
killall main

# 删除所有生成模型文件
rm -rf ./internal/model/do ./internal/model/entity ./internal/dao

# 重新生成模型文件
gf gen dao -c -v

# 删除所有生成service文件
rm -rf ./internal/service ./internal/logic/logic.go

# 重新生成service文件
gf gen service

# 优化生成文件的imports
goimports -w ./internal/service

# 可将任意文件打包后随着可执行文件一同发布，就是发布只需要一个运行文件就可以
echo -e "\n" | gf pack manifest/config ./internal/packed/packed.go -n packed

# 启用GF热编译
gf run main.go