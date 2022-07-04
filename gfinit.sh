#!/bin/bash

# 删除所有生成service文件
rm -rf ./internal/service ./internal/logic/logic.go

# 重新生成service文件
gf gen service

# 优化生成文件的imports
goimports -w ./internal/service