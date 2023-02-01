EA DLC Unlocker v2 - DLC unlocker for Origin and EA Desktop
made by anadius

Website: https://anadius.hermietkreeft.site/
Discord: https://anadius.hermietkreeft.site/discord
CS RIN thread: https://cs.rin.ru/forum/viewtopic.php?f=20&t=104412
Donate: https://anadius.hermietkreeft.site/donate

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
4. Download DLC files if needed - see the links in the CS RIN thread.

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
         * EA Desktop: "C:\Program Files\Electronic Arts\EA Desktop"
3. Copy the correct "version.dll" to the folder you opened in the previous step.
   If you use Origin copy it from "origin" folder.
   If you use EA Desktop copy it from "ea_desktop" folder.
4. Open "C:\Users\<your username>\AppData\Roaming", create "anadius" folder,
   open it, create "EA DLC Unlocker v2" folder, open it. Full path should be
   "C:\Users\<your username>\AppData\Roaming\anadius\EA DLC Unlocker v2"
5. Copy "config.ini" and any game config you want to the folder opened
   in the previous step.

And that's it, you have Unlocker v2 installed. If you want to uninstall it 
just delete that "version.dll" file and then delete
"C:\Users\<your username>\AppData\Roaming\anadius\EA DLC Unlocker v2" and
"C:\Users\<your username>\AppData\Local\anadius\EA DLC Unlocker v2" folders.

===============================
Additional info for Wine users:
===============================
I doubt you can run the setup script through Wine so use the manual
installation instructions above. Then try if the Unlocker works
(start the game, test if DLCs are unlocked). If it does - great!
If it doesn't follow these instructions:
1. Open "winecfg".
2. Open "Libraries" tab.
3. Select "version" and click "Add".
4. Select it on "Existing overrides" list and click "Edit".
5. Select "Native then Builtin" and press "OK".
6. Press OK.
Now the Unlocker should work.
