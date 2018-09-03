package pacman

import (
	"bytes"
	"image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/skatiyar/pacman/assets/images"
)

type Powers struct {
	Life          *ebiten.Image
	Invincibility *ebiten.Image
}

func LoadPowers() (*Powers, error) {
	pImage, pImageErr := png.Decode(bytes.NewReader(images.PowersPng))
	if pImageErr != nil {
		return nil, pImageErr
	}

	powers, powersErr := ebiten.NewImageFromImage(pImage, ebiten.FilterDefault)
	if powersErr != nil {
		return nil, powersErr
	}

	life, lifeErr := GetSprite(64, 64, 0, 0, powers)
	if lifeErr != nil {
		return nil, lifeErr
	}

	invinc, invincErr := GetSprite(64, 64, 67, 0, powers)
	if invincErr != nil {
		return nil, invincErr
	}

	return &Powers{
		Life:          life,
		Invincibility: invinc,
	}, nil
}
