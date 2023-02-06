/*
 * @Author: SpenserCai
 * @Date: 2023-02-02 11:55:11
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-06 16:49:26
 * @Description: file content
 */
package main

import (
	"fmt"
	"os"

	"github.com/SpenserCai/lgscore"
	lgsutils "github.com/SpenserCai/lgscore/utility"
)

var HomePath = os.Getenv("HOME")

var AltvPath = HomePath + "/.LGS-Helper/altv"
var AltvTmpPath = AltvPath + "/tmp"

func InitAltV() error {
	// 检查AltvPath是否存在,不存在则创建
	if _, err := os.Stat(AltvPath); os.IsNotExist(err) {
		os.MkdirAll(AltvPath, os.ModePerm)
	}
	// 检查AltvTmpPath是否存在,不存在则创建
	if _, err := os.Stat(AltvTmpPath); os.IsNotExist(err) {
		os.MkdirAll(AltvTmpPath, os.ModePerm)
	}
	// 检查AltvPath中是否存在altv.exe,不存在则下载
	if _, err := os.Stat(AltvPath + "/altv.exe"); os.IsNotExist(err) {
		downErr := new(lgsutils.NewWorkLGS).Download("https://cdn.altv.mp/launcher/release/x64_win32/altv.zip", AltvTmpPath+"/altv.zip")
		if downErr != nil {
			return downErr
		}
		// 解压
		unzipedList := []string{}
		unzipErr := new(lgsutils.FileLGS).UnZip(AltvPath+"/tmp/altv.zip", AltvPath, unzipedList)
		if unzipErr != nil {
			return unzipErr
		}
		// 删除临时文件
		delErr := os.Remove(AltvTmpPath + "/altv.zip")
		if delErr != nil {
			return delErr
		}
	}
	return nil
}

// atlv download:https://cdn.altv.mp/launcher/release/x64_win32/altv.zip
func main() {
	gtavApp := lgscore.SteamApp{
		AppId: "271590",
	}
	gtavApp.InitSteamApp()
	fmt.Println(gtavApp)
	InitAltV()
}
