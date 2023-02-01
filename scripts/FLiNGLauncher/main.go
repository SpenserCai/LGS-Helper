/*
 * @Author: SpenserCai
 * @Date: 2023-02-01 10:23:53
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-01 21:34:56
 * @Description: file content
 */
package main

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/andygrunwald/vdf"
)

type SteamApp struct {
	GamePath string
	PfxPath  string
	GameName string
	AppId    string
}

type GameInfo struct {
	GameName        string
	GameInstallPath string
}

// 定义json配置文件结构体type为string，appid为string，flingfilename为string，flingid为int
type FlingConfigItem struct {
	Type          string `json:"type"`
	Appid         string `json:"appid"`
	FlingFileName string `json:"flingfilename"`
	FlingID       string `json:"flingid"`
}

var steamPath = []string{
	".steam/steam",
	".local/share/Steam",
}

var flingTrainerWeb = "https://flingtrainer.com/trainer/"

// 通过环境变量获取home路径
var homePath = os.Getenv("HOME")

var flingTrainerPath = homePath + "/.LGS-Helper/FLiNG"
var flingTrainerTmpPath = homePath + "/.LGS-Helper/FLiNG/tmp"
var flingTrainerConfig = homePath + "/.LGS-Helper/FLiNG/config.json"

// 初始化Fling文件夹和配置文件
func InitFLiNG() error {
	// 创建Fling文件夹
	if err := os.MkdirAll(flingTrainerPath, 0755); err != nil {
		return err
	}
	// 创建Fling临时文件夹
	if err := os.MkdirAll(flingTrainerTmpPath, 0755); err != nil {
		return err
	}
	// 配置文件不存在则创建
	if _, err := os.Stat(flingTrainerConfig); os.IsNotExist(err) {
		// 创建配置文件并且写入默认值{}
		f, err := os.Create(flingTrainerConfig)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString("{}")
		if err != nil {
			return err
		}
	}
	return nil
}

func GetGeProtonPath(steamApp SteamApp) (string, error) {
	steamdataPath := strings.Split(steamApp.PfxPath, "/pfx")[0]
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

// 从指定目录通过appid获取游戏名
func GetGameInfo(appid string, appsPath string) (GameInfo, error) {
	// 定义一个结构体变量,默认值为nil
	gameInfo := GameInfo{}
	// 读取appsPath目录下的appmanifest_appid.acf文件
	f, err := os.Open(appsPath + "/appmanifest_" + appid + ".acf")
	if err != nil {
		return gameInfo, err
	}
	defer f.Close()
	p := vdf.NewParser(f)
	v, err := p.Parse()
	if err != nil {
		return gameInfo, err
	}
	// 将v转换成map从v中获取游戏名
	gameInfo.GameName = v["AppState"].(map[string]interface{})["name"].(string)
	gameInfo.GameInstallPath = v["AppState"].(map[string]interface{})["installdir"].(string)
	return gameInfo, nil
}

func GetSteamAppsPath(appid string) (SteamApp, error) {
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
					appsPath := value.(map[string]interface{})["path"].(string) + "/steamapps"
					gameInfo, err := GetGameInfo(appid, appsPath)
					if err != nil {
						return steamApp, err
					}
					tmpGamePath := value.(map[string]interface{})["path"].(string) + "/steamapps/common/" + gameInfo.GameInstallPath
					// 判断GamePath目录是否存在
					if _, err := os.Stat(tmpGamePath); err != nil {
						return steamApp, err
					}
					steamApp.GamePath = tmpGamePath
					steamApp.PfxPath = value.(map[string]interface{})["path"].(string) + "/steamapps/compatdata/" + appid + "/pfx"
					steamApp.GameName = gameInfo.GameName
					steamApp.AppId = appid
					return steamApp, nil
				}
			}

		}
	}
	return steamApp, nil
}

// 通过steamApp获取最新版fling的下载地址
func GetLastestFlingUrl(steamApp SteamApp) (string, error) {
	// 将steamApp.GameName中的符号都去掉，将空格替换为中划线得到flingname
	// 定义符号数组
	symbol := []string{":", "™", "®", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "[", "]", "{", "}", "|", "\\", ";", ":", "'", "\"", "<", ">", ",", ".", "/", "?"}
	// 将steamApp.GameName中的符号都去掉,得到flingname
	flingname := steamApp.GameName
	for i := 0; i < len(symbol); i++ {
		flingname = strings.Replace(flingname, symbol[i], "", -1)
	}
	// 将空格替换为中划线得到flingname
	flingname = strings.Replace(flingname, " ", "-", -1)
	// 将flingname转换为小写
	flingname = strings.ToLower(flingname)
	// flingTrainerWeb+flingname，支持301重定向，支持https，得到html
	resp, err := http.Get(flingTrainerWeb + flingname)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// 读取html
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// 将html转换为字符串
	bodyStr := string(body)
	// 在body中正则查找所有https://flingtrainer.com/attachments/数字 的字符串并存入数组
	reg := regexp.MustCompile(`https://flingtrainer.com/attachments/\d+`)
	flingTrainer := reg.FindAllString(bodyStr, -1)
	// fmt.Println(flingTrainer)
	// 如果数组长度为0，说明没有找到，返回空
	if len(flingTrainer) == 0 {
		return "", nil
	}
	// 如果数组长度不为0，说明找到了，返回第一个
	return flingTrainer[0], nil
}

