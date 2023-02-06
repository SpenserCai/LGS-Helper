/*
 * @Author: SpenserCai
 * @Date: 2023-02-02 11:55:11
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-06 14:01:32
 * @Description: file content
 */
package main

import (
	"fmt"

	"github.com/SpenserCai/lgscore"
)

// atlv download:https://cdn.altv.mp/launcher/release/x64_win32/altv.zip
func main() {
	v := lgscore.SteamApp{
		AppId: "271590",
	}
	v.InitSteamApp()
	fmt.Println(v)
}
