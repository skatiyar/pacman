// Package spritetools provides functions to
// manipulate ebiten images.
package spritetools

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// GetSprite returns a new image from source,
// of size width x height. Starting point of image is
// specified by the xoffset & yoffset.
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

// ScaleSprite returns a new image from source,
// scaled by the given x & y.
func ScaleSprite(src *ebiten.Image, x, y float64) (*ebiten.Image, error) {
	spriteW, spriteH := src.Size()
	sSprite, sSpriteErr := ebiten.NewImage(
		int(float64(spriteW)*x),
		int(float64(spriteH)*y),
		ebiten.FilterDefault)
	if sSpriteErr != nil {
		return nil, sSpriteErr
	}

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(x, y)
	if drawErr := sSprite.DrawImage(src, ops); drawErr != nil {
		return nil, drawErr
	}

	return sSprite, nil
}
