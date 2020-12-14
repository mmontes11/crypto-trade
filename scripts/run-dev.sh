#!/bin/bash

set -e

docker-compose -f docker-compose.services.yml up -d

tmux new-session -d -s crypto
tab=0

function new_tab() {
    name="$1"
    path="$2"
    tab=$(($tab + 1))
    tmux new-window -t crypto:"$tab" -n "$name"
    tmux send-keys -t crypto:"$tab" "cd $path; make run" enter
}

for ms in $(ls -d cmd/*); do
    name=$(basename "$ms")
    path="$ms"
    new_tab "$name" "$path"
done

tmux rename-window -t crypto:0 'workspace'
tmux select-window -t crypto:0

tmux attach -t crypto
