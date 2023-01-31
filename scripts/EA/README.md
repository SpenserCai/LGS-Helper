<!--
 * @Author: SpenserCai
 * @Date: 2023-01-30 23:51:56
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-01-31 10:16:12
 * @Description: file content
-->
<div align="center">

# EaUnLockV2 (EA DLC UnLocker V2 - Linux)

<img src="https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white" />
<img src="https://img.shields.io/badge/Steam-000000?style=for-the-badge&logo=steam&logoColor=white" />
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" /> 


<!-- prettier-ignore-start -->
<!-- markdownlint-disable-next-line MD036 -->
基于原始的<a  href="https://sims.tarac.nl/the-sims-4/the-sims-4-free-downloads/add-pirated-dlcs-to-your-legit-the-sims-4-game/">EA DLC UnLocker V2</a>开发，在Linux实现了解锁EA游戏DLC的功能，后期还会继续优化和集成。
<!-- prettier-ignore-end -->

</div>

<p align="center">
  <a href="./README_EN.md">English</a>
</p>

## 支持情况
  |                         游戏名                       | 状态 |                 已测试系统                                |
  | :-----------------------------------------------: | :--: | :-----------------------------------------------------: |
  |                    The Sims 4                     |  ✅  |                  Ubuntu 22.04,SteamOS                   |

## 1.安装The Sims 4
从steam免费下载

## 2.安装GE-Proton
https://github.com/GloriousEggroll/proton-ge-custom


## 3.首次运行The Sims 4
直接运行一次The Sim4 并且给小人安好家，然后保存退出

## 4.解锁DLC
```bash
git clone https://github.com/SpenserCai/LGS-Helper.git
cd LGS-Helper/scripts/EA
chmod +x EAUnLockV2
./EAUnLockV2
```

## 5.配置wine
程序运行后会弹出wine配置窗口，在函数库中添加version，并且确保下表显示的是原装先于内建，应用保存

## 6.运行The Sims 4
如果没有购买按钮说明解锁成功

## 7.安装DLC
从<a href="https://sims.tarac.nl/the-sims-4/the-sims-4-free-downloads/add-pirated-dlcs-to-your-legit-the-sims-4-game/">EA DLC UnLocker V2</a>下载DLC镜像（找到“All DLC’s in 1 iso download”，下载种子文件），将镜像解压或者挂载，把所有EP SP FP GP 开头的文件夹复制到The Sims 4的安装目录下，即可完成安装。