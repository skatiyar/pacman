package pacman

import (
	"bytes"
	"image"
	"image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/skatiyar/pacman/assets/images"
)

type Characters struct {
	Pacman *ebiten.Image
	Ghost1 *ebiten.Image
	Ghost2 *ebiten.Image
	Ghost3 *ebiten.Image
	Ghost4 *ebiten.Image
}

func LoadCharacters() (*Characters, error) {
	cImage, cImageErr := png.Decode(bytes.NewReader(images.CharactersPng))
	if cImageErr != nil {
		return nil, cImageErr
	}

	characters, charactersErr := ebiten.NewImageFromImage(cImage, ebiten.FilterDefault)
	if charactersErr != nil {
		return nil, charactersErr
	}

	pacman, pacmanErr := GetSprite(61, 64, 0, 0, characters)
	if pacmanErr != nil {
		return nil, pacmanErr
	}

	ghost1, ghost1Err := GetSprite(56, 64, 66, 0, characters)
	if ghost1Err != nil {
		return nil, ghost1Err
	}

	ghost2, ghost2Err := GetSprite(56, 64, 125, 0, characters)
	if ghost2Err != nil {
		return nil, ghost2Err
	}

	ghost3, ghost3Err := GetSprite(56, 64, 185, 0, characters)
	if ghost3Err != nil {
		return nil, ghost3Err
	}

	ghost4, ghost4Err := GetSprite(56, 64, 244, 0, characters)
	if ghost4Err != nil {
		return nil, ghost4Err
	}

	return &Characters{
		Pacman: pacman,
		Ghost1: ghost1,
		Ghost2: ghost2,
		Ghost3: ghost3,
		Ghost4: ghost4,
	}, nil
}

// Offset is from top left corner
func GetSprite(
	width, height int,
	xoffset, yoffset int,
	src *ebiten.Image,
) (*ebiten.Image, error) {
	sprite, spriteErr := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if spriteErr != nil {
		return nil, spriteErr
	}

	rect := image.Rect(xoffset, yoffset, xoffset+width, yoffset+height)

	ops := &ebiten.DrawImageOptions{}
	ops.SourceRect = &rect
	if drawErr := sprite.DrawImage(src, ops); drawErr != nil {
		return nil, drawErr
	}

	return sprite, nil
}
