EA DLC Unlocker v2 - DLC unlocker for Origin and EA Desktop
made by anadius

Website: https://anadius.su/
Discord: https://anadius.su/discord
CS RIN thread: https://cs.rin.ru/forum/viewtopic.php?f=20&t=104412
Donate: https://anadius.su/donate

==========================
Installation instructions:
==========================
1. Run "setup.bat". (Special note for Sims players since I'm tired of hearing
   "but how do I run the setup.bat file?" - the same way you run any program
   or game on your PC - by double clicking on it.)
2. Select "Install" to install the Unlocker. The script automatically detects
   if you have EA Desktop or Origin installed. If you have both it will detect
   EA Desktop only and game started from Origin won't have the DLCs unlocked!
   If this happens reinstall Origin and run the script again or use the manual
   installation instructions below.
3. Select "Add/Update game config" and select the game.
4. Download DLC files if needed - links are on my website and in CS RIN thread.
   (Special note for Sims players - YES, you need the DLC files. DLC Unlocker
   doesn't download anything.)

If your DLCs suddenly stop working - it's because EA app updated and removed the
DLC Unlocker. So simply install it again.

On Linux the only difference is that you run "setup_linux.sh" (double click on
it - if that doesn't work run it from the terminal) and if you have multiple
wine/proton prefixes with EA app installed - you select which one to use.
Then you get the same menu as on Windows, so follow the instructions above
from the second step forwrads.

============================
Uninstallation instructions:
============================
1. Run "setup.bat".
2. Select "Uninstall".

=====================
Updates and new DLCs:
=====================
When a game gets a new DLC check if I added it to the Unlocker. There's a
changelog in the second post of CS RIN thread. If it's a DLC for The Sims 4
then the game config should automatically update when you restart
Origin/EA Desktop. If it's a DLC for a different game - redownload
the Unlocker and add the game config again.
You can enable/disable automatic config updates in "config.ini" file.

No matter what game it is for - check if you need to download the DLC files.
For The Sims 4 each Pack and Kit requires DLC files.

When a game gets a new update - just update it. No need to touch the Unlocker.
When Origin gets a new update - just update it. No need to touch the Unlocker.
When EA Desktop gets a new update - update it and install the Unlocker again.

=================================
Manual installation instructions:
=================================
1. Disable Origin/EA Desktop from autostart and reboot your PC.
   This will make sure it's not running and isn't messing with files.
2. Open the folder with Origin/EA Desktop.
     - if you're on Windows - right click on the shortcut
       and select "open file location"
     - if you're using Wine or if you deleted the shortcut
       the default install locations are:
         * Origin: "C:\Program Files (x86)\Origin"
         * EA Desktop: "C:\Program Files\Electronic Arts\EA Desktop\EA Desktop"
3. Copy the correct "version.dll" to the folder you opened in the previous step.
   If you use Origin copy it from "origin" folder.
   If you use EA Desktop copy it from "ea_desktop" folder.
4. Open "C:\Users\<your username>\AppData\Roaming", create "anadius" folder,
   open it, create "EA DLC Unlocker v2" folder, open it. Full path should be
   "C:\Users\<your username>\AppData\Roaming\anadius\EA DLC Unlocker v2"
5. Copy "config.ini" and any game config you want to the folder opened
   in the previous step.
6. (EA Desktop only) Open command prompt as administrator and type:

   schtasks /Create /F /RL HIGHEST /SC ONCE /ST 00:00 /SD 01/01/2000 /TN copy_dlc_unlocker /TR "xcopy.exe /Y 'C:\Program Files\Electronic Arts\EA Desktop\EA Desktop\version.dll' 'C:\Program Files\Electronic Arts\EA Desktop\StagedEADesktop\EA Desktop\*'"

   If you get some error message change 01/01/2000 to 2000/01/01

And that's it, you have Unlocker v2 installed. If you want to uninstall it 
just delete that "version.dll" file and then delete
"C:\Users\<your username>\AppData\Roaming\anadius\EA DLC Unlocker v2" and
"C:\Users\<your username>\AppData\Local\anadius\EA DLC Unlocker v2" folders.
If you followed the 6th step above open command prompt as administrator and run:
schtasks /Delete /TN copy_dlc_unlocker /F
