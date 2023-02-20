# 清空release目录下所有文件和文件夹
###
 # @Author: SpenserCai
 # @Date: 2023-02-01 20:53:40
 # @version: 
 # @LastEditors: SpenserCai
 # @LastEditTime: 2023-02-11 15:30:00
 # @Description: file content
### 
# 判断当前目录是否是LGS-Helper，如果不是则退出，防止误操作，通过pwd命令获取当前目录，匹配最后一个/后的字符串
if [ `pwd | awk -F '/' '{print $NF}'` != "LGS-Helper" ]; then
    echo "当前目录不是LGS-Helper，请切换到LGS-Helper目录下执行"
    exit
fi
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
        rm -rf go.mod go.sum *.go *.md
        cd ../../../
    fi
done