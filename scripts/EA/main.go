/*
 * @Author: SpenserCai
 * @Date: 2023-01-30 17:53:47
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-05 10:39:04
 * @Description: file content
 */
package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/andygrunwald/vdf"
)

// 声明一个结构体，有GamePath和pfxPath两个字段
type SteamApp struct {
	GamePath string
	pfxPath  string
}

type WineDllOverrides struct {
	DllName string
	Mode    string
}

// 导入json包

// 定义一个数组存放steam默认路径.steam/steam和.local/share/Steam
var steamPath = []string{
	".steam/steam",
	".local/share/Steam",
}

// 通过环境变量获取home路径
var homePath = os.Getenv("HOME")

// 读取指定appid和游戏名的steamapps路径
func GetSteamAppsPath(appid string, gameName string) (SteamApp, error) {
	// 定义一个结构体变量,默认值为nil
	steamApp := SteamApp{}
	// 将steamPath数组中的路径和homePath拼接
	for i := 0; i < len(steamPath); i++ {
		appPath := homePath + "/" + steamPath[i] + "/steamapps"
		// 判断路径是否存在libaryfolders.vdf文件
		if _, err := os.Stat(appPath + "/libraryfolders.vdf"); err == nil {
			// 读取libaryfolders.vdf文件，遍历第一层
			f, err := os.Open(appPath + "/libraryfolders.vdf")
			if err != nil {
				return steamApp, err
			}
			defer f.Close()
			p := vdf.NewParser(f)
			v, err := p.Parse()
			if err != nil {
				return steamApp, err
			}
			// fmt.Println(v)
			// 遍历map->libraryfolders
			for _, value := range v["libraryfolders"].(map[string]interface{}) {
				// 判断map的apps中是否有名为appid的键，如果存在则设置steamApp的GamePath为map的path+"/steamapps/common/"+gameName
				if _, ok := value.(map[string]interface{})["apps"].(map[string]interface{})[appid]; ok {
					tmpGamePath := value.(map[string]interface{})["path"].(string) + "/steamapps/common/" + gameName
					// 判断GamePath目录是否存在
					if _, err := os.Stat(tmpGamePath); err != nil {
						return steamApp, err
					}
					steamApp.GamePath = tmpGamePath
					tmpPfxPath := value.(map[string]interface{})["path"].(string) + "/steamapps/compatdata/" + appid + "/pfx"
					// 判断pfxPath目录是否存在
					if _, err := os.Stat(tmpPfxPath); err != nil {
						for _, localPath := range steamPath {
							tmpPfxPath = homePath + "/" + localPath + "/steamapps/compatdata/" + appid + "/pfx"
							if _, err := os.Stat(tmpPfxPath); err == nil {
								steamApp.pfxPath = tmpPfxPath
								break
							}
						}
					} else {
						steamApp.pfxPath = tmpPfxPath
					}
					steamApp.pfxPath = tmpPfxPath
					return steamApp, nil
				}
			}

		}
	}
	return steamApp, nil
}

