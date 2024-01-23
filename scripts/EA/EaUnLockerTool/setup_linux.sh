#!/usr/bin/env bash

cd "$(dirname "$0")"

STEAM_PATHS=(
  "$HOME/.local/share/Steam"
  "$HOME/.steam/steam"
)

# There are other games supported by this Unlocker, but their
# Steam version doesn't require EA app.
KNOWN_IDS=(
  "1238060:Dead Space 3"
  "1426210:It Takes Two"
  "1222680:Need For Speed Heat"
  "1262560:Need For Speed Most Wanted"
  "1262580:Need For Speed Payback"
  "1846380:Need For Speed Unbound"
  "1172380:STAR WARS Jedi Fallen Order"
  "1774580:STAR WARS Jedi Survivor"
  "1222670:The Sims 4"
)

ea_app_parent="drive_c/Program Files/Electronic Arts/EA Desktop"
unlocker_dir="anadius/EA DLC Unlocker v2"

DLL_NAME="version.dll"
MAIN_CONFIG="config.ini"
EA_APP="$ea_app_parent/EA Desktop"
EA_APP_OLD="$ea_app_parent/StagedEADesktop/EA Desktop"
USERS_DIR="drive_c/users"
APPDATA_DIR="AppData/Roaming/$unlocker_dir"
LOCALAPPDATA_DIR="AppData/Local/$unlocker_dir"

unset ea_app_parent unlocker_dir

CLIENT="ea_desktop"
SRC_DLL="$CLIENT/$DLL_NAME"
SRC_CONFIG="$MAIN_CONFIG"

FILE_MISSING_MESSAGE="missing, you didn't extract all files."

######################################################################

function warn {
  printf "\e[31m$1\e[0m\n"
}

function fail {
  printf "\e[37;41mFatal error:\e[0m "
  warn "$1"
  echo ""
  read -s -p "Press enter to continue..."
  echo ""
  exit 1
}

function success {
  printf "\e[32m$1\e[0m\n"
}

function special {
  printf "\e[33m$1\e[0m$2\n"
}

function special2 {
  printf "\e[31m$1\e[0m$2\n"
}

function get_prefix_name {
  local appid steamapps game id name
  appid="$1"
  steamapps="$2"

  prefix_name="Unknown prefix ($appid)"
  prefix_config=""

  for game in "${KNOWN_IDS[@]}"; do
    id="${game%%:*}"
    if [ "$id" = "$appid" ]; then
      name="${game#*:}"
      prefix_name="$name"
      prefix_config="$name"
      return
    fi
  done

  name=$(grep -Po '"name"\s+"\K([^"]+)' "$steamapps/appmanifest_$appid.acf" 2>/dev/null)
  if [ ! -z "$name" ]; then
    prefix_name="$name"
    return
  fi
}

function check_prefix {
  local path name src
  path="$1"
  name="$2"
  src="$3"

  if [ ! -d "$path/$EA_APP" ]; then
    return
  fi

  if [ "$src" = "wine" ]; then
    config=""
    usr="$USER"
  elif [ "$src" = "steam" ]; then
    # this sets $prefix_name and $prefix_config
    get_prefix_name "$name" "$4"
    name="$prefix_name"
    config="$prefix_config"
    usr="steamuser"
  else
    return
  fi

  ALL_PREFIX_PATHS+=("$path")
  ALL_PREFIX_NAMES+=("$name")
  ALL_PREFIX_CONFIGS+=("$config")
  ALL_PREFIX_USERS+=("$usr")
}

function get_wine_prefix {
  local name
  if [ -z "$WINEPREFIX" ]; then
    WINEPREFIX="$HOME/.wine"
    name="Default Wine prefix"
  else
    name="Custom Wine prefix"
  fi

  check_prefix "$WINEPREFIX" "$name" "wine"
}

