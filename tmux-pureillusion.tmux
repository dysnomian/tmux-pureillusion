#!/usr/bin/env bash

BATTERY_DIR="$TMUX_PLUGIN_MANAGER_PATH/tmux-battery"
CPU_DIR="$TMUX_PLUGIN_MANAGER_PATH/tmux-cpu"

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

$CURRENT_DIR/tmux-pureillusion

