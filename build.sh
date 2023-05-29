#!/bin/bash

# 判断打包目录是否存在
buildDir='bin'
if [[ -d ${buildDir} ]]; then
  # 删除这个目录
  echo '发现上次打包的残余目录，正在将其删除...'
  rm -rf ${buildDir}
  echo '删除成功！'
fi

# 创建打包目录
echo '正在创建打包目录...'
mkdir ${buildDir}
echo '创建成功！'

# 构建项目
echo '正在构建项目...'
go build .
echo '构建成功！'

# 移动可执行文件到打包目录
echo '正在移动可执行文件到打包目录...'
mv greateme_ddns ${buildDir}
echo '移动成功！'

# 创建配置文件目录
echo '正在创建配置文件目录...'
mkdir ${buildDir}/conf
echo '创建成功！'

# 复制配置文件到打包目录
echo '正在复制配置文件到打包目录...'
cp conf/config.ini ${buildDir}/conf
echo '配置文件复制成功！'

echo '打包成功！'
