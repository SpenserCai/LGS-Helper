<!--
 * @Author: SpenserCai
 * @Date: 2023-01-30 23:51:56
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-02-02 13:44:59
 * @Description: file content
-->
<div align="center">

# EaUnLockV2 (EA DLC UnLocker V2 - Linux)
<img src="https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white" />
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" /> 
<img src="https://img.shields.io/badge/Steam-000000?style=for-the-badge&logo=steam&logoColor=white" />
<img src="https://img.shields.io/badge/EA%20Game-CA4245?style=for-the-badge&logo=ea&logoColor=white" />

<!-- prettier-ignore-start -->
<!-- markdownlint-disable-next-line MD036 -->
Developed based on the original <a  href="https://sims.tarac.nl/the-sims-4/the-sims-4-free-downloads/add-pirated-dlcs-to-your-legit-the-sims-4-game/">EA DLC UnLocker V2</a>, it realizes the function of unlocking EA game DLC on Linux, and will continue to optimize and integrate it later.
The original technology comes from <a href="https://github.com/acidicoala/Koalageddon">Koalageddon</a>
<!-- prettier-ignore-end -->

</div>

<p align="center">
  <a href="./README.md">中文</a>
</p>

## Support
  |                         Game                       | Status |                 Tested OS                               |
  | :------------------------------------------------: | :--: | :-------------------------------------------------------: |
  |                    The Sims 4                      |  ✅  |                  Ubuntu 22.04,SteamOS                     |

## 1.Install The Sims 4
Get it for free from Steam

## 2.Install GE-Proton
https://github.com/GloriousEggroll/proton-ge-custom


## 3.First Run The Sims 4
Run The Sim4 once and give the villain a home, then save and exit

## 4.UnLock DLC
```bash
git clone https://github.com/SpenserCai/LGS-Helper.git
cd LGS-Helper/release/scripts/EA
chmod +x EAUnLockV2
./EAUnLockV2
```

## 5.Config Wine
~~After the program runs, the wine configuration window will pop up, add the version in the function library, and make sure that the following table shows that the original version is prior to the built-in, and the application is saved.~~

nothing to do, auto config

## 6.Run The Sims 4
If there is no purchase button, the unlock is successful,then exit the game.

## 7.Install DLC
From <a href="https://sims.tarac.nl/the-sims-4/the-sims-4-free-downloads/add-pirated-dlcs-to-your-legit-the-sims-4-game/">EA DLC UnLocker V2</a> Download the DLC image (find "All DLC's in 1 iso download", download the seed file), unzip or mount the image, and copy all the folders starting with EP SP FP GP to In the installation directory of The Sims 4, the installation can be completed.
