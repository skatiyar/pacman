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
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/skatiyar/pacman/assets/fonts"
	"github.com/skatiyar/pacman/assets/images"
)

const MaxScoreView = 999999999
const MaxLifesView = 5

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
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	width, height := skin.Size()
	view, viewErr := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if viewErr != nil {
		return nil, viewErr
	}

	return func(state gameState, data *Data) (*ebiten.Image, error) {
		if clearErr := view.Clear(); clearErr != nil {
			return nil, clearErr
		}
		if drawErr := view.DrawImage(skin, &ebiten.DrawImageOptions{}); drawErr != nil {
			return nil, drawErr
		}
		if fpsErr := ebitenutil.DebugPrint(view,
			fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS())); fpsErr != nil {
			return nil, fpsErr
		}

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
				text.Draw(view, numstr, fontface, 344-(len(numstr)*15), 35, color.White)

				if lifes > MaxLifesView {
					lifes = MaxLifesView
				}
				ops := &ebiten.DrawImageOptions{}
				for i := 0; i < lifes; i++ {
					ops.GeoM.Reset()
					ops.GeoM.Scale(0.2, 0.2)
					ops.GeoM.Translate(float64(330-(16*i)), 42)
					if drawErr := view.DrawImage(powers.Life, ops); drawErr != nil {
						return nil, drawErr
					}
				}
			}
		}

		return view, nil
	}, nil
}