// 判断是否安装了最新版风灵月影
func IsInstallLastestFling(steamApp SteamApp, lastestUrl string) (bool, error) {
	// 从lastestUrl中正则查找最新版风灵月影的版本号（用/分割，取最后一个）
	flingVerison := strings.Split(lastestUrl, "/")[len(strings.Split(lastestUrl, "/"))-1]
	// 读取json文件，将内容转换成map
	jsonData, err := os.ReadFile(flingTrainerConfig)
	if err != nil {
		return false, err
	}
	var flingTrainerMap map[string]interface{}
	err = json.Unmarshal(jsonData, &flingTrainerMap)
	if err != nil {
		return false, err
	}
	// 判断是否存在名为appid的键，如果不存在，返回false
	if _, ok := flingTrainerMap[steamApp.AppId]; !ok {
		return false, nil
	}
	// 如果存在，将flingTrainerMap[steamApp.AppId]内容转换成FlingConfigItem
	flingConfigItem := FlingConfigItem{}
	flingConfigItemJson, err := json.Marshal(flingTrainerMap[steamApp.AppId])
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(flingConfigItemJson, &flingConfigItem)
	if err != nil {
		return false, err
	}
	// 判断flingConfigItem.Version是否等于flingVerison，如果等于，返回true，否则返回false
	if flingConfigItem.FlingID == flingVerison {
		return true, nil
	} else {
		return false, nil
	}

}

// 通过steamApp下载或更新风灵月影
func DownOrUpdateFling(steamApp SteamApp, isReDown bool) error {
	lastestUrl, err := GetLastestFlingUrl(steamApp)
	if err != nil {
		return err
	}
	if lastestUrl == "" {
		// 返回Can't find fling trainer(Game:GameName)
		return errors.New("Can't find fling trainer(Game:" + steamApp.GameName + ")")
	}
	// 判断是否安装了最新版风灵月影
	isInstallLastestFling, err := IsInstallLastestFling(steamApp, lastestUrl)
	if err != nil {
		return err
	}
	if isInstallLastestFling && !isReDown {
		fmt.Println("Already install lastest fling trainer(Game:" + steamApp.GameName + ")")
		return nil
	}
	// 获取原始的flingConfigItem
	oldFlingConfigItem, err := GetFlingConfigItemByAppid(steamApp.AppId)
	if err != nil {
		return err
	}
	// 如果存在老的FlingFileName，删除老的FlingFileName
	if oldFlingConfigItem.FlingFileName != "" {
		// 判断文件是否存在，存在则删除
		if _, err := os.Stat(flingTrainerPath + "/" + oldFlingConfigItem.FlingFileName); err == nil {
			err = os.Remove(flingTrainerPath + "/" + oldFlingConfigItem.FlingFileName)
			if err != nil {
				return err
			}
		}
	}
	// 从lastestUrl中正则查找最新版风灵月影的版本号（用/分割，取最后一个）
	flingVerison := strings.Split(lastestUrl, "/")[len(strings.Split(lastestUrl, "/"))-1]
	// 从lastestUrl下载风灵月影支持301重定向，支持https，保存到临时目录(flingTrainerTmpPath)文件名为flingVerison+.zip
	fmt.Println("Downloading fling trainer(Game:" + steamApp.GameName + ")")
	fmt.Println("Fling trainer url:" + lastestUrl)
	resp, err := http.Get(lastestUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 创建临时文件
	flingTrainerTmpFile, err := os.Create(flingTrainerTmpPath + "/" + flingVerison + ".zip")
	if err != nil {
		return err
	}
	defer flingTrainerTmpFile.Close()
	// 将下载的内容写入临时文件
	_, err = io.Copy(flingTrainerTmpFile, resp.Body)
	if err != nil {
		return err
	}
	// 解压临时文件到flingTrainerPath,获取解压后的exe文件名，并删除临时文件
	fmt.Println("Unzip fling trainer(Game:" + steamApp.GameName + ")")
	zipReader, err := zip.OpenReader(flingTrainerTmpPath + "/" + flingVerison + ".zip")
	if err != nil {
		return err
	}
	defer zipReader.Close()
	var exeFileName string
	for _, file := range zipReader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name, ".exe") {
			exeFileName = file.Name
		}
		// 解压文件到flingTrainerPath
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()
		fileWriter, err := os.Create(flingTrainerPath + "/" + file.Name)
		if err != nil {
			return err
		}
		defer fileWriter.Close()
		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			return err
		}
	}
	// 删除临时文件
	err = os.Remove(flingTrainerTmpPath + "/" + flingVerison + ".zip")
	if err != nil {
		return err
	}
	fmt.Println("Save fling trainer(Game:" + steamApp.GameName + ") to " + flingTrainerPath + "/" + exeFileName)
	flingConfigItem := FlingConfigItem{
		FlingID:       flingVerison,
		FlingFileName: exeFileName,
		Type:          "steam",
		Appid:         steamApp.AppId,
	}
	jsonData, err := os.ReadFile(flingTrainerConfig)
	if err != nil {
		return err
	}
	var flingTrainerMap map[string]interface{}
	err = json.Unmarshal(jsonData, &flingTrainerMap)
	if err != nil {
		return err
	}
	flingTrainerMap[steamApp.AppId] = flingConfigItem
	// fmt.Println(flingTrainerMap)
	// 格式化json
	flingTrainerMapJson, err := json.MarshalIndent(flingTrainerMap, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(flingTrainerConfig, flingTrainerMapJson, 0644)
	if err != nil {
		return err
	}
	return nil

}

