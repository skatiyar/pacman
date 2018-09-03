package pacman

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/skatiyar/pacman/assets/images"
)

type Walls struct {
	ActiveCorner   *ebiten.Image
	ActiveSide     *ebiten.Image
	InActiveCorner *ebiten.Image
	InActiveSide   *ebiten.Image
}

func LoadWalls() (*Walls, error) {
	wImage, wImageErr := png.Decode(bytes.NewReader(images.WallsPng))
	if wImageErr != nil {
		return nil, wImageErr
	}

	walls, wallsErr := ebiten.NewImageFromImage(wImage, ebiten.FilterDefault)
	if wallsErr != nil {
		return nil, wallsErr
	}

	inactiveCorner, inactiveCornerErr := GetSprite(12, 12, 0, 0, walls)
	if inactiveCornerErr != nil {
		return nil, inactiveCornerErr
	}

	inactiveSide, inactiveSideErr := GetSprite(40, 12, 12, 0, walls)
	if inactiveSideErr != nil {
		return nil, inactiveSideErr
	}

	activeCorner, activeCornerErr := GetSprite(12, 12, 52, 0, walls)
	if activeCornerErr != nil {
		return nil, activeCornerErr
	}

	activeSide, activeSideErr := GetSprite(40, 12, 64, 0, walls)
	if activeSideErr != nil {
		return nil, activeSideErr
	}

	return &Walls{
		ActiveCorner:   activeCorner,
		ActiveSide:     activeSide,
		InActiveCorner: inactiveCorner,
		InActiveSide:   inactiveSide,
	}, nil
}
