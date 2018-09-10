#!/bin/bash
set -e

cd build/pacman-pages
yarn build
cd -
git subtree push --prefix build/pacman-pages/dist origin gh-pages
