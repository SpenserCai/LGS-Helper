<# : batch portion
@echo off & setlocal

pushd "%~dp0"

set "p1=%~f0"
set p=%p1:^=%
set p=%p:@=%
set p=%p:&=%
if not "%p1%"=="%p%" goto :badpath

if not "%~nx0"=="setup.bat" goto :badname

echo marco | findstr /C:"polo" >nul
if %ERRORLEVEL% EQU 0 goto :wineskip
echo marco | findstr /V /C:"polo" >nul
if %ERRORLEVEL% NEQ 0 goto :wineskip
echo. >nul || goto :wineskip

echo "%~dp0" | findstr /V /C:"%TEMP%" >nul
if %ERRORLEVEL% NEQ 0 goto :temp

set "script_path=%~f0"
set "arg1=%1"

echo "Starting the script... %~f0"
powershell -noprofile "$_PSCommandPath = [Environment]::GetEnvironmentVariable('script_path', 'Process'); iex ((Get-Content -LiteralPath $_PSCommandPath) | out-string)"
if %ERRORLEVEL% EQU 0 goto :EOF

if %ERRORLEVEL% LSS 0 exit /B %ERRORLEVEL%

pause
goto :EOF

:wineskip
echo It looks like you're trying to run this script through Wine - it won't work without PowerShell installed. Follow the manual installation instructions in CS RIN thread - see the readme file. And if you want to try with this script anyway - remove the lines that end with "goto :wineskip"
pause
goto :EOF

:temp
echo It looks like you're trying to run this script from inside the archive. Make sure you extract the file first.
pause
goto :EOF

:badname
echo Don't rename this script, leave it as "setup.bat"!
pause
goto :EOF

:badpath
echo %~dp0
echo You put the Unlocker in a path that will break the setup script. Move it somewhere else, for example "C:\unlocker" or "D:\unlocker". The problematic characters are: @^&^^
pause
goto :EOF
: end batch / begin powershell #>

function Get-Env {
    param (
        [string]$name
    )

    Return [Environment]::GetEnvironmentVariable($name, 'Process')
}

$arg1 = Get-Env 'arg1'

$ErrorActionPreference = 'stop'
Set-Location -LiteralPath (Split-Path -parent $_PSCommandPath)
$FileMissingMessage = ' missing, you didn''t extract all files'
$GameConfigPrefix = 'g_'
$GameConfigSuffix = '.ini'
Clear-Host

