#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh bsmaster users
# ./genModel.sh bsmaster bike
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./${tables}Model

# 数据库配置
host=192.168.0.214
port=32432
dbname=$1 # 库名
username=root
passwd=123456


echo "开始创建库：$dbname 的表：$tables"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero

#-cache=false 如果需要redis则打开
