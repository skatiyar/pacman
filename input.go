package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func SpacePressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}
