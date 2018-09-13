package pacman

import (
	"image/color"
	"strconv"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/skatiyar/pacman/assets"
	"github.com/skatiyar/pacman/spritetools"
)

const MaxScoreView = 999999999
const MaxLifes = 7

func SkinView(
	skin *ebiten.Image,
	powers *assets.Powers,
	arcadeFont *truetype.Font,
) (func(state gameState, data *Data) (*ebiten.Image, error), error) {
	fontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    28,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	width, height := skin.Size()
	view, viewErr := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if viewErr != nil {
		return nil, viewErr
	}

	life, lifeErr := spritetools.ScaleSprite(powers.Life, 0.5, 0.5)
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
