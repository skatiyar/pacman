# Pacman

[![Go to game](https://img.shields.io/badge/Try%20Game-Online-green.svg)](https://skatiyar.github.io/pacman)
[![Go Report Card](https://goreportcard.com/badge/github.com/skatiyar/pacman)](https://goreportcard.com/report/github.com/skatiyar/pacman)

Classic Pacman with procedurally generated infinite vertical maze.

![Sample](https://raw.githubusercontent.com/skatiyar/pacman/master/pacman.gif)

## Motivation

I came across Eller's algorithm for maze generation, a few months back. Eller's algorithm creates a perfect maze, by generating next row, on basis of current row. Giving us ability to create maze with infinite rows.

Since then I have been toying with idea of creating a game around it. It wasn't until a few days ago that I finally decided to use Pacman as the basis for game. I had experimented with [Ebiten](https://github.com/hajimehoshi/ebiten) 2D game engine a bit and this gave me a good opportunity to use it. For maze generation I slightly modified Eller's algorithm to create non-perfect mazes.

## Build

Using `go get` & without go modules.

```shell
$ go get -u github.com/skatiyar/pacman
$ cd skatiyar/pacman
$ go get ./...
$ cd build/pacman #goto build dir
$ go build -o pacman main.go
$ ./pacman
```

Using `git clone` & go modules.

```shell
$ git clone https://github.com/skatiyar/pacman.git
$ cd pacman/build/pacman #goto build dir
$ go build -o pacman main.go
$ ./pacman
```

## Build gh-pages

Golang code is converted to JS by using [gopherjs](https://github.com/gopherjs/gopherjs). Ebiten supports browsers by using webgl.

Note: Setup repo beforehand, as shown above.

To build just go code.

```shell
$ go get -u github.com/gopherjs/gopherjs
$ cd pacman/build/pacman
$ gopherjs build --tags=pacman --output=pacman.js
```

To build gh-pages.

```shell
$ go get -u github.com/gopherjs/gopherjs
$ cd pacman/build/pacman-pages
$ yarn install && yarn build
```

## How to play

- Use `arrow keys` to move pacman.
- Gain points by eating `dots`.
- Ghosts try to chase player and on collision player `looses` a life.
- Player starts with 5 lives and can have upto 7.
- Collect `diamond` to increase lives.
- Use `flask` to gain ability to destroy ghosts, ability lasts for `10 Sec` & ghosts try to runaway from player.
- `Eating` a running away ghost gives a bonus of `200 points`.

## Thanks to

- [Hajimehoshi](https://github.com/hajimehoshi) for Ebiten 2D engine.
- [Classic Gaming](http://www.classicgaming.cc) for Pacman & Ghost character-art & sounds.
- [Golang](https://golang.org) community for awesome tools and libraries.
