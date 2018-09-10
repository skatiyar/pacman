package pacman

import (
	"bytes"
	"fmt"
	"image/color"
	"image/png"
	"strconv"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/skatiyar/pacman/assets/fonts"
	"github.com/skatiyar/pacman/assets/images"
)

const MaxScoreView = 999999999
const MaxLifes = 7

func LoadSkin() (*ebiten.Image, error) {
	sImage, sImageErr := png.Decode(bytes.NewReader(images.SkinPng))
	if sImageErr != nil {
		return nil, sImageErr
	}

	skin, skinErr := ebiten.NewImageFromImage(sImage, ebiten.FilterDefault)
	if skinErr != nil {
		return nil, skinErr
	}

	return skin, nil
}

func LoadArcadeFont() (*truetype.Font, error) {
	return truetype.Parse(fonts.ArcadeTTF)
}

func SkinView(
	skin *ebiten.Image,
	powers *Powers,
	arcadeFont *truetype.Font,
) (func(state gameState, data *Data) (*ebiten.Image, error), error) {
	fontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    28,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	smallFontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	width, height := skin.Size()
	view, viewErr := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if viewErr != nil {
		return nil, viewErr
	}

	life, lifeErr := ScaleSprite(powers.Life, 0.5, 0.5)
	if lifeErr != nil {
		return nil, lifeErr
	}

	return func(state gameState, data *Data) (*ebiten.Image, error) {
		if clearErr := view.Clear(); clearErr != nil {
			return nil, clearErr
		}
		if drawErr := view.DrawImage(skin, &ebiten.DrawImageOptions{}); drawErr != nil {
			return nil, drawErr
		}

		text.Draw(
			view,
			fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()),
			smallFontface,
			24, 18, color.White)

		switch state {
		case GameStart:
			fallthrough
		case GamePause:
			fallthrough
		case GameOver:
			if data != nil {
				score := data.score
				lifes := data.lifes

				if score > MaxScoreView {
					score = MaxScoreView
				}
				numstr := strconv.Itoa(score)
				text.Draw(view, numstr, fontface, 682-(len(numstr)*27), 64, color.White)

				if lifes > MaxLifes {
					lifes = MaxLifes
				}

				ops := &ebiten.DrawImageOptions{}
				width, _ := life.Size()
				for i := 0; i < lifes; i++ {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(650-(width*i)), 80)
					if drawErr := view.DrawImage(life, ops); drawErr != nil {
						return nil, drawErr
					}
				}
			}
		}

		return view, nil
	}, nil
}