// 通过appid读取flingConfigItem
func GetFlingConfigItemByAppid(appid string) (FlingConfigItem, error) {
	flingConfigItem := FlingConfigItem{}
	jsonData, err := os.ReadFile(flingTrainerConfig)
	if err != nil {
		return flingConfigItem, err
	}
	var flingTrainerMap map[string]interface{}
	err = json.Unmarshal(jsonData, &flingTrainerMap)
	if err != nil {
		return flingConfigItem, err
	}
	// 读取flingConfigItem转换为FlingConfigItem对象
	flingConfigItemJson, err := json.Marshal(flingTrainerMap[appid])
	if err != nil {
		return flingConfigItem, err
	}
	err = json.Unmarshal(flingConfigItemJson, &flingConfigItem)
	if err != nil {
		return flingConfigItem, err
	}
	// fmt.Println(flingConfigItem)
	return flingConfigItem, nil
}

// 运行风灵月影
func RunFling(steamApp SteamApp) error {
	// 读取flingConfigItem转换为FlingConfigItem对象
	flingConfigItem, err := GetFlingConfigItemByAppid(steamApp.AppId)
	if err != nil {
		return err
	}
	// 风灵月影exe路径flingTrainerPath + "/" + flingConfigItem.FlingFileName
	flingExePath := flingTrainerPath + "/" + flingConfigItem.FlingFileName
	// 裁减steamApp.GamePath字符串/steamapps以及后面的内容
	steamInstallPath := strings.Split(steamApp.GamePath, "/steamapps")[0]
	// 去掉steamApp.pfxPath字符串的/pfx
	steamdataPath := strings.Split(steamApp.PfxPath, "/pfx")[0]
	// protonPath
	protonPath, err := GetGeProtonPath(steamApp)
	if err != nil {
		return err
	}
	commandString := fmt.Sprintf("STEAM_COMPAT_CLIENT_INSTALL_PATH=\"%s\" STEAM_COMPAT_DATA_PATH=\"%s\" WINEPREFIX=\"%s\" \"%s\" run \"%s\"", steamInstallPath, steamdataPath, steamApp.PfxPath, protonPath, flingExePath)
	// fmt.Println(commandString)
	fmt.Printf("Run fling trainer(Game:%s):\n", steamApp.GameName)
	fmt.Printf("  Proton:%s\n", protonPath)
	fmt.Printf("  WinePrefix:%s\n", steamApp.PfxPath)
	fmt.Printf("  FlingTrainer:%s\n", flingExePath)
	// 异步执行命令
	cmd := exec.Command("sh", "-c", commandString)
	cmd.Start()
	return nil

}

func main() {
	// 参数：--appid=582010 --redown
	// 获取appid参数
	appid := flag.String("appid", "", "Enter Steam Game Appid")
	isReDown := flag.Bool("redown", false, "Force Re-Download")
	flag.Parse()
	if *appid == "" {
		fmt.Println("Please enter the appid parameter")
		return
	}
	// TODO：添加跳过检查更新参数（不存在的修改器将无法自动下载）
	InitFLiNG()
	steamApp, err := GetSteamAppsPath(*appid)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	downErr := DownOrUpdateFling(steamApp, *isReDown)
	if downErr != nil {
		fmt.Println(err)
	}
	runErr := RunFling(steamApp)
	if runErr != nil {
		fmt.Println(err)
	}
}
