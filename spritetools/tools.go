// Spritetools package provides functions to
// manipulate ebiten images.
package spritetools

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

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

func ScaleSprite(sprite *ebiten.Image, x, y float64) (*ebiten.Image, error) {
	spriteW, spriteH := sprite.Size()
	sSprite, sSpriteErr := ebiten.NewImage(
		int(float64(spriteW)*x),
		int(float64(spriteH)*y),
		ebiten.FilterDefault)
	if sSpriteErr != nil {
		return nil, sSpriteErr
	}

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(x, y)
	if drawErr := sSprite.DrawImage(sprite, ops); drawErr != nil {
		return nil, drawErr
	}

	return sSprite, nil
}
