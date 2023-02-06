/*
 * @Author: SpenserCai
 * @Date: 2023-02-02 11:55:11
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-06 14:51:24
 * @Description: file content
 */
package main

import (
	"fmt"

	"github.com/SpenserCai/lgscore"
)

// atlv download:https://cdn.altv.mp/launcher/release/x64_win32/altv.zip
func main() {
	gtavApp := lgscore.SteamApp{
		AppId: "271590",
	}
	gtavApp.InitSteamApp()
	fmt.Println(gtavApp)
}