function get_steam_prefixes {
  local steamapps_path steam_path tmp library library_steamapps prefix
  for steam_path in ${STEAM_PATHS[@]}; do
    tmp="$steam_path/steamapps"
    if [ -d $tmp ]; then
      steamapps_path=$tmp
      break
    fi
  done

  if [ -z "$steamapps_path" ]; then
    return
  fi

  while read -r library ; do
    library_steamapps="$library/steamapps"
    for prefix in "$library_steamapps/compatdata"/*; do
      check_prefix "$prefix/pfx" "$(basename "$prefix")" "steam" "$library_steamapps"
    done
  done < <(grep -Po '"path"\s+"\K([^"]+)' "$steamapps_path/libraryfolders.vdf")
}

function show_prefix_menu {
  local tmp i choice

  declare -i tmp
  while true; do
    printf "Multiple wine/proton prefixes found!\n"
    printf "Which one do you want to manage?\n\n"

    for i in "${!ALL_PREFIX_PATHS[@]}"; do
      tmp=i+1
      special "$tmp" ". ${ALL_PREFIX_NAMES[$i]}"
    done
    special2 "q" ". Quit"

    read -p "Choose option number and press enter: " choice

    if [ "$choice" = "q" ]; then exit 0; fi
    clear
    
    case $choice in
      ''|*[!0-9]*) tmp=-1 ;;
      *) tmp=choice-1 ;;
    esac

    prefix_path="${ALL_PREFIX_PATHS[$tmp]}"
    if [ $tmp -lt 0 ] || [ -z "$prefix_path" ]; then
      warn "Bad choice!\n"
    else
      prefix_name="${ALL_PREFIX_NAMES[$tmp]}"
      prefix_config="${ALL_PREFIX_CONFIGS[$tmp]}"
      prefix_user="${ALL_PREFIX_USERS[$tmp]}"
      break
    fi
  done
}

function create_config_directory {
  if [ ! -d "$CONFIGS_DIR" ]; then
    mkdir -p "$CONFIGS_DIR" || fail "Could not create the configs folder."
    success "Configs folder created!"
  fi
}

function delete_if_exists {
  if [ -f "$1" ]; then
    rm -f "$1" || fail "Could not delete file: $1"
  fi
}

function install_unlocker {
  echo "Installing..."

  if [ ! -f "$SRC_DLL" ]; then
    fail "$SRC_DLL $FILE_MISSING_MESSAGE"
  fi
  if [ ! -f "$SRC_CONFIG" ]; then
    fail "$SRC_CONFIG $FILE_MISSING_MESSAGE"
  fi

  create_config_directory
  cp -f "$SRC_CONFIG" "$DST_CONFIG" || fail "Could not copy the main config."
  success "Main config copied!"

  echo "" >> "$REG"
  echo "[Software\\\\Wine\\\\DllOverrides]" >> "$REG"
  echo "\"version\"=\"native,builtin\"" >> "$REG"

  cp -f "$SRC_DLL" "$DST_DLL" || fail "Could not install the Unlocker."
  cp -f "$SRC_DLL" "$DST_DLL2" 2>/dev/null
  success "DLC Unlocker installed!"
}

function uninstall_unlocker {
  echo "Uninstalling..."

  if [ -d "$CONFIGS_DIR" ]; then
    rm -r -f "$CONFIGS_DIR" || fail "Could not delete the configs folder."
  fi
  success "Configs folder deleted!"
  rmdir "$(dirname "$CONFIGS_DIR")" 2>/dev/null

  if [ -d "$LOGS_DIR" ]; then
    rm -r -f "$LOGS_DIR" || fail "Could not delete the logs folder."
  fi
  success "Logs folder deleted!"
  rmdir "$(dirname "$LOGS_DIR")" 2>/dev/null
  
  delete_if_exists "$DST_DLL"
  delete_if_exists "$DST_DLL2"
}

function open_configs_folder {
  if [ -d "$CONFIGS_DIR" ]; then
    xdg-open "$CONFIGS_DIR" && success "Configs folder opened!"
  else
    warn "Configs folder not found. Install the Unlocker first."
  fi
}

function open_logs_folder {
  if [ -d "$LOGS_DIR" ]; then
    xdg-open "$LOGS_DIR" && success "Logs folder opened!"
  else
    warn "Logs folder not found. Install the Unlocker and run the game first."
  fi

}

function show_configs_menu {
  local tmp i choice names
  names=("$@")

  declare -i tmp
  while true; do
    special "Game configs" ":"

    for i in "${!names[@]}"; do
      tmp=i+1
      special "$tmp" ". ${names[$i]}"
    done
    special2 "b" ". Back"

    read -p "Choose option number and press enter: " choice
    clear

    if [ "$choice" = "b" ]; then
      echo "No game config selected."
      return 255
    fi
    
    case $choice in
      ''|*[!0-9]*) tmp=-1 ;;
      *) tmp=choice-1 ;;
    esac

    chosen_game="${names[$tmp]}"
    if [ $tmp -lt 0 ] || [ -z "$chosen_game" ]; then
      warn "Bad choice!\n"
    else
      break
    fi
  done
}
function add_game_config {
  local names path game chosen_config
  if [ -z "$prefix_config" ]; then
    names=()
    for path in g_*.ini; do
      names+=("${path:2: -4}")
    done

    show_configs_menu "${names[@]}" || return

    game="$chosen_game"
  else
    game="$prefix_config"
  fi

  chosen_config="g_$game.ini"
  if [ ! -f "$chosen_config" ]; then
    fail "$chosen_config $FILE_MISSING_MESSAGE"
  fi

  special "$game" " config selected."

  create_config_directory
  cp -f "$chosen_config" "$CONFIGS_DIR" || fail "Could not copy the game config."
  success "Game config copied!"

  rm -f "$LOGS_DIR/$game.etag" 2>/dev/null
}

function show_main_menu {
  local choice

  #echo "path:$prefix_path; name:$prefix_name; config:$prefix_config;"

  while true; do
    echo "Prefix: $prefix_name ($prefix_path)"
    printf "DLC Unlocker "
    if [ -f "$DST_DLL" ] && [ -f "$DST_CONFIG" ]; then
      success "installed"
    else
      warn "not installed"
    fi

    special "1" ". Install EA DLC Unlocker"
    special "2" ". Add/Update game config"
    special "3" ". Open folder with installed configs"
    special "4" ". Open folder with log file"
    special "5" ". Uninstall EA DLC Unlocker"
    special2 "q" ". Quit"

    read -p "Choose option number and press enter: " choice
    if [ "$choice" = "q" ]; then exit 0; fi
    clear

    if [ "$choice" = "1" ]; then install_unlocker
    elif [ "$choice" = "2" ]; then add_game_config
    elif [ "$choice" = "3" ]; then open_configs_folder
    elif [ "$choice" = "4" ]; then open_logs_folder
    elif [ "$choice" = "5" ]; then uninstall_unlocker
    else warn "Bad option!"
    fi

    echo ""
  done
}

######################################################################

ALL_PREFIX_PATHS=()
ALL_PREFIX_NAMES=()
ALL_PREFIX_CONFIGS=()
ALL_PREFIX_USERS=()

get_wine_prefix
get_steam_prefixes

clear

if [ ${#ALL_PREFIX_PATHS[@]} -eq 0 ]; then
  fail "No prefixes found. Run the game once and then try again."
elif [ ${#ALL_PREFIX_PATHS[@]} -eq 1 ]; then
  prefix_path="${ALL_PREFIX_PATHS[0]}"
  prefix_name="${ALL_PREFIX_NAMES[0]}"
  prefix_config="${ALL_PREFIX_CONFIGS[0]}"
  prefix_user="${ALL_PREFIX_USERS[0]}"
else
  show_prefix_menu
fi

REG="$prefix_path/user.reg"
DST_DLL="$prefix_path/$EA_APP/$DLL_NAME"
DST_DLL2="$prefix_path/$EA_APP_OLD/$DLL_NAME"
CONFIGS_DIR="$prefix_path/$USERS_DIR/$prefix_user/$APPDATA_DIR"
LOGS_DIR="$prefix_path/$USERS_DIR/$prefix_user/$LOCALAPPDATA_DIR"
DST_CONFIG="$CONFIGS_DIR/$MAIN_CONFIG"

show_main_menu