function Fail {
    param (
        [string]$message
    )

    Write-Host `n'Fatal error:' -NoNewline -BackgroundColor red -ForegroundColor white
    Warn (' ' + $message)
    Write-Host 'Script path:' $_PSCommandPath
    Write-Host
    Exit 1
}

function Warn {
    param (
        [string]$message
    )

    Write-Host $message -ForegroundColor red
}

function Success {
    param (
        [string]$message
    )

    Write-Host $message -ForegroundColor green
}

function Special {
    param (
        [string]$yellow,
        [string]$suffix
    )

    Write-Host $yellow -NoNewline -ForegroundColor yellow
    Write-Host $suffix
}

function Special2 {
    param (
        [string]$red,
        [string]$suffix
    )

    Write-Host $red -NoNewline -ForegroundColor red
    Write-Host $suffix
}

function Force-Stop-Clients {
    If ($client -Eq 'origin') {
        $wildcard = 'Origin*'
    }
    Else {
        $wildcard = 'EA*'
    }
    Stop-Process -Force -Name $wildcard
    Wait-Process -Name $wildcard -Timeout 10
}

function Delete-Folder-Recursively {
    param (
        [string]$directory
    )

    If (Test-Path -LiteralPath $directory) {
        Get-ChildItem -LiteralPath $directory -Force -Recurse | Remove-Item -Force
        Remove-Item -LiteralPath $directory -Force
    }
}

function Delete-Folder-If-Empty {
    param (
        [string]$directory
    )

    If ((Test-Path -LiteralPath $directory) -And ((Get-ChildItem -LiteralPath $directory).Count -Eq 0)) {
        Remove-Item -LiteralPath $directory -Force
    }
}

function Get-Client-Path-From-Registry {
    param (
        [string]$RegistryPath
    )

    $path = (Get-ItemProperty -Path ('Registry::HKEY_LOCAL_MACHINE\SOFTWARE\' + $RegistryPath) -Name ClientPath).ClientPath
    Return (Resolve-Path -LiteralPath (Join-Path $path '..'))
}

function Is-Admin {
    Return ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
}

function Delete-If-Exists {
    param (
        [string]$path
    )

    If (Test-Path -LiteralPath $path) {
        Remove-Item -LiteralPath $path -Force
    }
}

function Remove-Old-Unlocker {
    Delete-If-Exists (Join-Path $client_path 'version_o.dll')
    Delete-If-Exists (Join-Path $client_path 'winhttp.dll')
    Delete-If-Exists (Join-Path $client_path 'winhttp_o.dll')
    Get-ChildItem -LiteralPath $client_path -Force | Where-Object {($_.Name.EndsWith('.ini') -And $_.Name.StartsWith('w_'))} | Remove-Item -Force
}

function Common-Setup-Real {
    param (
        [string]$action
    )

    Try {
        Force-Stop-Clients
        Remove-Old-Unlocker
        If ($action -Eq 'install') {
            Copy-Item $srcdll -Destination $dstdll -Force
            Return 0
        }
        ElseIf ($action -Eq 'uninstall') {
            Delete-If-Exists $dstdll
            Return 0
        }
        Else {
            Return -1
        }
    }
    Catch {
        If (Is-Admin) {
            Return -1
        }
        Else {
            Special 'Requesting administrator rights...'
            Try {
                $process = Start-Process -FilePath setup.bat -Verb RunAs -WorkingDirectory . -ArgumentList $action -PassThru -Wait
            }
            Catch {
                Fail 'Failed to get administrator rights. Run this script as administrator yourself. Right click on it and then "Run as administrator".'
            }
            Return $process.ExitCode
        }
    }
}

function Common-Setup {
    param (
        [string]$action
    )

    $result = Common-Setup-Real $action
    If ($result -Ne 0) {
        Fail ('Permission error. Could not ' + $action + ' the Unlocker.')
    }
}

function Create-Config-Directory {
    Try {
        New-Item -Path $appdatadir -ItemType 'Directory' -Force | Out-Null
    }
    Catch {
        Fail 'Could not create the configs folder.'
    }
}

function Install-Unlocker {
    Write-Host 'Installing...'

    If (-Not (Test-Path $srcdll)) {
        Fail ($srcdll + $FileMissingMessage + ' or your anti-virus deleted it.')
    }
    If (-Not (Test-Path $srcconfig)) {
        Fail ($srcconfig + $FileMissingMessage + '.')
    }

    Create-Config-Directory
    Success 'Configs folder created!'
    Try {
        Copy-Item $srcconfig -Destination $dstconfig -Force
        Success 'Main config copied!'
    }
    Catch {
        Fail 'Could not copy the main config.'
    }

    Common-Setup 'install'
    Success 'DLC Unlocker installed!'
}

function Uninstall-Unlocker {
    Write-Host 'Uninstalling...'

    Try {
        Delete-Folder-Recursively $appdatadir
        Delete-Folder-If-Empty (Join-Path $appdatadir '..')
        Success 'Configs folder deleted!'
    }
    Catch {
        Warn 'Could not delete the configs folder.'
    }
    Try {
        Delete-Folder-Recursively $localappdatadir
        Delete-Folder-If-Empty (Join-Path $localappdatadir '..')
        Success 'Logs folder deleted!'
    }
    Catch {
        Warn 'Could not delete the logs folder.'
    }

    Common-Setup 'uninstall'
    Success 'DLC Unlocker uninstalled!'
}

function Open-Configs-Folder {
    If (Test-Path -LiteralPath $appdatadir) {
        Invoke-Item -LiteralPath $appdatadir
        Success 'Configs folder opened!'
    }
    Else {
        Warn 'Configs folder not found. Install the Unlocker first.'
    }
}

function Open-Logs-Folder {
    If (Test-Path -LiteralPath $localappdatadir) {
        Invoke-Item -LiteralPath $localappdatadir
        Success 'Logs folder opened!'
    }
    Else {
        Warn 'Logs folder not found. Install the Unlocker and run EA Desktop/Origin first.'
    }
}

function Add-Game-Config {
    Try {
        [string[]] $configs = Get-ChildItem -Path '.' | Where-Object {($_.Name.EndsWith($GameConfigSuffix) -And $_.Name.StartsWith($GameConfigPrefix))} | %{ $_.Name.Substring(2, ($_.Name.Length-6)) }
    }
    Catch {
        $configs = @()
    }

    If ($configs.Length -Eq 0) {
        Fail ('Game configs' + $FileMissingMessage + '.')
    }
    Else {
        While ($True) {
            Special 'Game configs' ':'
            For ($i = 0; $i -Lt $configs.Length; $i++) {
                Special ($i + 1) ('. ' + $configs[$i])
            }
            Special2 'b' '. Go back'

            $choice = Read-Host -Prompt `n'Choose option number'
            Clear-Host

            If ($choice -Eq 'b') {
                Write-Host 'No game config selected.'
                Return
            }
            Try {
                $game = $configs.Get(([int] $choice) - 1)
                Break
            }
            Catch {}
            Warn 'Bad option!'
            Write-Host
        }
    }

    Create-Config-Directory
    Special $game ' config selected.'
    Try {
        Copy-Item ($GameConfigPrefix + $game + $GameConfigSuffix) -Destination $appdatadir -Force
        Success 'Game config copied!'
    }
    Catch {
        Fail 'Could not copy the game config.'
    }

    Try {
        Delete-If-Exists (Join-Path $localappdatadir ($game + '.etag'))
    }
    Catch {}
}

