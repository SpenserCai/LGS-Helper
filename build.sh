# 清空release目录下所有文件和文件夹
###
 # @Author: SpenserCai
 # @Date: 2023-02-01 20:53:40
 # @version: 
 # @LastEditors: SpenserCai
 # @LastEditTime: 2023-02-01 20:59:10
 # @Description: file content
### 
rm -rf release/
# 创建release目录
mkdir release
# 在release目录下创建scripts目录，并将scripts目录下所有文件复制到release/scripts目录下
mkdir release/scripts
cp -r scripts/* release/scripts/
# 在所有不为空的release/scripts的子目录下执行go mod tidy && go build
for dir in `ls release/scripts`; do
    if [ -n "$(ls -A release/scripts/$dir)" ]; then
        cd release/scripts/$dir
        go mod tidy
        go build
        rm -rf go.mod go.sum main.go
        cd ../../../
    fi
done