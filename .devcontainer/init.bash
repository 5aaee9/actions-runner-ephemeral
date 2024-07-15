#!/usr/bin/env bash

set -xeuo pipefail

sudo apt install -y acl && sudo setfacl -k /tmp
nix profile install nixpkgs#direnv nixpkgs#nix-direnv
mkdir -p ~/.config/direnv

echo 'source $HOME/.nix-profile/share/nix-direnv/direnvrc' >> ~/.config/direnv/direnvrc
echo 'eval \"$(direnv hook bash)\"' >> ~/.bashrc
echo 'eval \"$(direnv hook zsh)\"' >> ~/.zshrc

direnv allow
# nix shell nixpkgs#acl --command sudo setfacl -k /tmp
nix print-dev-env > /dev/null