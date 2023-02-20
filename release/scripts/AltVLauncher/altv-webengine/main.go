/*
 * @Author: SpenserCai
 * @Date: 2023-02-11 00:14:10
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-11 15:46:19
 * @Description: file content
 */
package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

// 获取当前程序的路径
func getCurrentPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return filepath.Dir(path)
}

// 获取当前传入的参数,并用空格连接成字符串
func getArgs() string {
	argStr := ""
	for i := 1; i < len(os.Args); i++ {
		argStr += os.Args[i] + " "
	}
	return argStr
}

func main() {
	args := getArgs() + " --in-process-gpu"
	cmd := exec.Command(getCurrentPath()+"\\altv-webengine.old.exe", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
