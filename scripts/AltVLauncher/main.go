/*
 * @Author: SpenserCai
 * @Date: 2023-02-02 11:55:11
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-06 22:35:39
 * @Description: file content
 */
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

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

func KillAltv() {
	// 杀死altv进程
	// ps -ef | grep altv.exe | grep -v grep | awk '{print $2}'
	psCmd := exec.Command("sh", "-c", "ps -ef | grep altv.exe | grep -v grep | awk '{print $2}'")
	// 一次性读取所有输出
	psOut, _ := psCmd.Output()
	// 将[]byte转换为string
	psOutString := string(psOut)
	// 将string转换为[]string
	psOutStringList := strings.Split(psOutString, "\n")
	fmt.Println(psOutStringList)
	// 如果有进程则杀死
	if len(psOutStringList) > 1 {
		// 杀死进程
		for _, pid := range psOutStringList {
			if pid != "" {
				killCmd := exec.Command("sh", "-c", "kill -9 "+pid)
				killCmd.Run()
			}
		}
	}
}

func FirstRunAltvCheck(cmd *exec.Cmd) {
	// 定时检测是否存在包含PlayGTAV.exe的进程，如果有则杀死（获取进程id调用kill -9）并且停止命令执行
	for {
		// 检测是否存在包含PlayGTAV.exe的进程
		// ps -ef | grep PlayGTAV.exe | grep -v grep | awk '{print $2}'
		psCmd := exec.Command("sh", "-c", "ps -ef | grep PlayGTAV.exe | grep -v grep | awk '{print $2}'")
		// 一次性读取所有输出
		psOut, _ := psCmd.Output()
		// 将[]byte转换为string
		psOutString := string(psOut)
		// 将string转换为[]string
		psOutStringList := strings.Split(psOutString, "\n")
		// 如果有进程则杀死
		if len(psOutStringList) > 1 {
			// 杀死altv进程
			// 杀死进程
			for _, pid := range psOutStringList {
				if pid != "" {
					killCmd := exec.Command("sh", "-c", "kill -9 "+pid)
					killCmd.Run()
				}
			}
			cmd.Process.Kill()
			break
		}
		// 如果没有进程则继续等待3秒
		time.Sleep(3 * time.Second)
	}
}

func GetRunAltvCommand(steamApp lgscore.SteamApp) string {
	commandString := fmt.Sprintf("STEAM_COMPAT_CLIENT_INSTALL_PATH=\"%s\" STEAM_COMPAT_DATA_PATH=\"%s\" WINEPREFIX=\"%s\" \"%s\" run \"%s\"",
		strings.Split(steamApp.Game.GameInstallPath, "/steamapps")[0],
		strings.Split(steamApp.PfxPath, "/pfx")[0], steamApp.PfxPath,
		steamApp.ProtonPath, AltvPath+"/altv.exe")
	return commandString
}

func FirstRunAltv(steamApp lgscore.SteamApp) error {
	// 检查是否第一次运行
	if _, err := os.Stat(AltvPath + "/lgcaltv.lock"); os.IsNotExist(err) {
		// 向altv.toml写入“branch = 'release'”，如果没有则创建
		altvToml := AltvPath + "/altv.toml"
		if _, err := os.Stat(altvToml); os.IsNotExist(err) {
			_, err := os.Create(altvToml)
			if err != nil {
				return err
			}
		}
		// 写入branch = 'release'
		altvTomlFile, err := os.OpenFile(altvToml, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}
		_, err = altvTomlFile.WriteString("branch = 'release'\n")
		if err != nil {
			return err
		}
		winePath := "Z:" + strings.ReplaceAll(steamApp.Game.GameInstallPath, "/", "\\")
		_, err = altvTomlFile.WriteString("gtapath = '" + winePath + "'\n")
		if err != nil {
			return err
		}
		altvTomlFile.Close()
		commandString := GetRunAltvCommand(steamApp)
		fmt.Println(commandString)
		// 切换到AltvPath目录下同步执行命令，如果检测到输出中有“Steam client's requirements are satisfied”则停止命令执行
		cmd := exec.Command("sh", "-c", commandString)
		cmd.Dir = AltvPath
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		go FirstRunAltvCheck(cmd)
		// 开始执行命令
		cmd.Run()
		KillAltv()
		// 设置超链接
		setLinkCmd := exec.Command("sh", "-c", "ln -s ../libs/chrome_elf.dll . &&  ln -s ../libs/icudtl.dat . &&  ln -s ../libs/libce2.dll . &&  ln -s ../libs/snapshot_blob.bin . &&  ln -s ../libs/v8_context_snapshot.bin .")
		setLinkCmd.Dir = AltvPath + "/cef"
		setLinkCmd.Run()
		//创建锁文件
		_, createLockErr := os.Create(AltvPath + "/lgcaltv.lock")
		if createLockErr != nil {
			fmt.Println(createLockErr)
		}
	}
	return nil
}

func LaunchAltv(steamApp lgscore.SteamApp) {
	commandString := GetRunAltvCommand(steamApp)
	runGtaCmd := exec.Command("sh", "-c", "steam steam://rungameid/271590")
	runGtaCmd.Run()
	// 等待10秒
	time.Sleep(10 * time.Second)
	runAltvCmd := exec.Command("sh", "-c", commandString)
	runAltvCmd.Dir = AltvPath
	runAltvCmd.Start()
}

// atlv download:https://cdn.altv.mp/launcher/release/x64_win32/altv.zip
func main() {
	gtavApp := lgscore.SteamApp{
		AppId: "271590",
	}
	gtavApp.InitSteamApp()
	fmt.Println(gtavApp)
	InitAltV()
	err := FirstRunAltv(gtavApp)
	if err != nil {
		fmt.Println(err)
	}
	LaunchAltv(gtavApp)
}
