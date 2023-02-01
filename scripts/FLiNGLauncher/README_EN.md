<!--
 * @Author: SpenserCai
 * @Date: 2023-02-01 10:29:02
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-01 23:04:59
 * @Description: file content
-->
<div align="center">

# FLiNGLauncher

<img src="https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white" />
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" /> 
<img src="https://img.shields.io/badge/Steam-000000?style=for-the-badge&logo=steam&logoColor=white" />
<img src="https://img.shields.io/badge/Epic%20Games-313131?style=for-the-badge&logo=Epic%20Games&logoColor=white" />


<!-- prettier-ignore-start -->
<!-- markdownlint-disable-next-line MD036 -->
FLiNGLauncher for Linux distributions, supports running Luna on SteamDeck and Linux. As we all know, running games on Linux depends on Proton/Wine, so different games are required to start, and a series of commands are input. FLiNGLauncher was born to solve this problem. It can target the specified game ID and platform by specifying the game ID and platform. The game runs Fling Trainer.
<!-- prettier-ignore-end -->

</div>

<p align="center">
  <a href="./README.md">中文</a>
</p>

## 1.Foreword
There is a saying in the Jianghu: "You may not know who <a href="https://flingtrainer.com/">FLiNG</a> is, but you must have used the modifier developed by him!" . We used to use complex CE, and it took half a day to adjust the value every time, until FLiNG developed Fling Trainer, everything became easier, and the disciples of Fling Shadow Sect have spread all over the world since then!

On Windows system, we can use Fling Trainer very easily, but on Linux, it needs quite complicated commands and finding a bunch of paths to run normally, so LGS-Helper provides you with a very simple solution: FLiNGLauncher, You only need to provide the game ID, and the program will automatically download, update, and run the modifier.


## 2.Instructions
I have compiled executable files for you, you can use them directly, and of course you can also compile them yourself.
```bash
git clone https://github.com/SpenserCai/LGS-Helper.git
cd LGS-Helper/release/scripts/FLiNGLauncher
chmod +x ./FLiNGLauncher
./FLiNGLauncher -appid SteamGameID
```
If you need to force a re-download of modifiers, you can use the `-redown` parameter.
```bash
./FLiNGLauncher -appid SteamGameID -redown
```

## 3.Compatibility
I tested Monster Hunter Rise and Monster Hunter World, and they can download, update, and run the modifiers normally. If you find any modifiers that cannot run normally, please file an issue.

## 4.Special Thanks
  <a href="https://flingtrainer.com/" title="风灵月影"><img src="https://flingtrainer.com/cn/community/data/avatars/l/0/1.jpg?1584477493" width=70px /></a>