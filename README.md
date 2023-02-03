<!--
 * @Author: SpenserCai
 * @Date: 2023-01-30 17:53:03
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-02 12:08:52
 * @Description: file content
-->
<div align="center">

# LGS-Helper
<img src="https://img.shields.io/github/license/SpenserCai/LGS-Helper?color=green" alt="license">
<img src="https://img.shields.io/badge/Go-1.19+-blue" alt="go">
<a href="https://jq.qq.com/?_wv=1027&k=htcRNUvM">
    <img src="https://img.shields.io/badge/QQ%E7%BE%A4-246554357-blueviolet?style=flat-square" alt="QQ Chat Group">
</a>
![GitHub all releases](https://img.shields.io/github/downloads/SpenserCai/LGS-Helper/total)
<a href="https://discord.gg/3P7K6EzYRW">
    <img src="https://discordapp.com/api/guilds/1070553912156885203/widget.png?style=shield" alt="Discord Server">
</a>

LGS-Helper(Linux Game Script) A game tool under Linux written in Golang, designed to help linux gamers, steamdeck players to solve, localize, unlock and other functions under non-windows platforms.

</div>

# Scripts

  |                         Script                                       |                            Detail                         |
  | :------------------------------------------------------------------: | :-------------------------------------------------------: |
  |  <a href="./scripts/EA">EA DLC UnLocker V2 - Linux</a>               |                 Unocker EA DLC on Linux                   |
  |  <a href="./scripts/FLiNGLauncher">FLiNGLauncher</a>                 |                 Fling Trainer on Linux                    |

# Build
You need go 1.19+ to build this project.
```bash
git clone https://github.com/SpenserCai/LGS-Helper.git
cd LGS-Helper
chmod +x build.sh
sh ./build.sh
```
then you will get the binary file in the `release` directory.

# Usage
## Clone Project
If you haven't got the project yet, you can get it through the following command and switch to the project directory:
```bash
git clone https:://github.com/SpenserCai/LGS-Helper.git
cd LGS-Helper
```
## Use EA DLC UnLocker V2 - Linux

```bash
cd release/scripts/EA
chmod +x ./EaUnLockerV2
./EaUnLockerV2
```

## Use FLiNGLauncher

```bash
cd release/scripts/FLiNGLauncher
chmod +x ./FLiNGLauncher
# Run Monster Hunter Rise Trainer,if you want force redownload the trainer,add -redown
./FLiNGLauncher -appid 1446780 
```

# Other Game
If you have any localizations of games that unlock need to run on Linux or SteamDeck, feel free to file an issue!
```md
Game name: xxxx
Type: Localized/Unlocked etc.
How to use it on win: https://xxxx
```
