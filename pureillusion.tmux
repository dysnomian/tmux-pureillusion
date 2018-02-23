#!/usr/bin/env bash

# listing installed plugins
ls -1 "$plugin_path"

$($TMUX_PLUGIN_MANAGER_PATH/tmux-pureillusion/tmux-pureillusion)
