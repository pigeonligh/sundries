#!/bin/sh

set -e

function install_tmux_entrance {
    go install ./cmd/tmux-entrance

    cp scripts/tmux.sh $HOME/.pigeonligh-sundries/tmux.sh
}

function install {
    mkdir -p $HOME/.pigeonligh-sundries
    cp scripts/source.sh $HOME/.pigeonligh-sundries/source
    
    install_tmux_entrance

    echo "write the following command into your .zshrc and re-source:"
    echo ""
    echo '    source $HOME/.pigeonligh-sundries/source'
    echo ""
}

install
