package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func spacePressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func spaceReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace)
}

func upKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyUp)
}

func upKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyUp)
}

func downKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyDown)
}

func downKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyDown)
}

func leftKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyLeft)
}

func leftKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyLeft)
}

func rightKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyRight)
}

func rightKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyRight)
}