func CopyFile(src string, dst string) error {
	// 读取源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	// 拷贝文件
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

func UnLockEaGameDlc(steamApp SteamApp) error {
	// version.dll劫持
	eaDesktopBasePath := steamApp.pfxPath + "/drive_c/Program Files/Electronic Arts/EA Desktop"
	if _, err := os.Stat(eaDesktopBasePath); err != nil {
		return err
	}
	if _, err := os.Stat(eaDesktopBasePath + "/EA Desktop"); err == nil {
		CopyFile("./EaUnLockerTool/ea_desktop/version.dll", eaDesktopBasePath+"/EA Desktop/version.dll")
		fmt.Printf("From %s to %s \n", "./EaUnLockerTool/ea_desktop/version.dll", eaDesktopBasePath+"/EA Desktop/version.dll")
	} else {
		CopyFile("./EaUnLockerTool/ea_desktop/version.dll", eaDesktopBasePath+"/version.dll")
		fmt.Printf("From %s to %s \n", "./EaUnLockerTool/ea_desktop/version.dll", eaDesktopBasePath+"/version.dll")
	}
	// 创建配置文件
	// 递归创建steamApp.pfxPath+"/drive_c/users/steamuser/AppData/Roaming/anadius/EA DLC Unlocker v2目录
	eaUnlockConfig := steamApp.pfxPath + "/drive_c/users/steamuser/AppData/Roaming/anadius/EA DLC Unlocker v2"
	fmt.Printf("eaUnLockConfig: %s \n", eaUnlockConfig)
	if err := os.MkdirAll(eaUnlockConfig, 0755); err != nil {
		return err
	}
	// 复制配置文件
	// 用/分割steamApp.GamePath，取最后一个元素作为游戏名
	gameConfigName := "g_" + strings.Split(steamApp.GamePath, "/")[len(strings.Split(steamApp.GamePath, "/"))-1] + ".ini"
	CopyFile("./EaUnLockerTool/"+gameConfigName, eaUnlockConfig+"/"+gameConfigName)
	CopyFile("./EaUnLockerTool/config.ini", eaUnlockConfig+"/config.ini")
	fmt.Printf("From %s to %s \n", "./EaUnLockerTool/"+gameConfigName, eaUnlockConfig+"/"+gameConfigName)
	fmt.Printf("From %s to %s \n", "./EaUnLockerTool/config.ini", eaUnlockConfig+"/config.ini")
	return nil

}

func GetGeProtonPath(steamApp SteamApp) (string, error) {
	steamdataPath := strings.Split(steamApp.pfxPath, "/pfx")[0]
	// 读取steamdataPath+"/config_info"文件
	configInfo, err := os.ReadFile(steamdataPath + "/config_info")
	if err != nil {
		return "", err
	}
	// 第一行是版本号，第二行是proton路径
	protonVersion := strings.Split(string(configInfo), "\n")[0]
	protonPath := strings.Split(string(configInfo), "\n")[1]
	// 去除第二行protonVersion后面的部分
	protonPath = strings.Split(protonPath, protonVersion)[0] + protonVersion + "/proton"
	return protonPath, nil

}

func UpdataWineCfg(steamApp SteamApp) error {
	// 裁减steamApp.GamePath字符串/steamapps以及后面的内容
	steamInstallPath := strings.Split(steamApp.GamePath, "/steamapps")[0]
	// 去掉steamApp.pfxPath字符串的/pfx
	steamdataPath := strings.Split(steamApp.pfxPath, "/pfx")[0]
	// protonPath
	protonPath, err := GetGeProtonPath(steamApp)
	if err != nil {
		return err
	}
	commandString := fmt.Sprintf("STEAM_COMPAT_CLIENT_INSTALL_PATH=\"%s\" STEAM_COMPAT_DATA_PATH=\"%s\" WINEPREFIX=\"%s\" \"%s\" run %s/drive_c/windows/system32/winecfg.exe", steamInstallPath, steamdataPath, steamApp.pfxPath, protonPath, steamApp.pfxPath)
	fmt.Println(commandString)
	// 异步执行命令
	cmd := exec.Command("sh", "-c", commandString)
	cmd.Start()
	return nil

}

func UpdataDllOverrides(steamApp SteamApp) error {
	// TODO:通过修改该[Software\\Wine\\DllOverrides]下的"version"="native,builtin"来实现
	// 读取steamApp.pfxPath+"/user.reg"文件
	userReg, err := os.ReadFile(steamApp.pfxPath + "/user.reg")
	if err != nil {
		return err
	}
	// 找到[Software\\Wine\\DllOverrides]开头的行号
	dllOverridesLine := 0
	dllOverridesEndLine := 0
	versionSet := "\"version\"=\"native,builtin\""
	dllOverridesFlag := "[Software\\\\Wine\\\\DllOverrides]"
	for i := 0; i < len(strings.Split(string(userReg), "\n")); i++ {
		if strings.Contains(strings.Split(string(userReg), "\n")[i], dllOverridesFlag) {
			dllOverridesLine = i
			break
		}
	}
	// 找到[Software\\Wine\\DllOverrides]结尾的行号,结尾行没有任何内容
	for i := dllOverridesLine + 1; i < len(strings.Split(string(userReg), "\n")); i++ {
		if strings.Contains(strings.Split(string(userReg), "\n")[i], " ") {
			dllOverridesEndLine = i
			break
		}
	}
	// fmt.Println(strings.Split(string(userReg), "\n")[dllOverridesLine : dllOverridesLine+50])
	fmt.Printf("dllOverridesLine:%d dllOverridesEndLine:%d", dllOverridesLine, dllOverridesEndLine)
	// 在dllOverridesLine和dllOverridesEndLine之间查找versionSet是否存在
	versionSetLine := -1
	for i := dllOverridesLine; i < dllOverridesEndLine; i++ {
		if strings.Contains(strings.Split(string(userReg), "\n")[i], versionSet) {
			versionSetLine = i
			break
		}
	}
	// 如果不存在则在dllOverridesEndLine之前插入versionSet
	if versionSetLine == -1 {
		userRegStringArray := strings.Split(string(userReg), "\n")
		userRegStringArray = append(userRegStringArray[:dllOverridesEndLine-1], append([]string{versionSet}, userRegStringArray[dllOverridesEndLine-1:]...)...)
		userRegString := strings.Join(userRegStringArray, "\n")
		// 写入steamApp.pfxPath+"/user.reg.local"文件
		err := os.WriteFile(steamApp.pfxPath+"/user.reg", []byte(userRegString), 0644)
		if err != nil {
			return err
		}
	}
	return nil

}

func main() {
	// 将steamPath数组中的路径和homePath拼接
	for i := 0; i < len(steamPath); i++ {
		// 输出
		fmt.Println(homePath + "/" + steamPath[i])
	}
	steamApp, err := GetSteamAppsPath("1222670", "The Sims 4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("GamePath: %s\n", steamApp.GamePath)
	fmt.Printf("pfxPath: %s\n", steamApp.pfxPath)
	unLockErr := UnLockEaGameDlc(steamApp)
	if unLockErr != nil {
		fmt.Println(unLockErr)
	}
	// UpdataWineCfg(steamApp)
	updataDllErr := UpdataDllOverrides(steamApp)
	if updataDllErr != nil {
		fmt.Println(updataDllErr)
	}
}
