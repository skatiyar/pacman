package pacman

import (
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

func GridView(
	characters *Characters,
	powers *Powers,
	arcadeFont *truetype.Font,
	mazeView func(state gameState, data *Data) (*ebiten.Image, error),
) (func(state gameState, data *Data) (*ebiten.Image, error), error) {
	fontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	pacman, pacmanErr := ScaleSprite(characters.Pacman, 0.27, 0.27)
	if pacmanErr != nil {
		return nil, pacmanErr
	}

	ghost1, ghost1Err := ScaleSprite(characters.Ghost1, 0.27, 0.27)
	if ghost1Err != nil {
		return nil, ghost1Err
	}

	ghost2, ghost2Err := ScaleSprite(characters.Ghost2, 0.27, 0.27)
	if ghost2Err != nil {
		return nil, ghost2Err
	}

	ghost3, ghost3Err := ScaleSprite(characters.Ghost3, 0.27, 0.27)
	if ghost3Err != nil {
		return nil, ghost3Err
	}

	ghost4, ghost4Err := ScaleSprite(characters.Ghost4, 0.27, 0.27)
	if ghost4Err != nil {
		return nil, ghost4Err
	}

	view, viewErr := ebiten.NewImage(32*Columns, 512, ebiten.FilterDefault)
	if viewErr != nil {
		return nil, viewErr
	}

	return func(state gameState, data *Data) (*ebiten.Image, error) {
		if clearErr := view.Clear(); clearErr != nil {
			return nil, clearErr
		}
		if fillErr := view.Fill(color.Black); fillErr != nil {
			return nil, fillErr
		}

		ops := &ebiten.DrawImageOptions{}
		switch state {
		case GameLoading:
			text.Draw(view, "PRESS SPACE", fontface, 75, 240, color.White)
			text.Draw(view, "TO BEGIN", fontface, 100, 270, color.White)
		case GameStart, GamePause:
			mazeView, mazeViewErr := mazeView(state, data)
			if mazeViewErr != nil {
				return nil, mazeViewErr
			}

			ops.GeoM.Reset()
			ops.GeoM.Translate(0, -(float64(len(data.grid)*32) - (512 + data.gridOffsetY)))
			if drawErr := view.DrawImage(mazeView, ops); drawErr != nil {
				return nil, drawErr
			}

			ops.GeoM.Reset()
			pwidth, pheight := pacman.Size()
			switch data.pacman.direction {
			case North:
				ops.GeoM.Rotate(-1.5708)
				ops.GeoM.Translate(
					data.pacman.posX-float64(pwidth/2),
					512-(data.pacman.posY-float64(pheight-(pheight/2))))
			case East:
				ops.GeoM.Translate(
					data.pacman.posX-float64(pwidth/2),
					512-(data.pacman.posY+float64(pheight/2)))
			case South:
				ops.GeoM.Rotate(1.5708)
				ops.GeoM.Translate(
					data.pacman.posX+float64(pwidth/2),
					512-(data.pacman.posY+float64(pheight/2)))
			case West:
				ops.GeoM.Rotate(3.14159)
				ops.GeoM.Translate(
					data.pacman.posX+float64(pwidth/2),
					512-(data.pacman.posY-float64(pheight-(pheight/2))))
			}
			if drawErr := view.DrawImage(pacman, ops); drawErr != nil {
				return nil, drawErr
			}

			for i := 0; i < len(data.ghosts); i++ {
				ghost := data.ghosts[i]
				ghostImg := ghost1
				switch ghost.kind {
				case Ghost2:
					ghostImg = ghost2
				case Ghost3:
					ghostImg = ghost3
				case Ghost4:
					ghostImg = ghost4
				}
				gwidth, gheight := ghostImg.Size()
				ops.GeoM.Reset()
				ops.GeoM.Translate(
					data.ghosts[i].posX-float64(gwidth/2),
					512-(data.ghosts[i].posY-float64(gheight+(gheight/2))))
				if drawErr := view.DrawImage(ghostImg, ops); drawErr != nil {
					return nil, drawErr
				}
			}

			if state == GamePause {
				back, backErr := ebiten.NewImage(188, 90, ebiten.FilterDefault)
				if backErr != nil {
					return nil, backErr
				}
				if fillErr := back.Fill(color.Black); fillErr != nil {
					return nil, fillErr
				}

				text.Draw(back, "GAME PAUSED", fontface, 6, 22, color.White)
				text.Draw(back, "PRESS SPACE", fontface, 6, 52, color.White)
				text.Draw(back, "TO CONTINUE", fontface, 6, 82, color.White)

				ops.GeoM.Reset()
				ops.GeoM.Translate(66, 206)
				if drawErr := view.DrawImage(back, ops); drawErr != nil {
					return nil, drawErr
				}
			}
		case GameOver:
		}

		return view, nil
	}, nil
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
