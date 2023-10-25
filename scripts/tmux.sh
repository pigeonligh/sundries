function tmux {
    while true; do
        tmux-entrance
        if [ $(echo $?) != "0" ]; then
            break
        fi
    done
}
