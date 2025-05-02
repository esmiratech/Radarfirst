#!/bin/sh

unicode_box() {
  local message="$1"
  local length=${#message}
  local border_top="╔═"
  local border_bottom="╚═"
  local fill_line="║ "

  for (( i=0; i<length; i++ )); do
    border_top+="═"
    border_bottom+="═"
  done

  border_top+="═╗"
  border_bottom+="═╝"

  echo "$border_top"
  echo "$fill_line$message ║"
  echo "$border_bottom"
}

pause() {
    echo ""
    echo "Press any key to continue..."
    read -n1 -s
    clear
}