$client = 'ea_desktop'
$client_name = 'EA Desktop'
Try {
    $client_path = Get-Client-Path-From-Registry 'Electronic Arts\EA Desktop'
}
Catch {
    $client = 'origin'
    $client_name = 'Origin'
    Try {
        $client_path = Get-Client-Path-From-Registry 'WOW6432Node\Origin'
    }
    Catch {
        Try {
            $client_path = Get-Client-Path-From-Registry 'Origin'
        }
        Catch {
            Fail 'EA Desktop/Origin not found, reinstall one of them.'
        }
    }
}

$srcdll = Join-Path $client 'version.dll'
$dstdll = Join-Path $client_path 'version.dll'

$commondir = 'anadius\EA DLC Unlocker v2'
$appdatadir = Join-Path (Get-Env 'AppData') $commondir
$localappdatadir = Join-Path (Get-Env 'LocalAppData') $commondir

$srcconfig = 'config.ini'
$dstconfig = Join-Path $appdatadir 'config.ini'

If (($arg1 -Eq 'install') -Or ($arg1 -Eq 'uninstall')) {
    Exit (Common-Setup-Real $arg1)
}

While ($True) {
    Special $client_name ' detected'
    Special '1' '. Install EA DLC Unlocker'
    Special '2' '. Add/Update game config'
    Special '3' '. Open folder with installed configs'
    Special '4' '. Open folder with log file'
    Special '5' '. Uninstall EA DLC Unlocker'
    Special2 'q' '. Quit'

    $choice = Read-Host -Prompt `n'Choose option number'
    Clear-Host

    If     ($choice -Eq '1') { Install-Unlocker }
    ElseIf ($choice -Eq '2') { Add-Game-Config }
    ElseIf ($choice -Eq '3') { Open-Configs-Folder }
    ElseIf ($choice -Eq '4') { Open-Logs-Folder }
    ElseIf ($choice -Eq '5') { Uninstall-Unlocker }
    ElseIf ($choice -Eq 'q') { Exit 0 }
    Else { Warn 'Bad option!' }
    Write-Host
}
