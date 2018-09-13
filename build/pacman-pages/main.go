// +build pacman jsgo

package main

import (
	"github.com/skatiyar/pacman"
)

func main() {
	game, gameErr := pacman.NewGame()
	if gameErr != nil {
		panic(gameErr)
	}

	if runErr := game.Run(); runErr != nil {
		panic(runErr)
	}
}
