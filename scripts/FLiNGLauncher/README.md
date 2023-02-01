<!--
 * @Author: SpenserCai
 * @Date: 2023-02-01 10:28:50
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-01 22:53:20
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
运行于Linux发行版上的风灵月影启动器，支持在SteamDeck和Linux上运行风灵月影。众所周知，Linux上运行游戏依赖于Proton/Wine，所以启动时需要不同的游戏，输入一大串命令，FLiNGLauncher就是为了解决这个问题而诞生的，他可以通过指定游戏ID，平台，的方式来针对指定的游戏运行风灵月影。
<!-- prettier-ignore-end -->

</div>

<p align="center">
  <a href="./README_EN.md">English</a>
</p>

## 1.前言
江湖上流传着这样一句话：“你可能不知道<a href="https://flingtrainer.com/">风灵月影</a>是谁，但你一定用过他开发的修改器！”。曾经我们用着复杂的CE，每次修改数值都要定位半天，直到风大开发了Fling Trainer一切都变得简单了，风灵月影宗弟子从此遍布全球！

在Windows系统上我们可以非常轻松的使用Fling Trainer，但是在Linux则需要相当复杂的命令，以及找到一堆路经才能正常运行，因此LGS-Helper为大家提供了一个非常简单的解决方案：FLiNGLauncher，你只需要提供游戏ID，程序将回自动下载、更新、运行修改器。


## 2.使用方法
我已经为大家编译好了可执行文件，你可以直接使用，当然也能自行编译。
```bash
git clone https://github.com/SpenserCai/LGS-Helper.git
cd LGS-Helper/release/scripts/FLiNGLauncher
chmod +x ./FLiNGLauncher
./FLiNGLauncher -appid Steam游戏ID
```
如果需要强制重新下载修改器，可以使用`-redown`参数。
```bash
./FLiNGLauncher -appid Steam游戏ID -redown
```

## 3.兼容性
我测试了怪物猎人崛起和怪物猎人世界，都能正常下载、更新、运行修改器，如果发现有修改器无法正常运行，请提issue。

## 4.特别鸣谢
  <a href="https://flingtrainer.com/" title="风灵月影"><img src="https://flingtrainer.com/cn/community/data/avatars/l/0/1.jpg?1584477493" width=70px /></a>