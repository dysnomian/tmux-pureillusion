#!/usr/bin/env bash

BATTERY="$TMUX_PLUGIN_MANAGER_PATH/tmux-battery/tmux-battery.tmux"
CPU="$TMUX_PLUGIN_MANAGER_PATH/tmux-cpu/tmux-cpu.tmux"

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

$CURRENT_DIR/tmux-pureillusion

$BATTERY
$CPU
