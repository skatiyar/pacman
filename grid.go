package pacman

import (
	"image/color"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/skatiyar/pacman/assets"
	"github.com/skatiyar/pacman/spritetools"
)

const GridViewSize = 1024

var GrayColor = color.RGBA{236, 240, 241, 255.0}

func GridView(
	characters *assets.Characters,
	powers *assets.Powers,
	arcadeFont *truetype.Font,
	mazeView func(state gameState, data *Data) (*ebiten.Image, error),
) (func(state gameState, data *Data) (*ebiten.Image, error), error) {
	fontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	pacman, pacmanErr := spritetools.ScaleSprite(characters.Pacman, 0.5, 0.5)
	if pacmanErr != nil {
		return nil, pacmanErr
	}

	ghost1, ghost1Err := spritetools.ScaleSprite(characters.Ghost1, 0.5, 0.5)
	if ghost1Err != nil {
		return nil, ghost1Err
	}

	ghost2, ghost2Err := spritetools.ScaleSprite(characters.Ghost2, 0.5, 0.5)
	if ghost2Err != nil {
		return nil, ghost2Err
	}

	ghost3, ghost3Err := spritetools.ScaleSprite(characters.Ghost3, 0.5, 0.5)
	if ghost3Err != nil {
		return nil, ghost3Err
	}

	ghost4, ghost4Err := spritetools.ScaleSprite(characters.Ghost4, 0.5, 0.5)
	if ghost4Err != nil {
		return nil, ghost4Err
	}

	life, lifeErr := spritetools.ScaleSprite(powers.Life, 0.5, 0.5)
	if lifeErr != nil {
		return nil, lifeErr
	}

	invinci, invinciErr := spritetools.ScaleSprite(powers.Invincibility, 0.5, 0.5)
	if invinciErr != nil {
		return nil, invinciErr
	}

	view, viewErr := ebiten.NewImage(64*Columns, GridViewSize, ebiten.FilterDefault)
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
			text.Draw(view, "PRESS SPACE", fontface, 320-176, 512-(10+32), color.White)
			text.Draw(view, "TO BEGIN", fontface, 320-128, 512+(10), color.White)
		case GameStart, GamePause, GameOver:
			mazeView, mazeViewErr := mazeView(state, data)
			if mazeViewErr != nil {
				return nil, mazeViewErr
			}

			ops.GeoM.Reset()
			ops.GeoM.Translate(0,
				-(float64(len(data.grid)*CellSize) - (GridViewSize + data.gridOffsetY)))
			if drawErr := view.DrawImage(mazeView, ops); drawErr != nil {
				return nil, drawErr
			}

			for i := 0; i < len(data.powers); i++ {
				power := data.powers[i]
				powerImg := life
				if power.kind == Invincibility {
					powerImg = invinci
				}
				pwidth, pheight := powerImg.Size()
				ops.GeoM.Reset()
				ops.GeoM.Translate(
					float64((data.powers[i].cellX*CellSize)+pwidth/2),
					-(float64(((data.powers[i].cellY*CellSize)+(CellSize/2))+pheight/2) -
						(GridViewSize + data.gridOffsetY)))
				if drawErr := view.DrawImage(powerImg, ops); drawErr != nil {
					return nil, drawErr
				}
			}

			ops.GeoM.Reset()
			pwidth, pheight := pacman.Size()
			switch data.pacman.direction {
			case North:
				ops.GeoM.Rotate(-1.5708)
				ops.GeoM.Translate(
					data.pacman.posX-float64(pwidth/2),
					GridViewSize-(data.pacman.posY-float64(pheight-(pheight/2))))
			case East:
				ops.GeoM.Translate(
					data.pacman.posX-float64(pwidth/2),
					GridViewSize-(data.pacman.posY+float64(pheight/2)))
			case South:
				ops.GeoM.Rotate(1.5708)
				ops.GeoM.Translate(
					data.pacman.posX+float64(pwidth/2),
					GridViewSize-(data.pacman.posY+float64(pheight/2)))
			case West:
				ops.GeoM.Rotate(3.14159)
				ops.GeoM.Translate(
					data.pacman.posX+float64(pwidth/2),
					GridViewSize-(data.pacman.posY-float64(pheight-(pheight/2))))
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
				if data.invincible {
					ops.ColorM.ChangeHSV(0, 0, 1)
				}
				ops.GeoM.Translate(
					data.ghosts[i].posX-float64(gwidth/2),
					(GridViewSize+data.gridOffsetY)-
						(data.ghosts[i].posY+float64(gheight-(gheight/2))))
				if drawErr := view.DrawImage(ghostImg, ops); drawErr != nil {
					return nil, drawErr
				}
			}

			if state == GamePause {
				back, backErr := ebiten.NewImage(389, 130, ebiten.FilterDefault)
				if backErr != nil {
					return nil, backErr
				}
				if fillErr := back.Fill(color.Black); fillErr != nil {
					return nil, fillErr
				}

				text.Draw(back, "GAME PAUSED", fontface, 24, 65-(10), color.White)
				text.Draw(back, "PRESS SPACE", fontface, 24, 65+(10+31), color.White)

				ops.GeoM.Reset()
				ops.GeoM.Translate(320-(389/2), 512-(130/2))
				if drawErr := view.DrawImage(back, ops); drawErr != nil {
					return nil, drawErr
				}
			} else if state == GameOver {
				back, backErr := ebiten.NewImage(389, 130, ebiten.FilterDefault)
				if backErr != nil {
					return nil, backErr
				}
				if fillErr := back.Fill(color.Black); fillErr != nil {
					return nil, fillErr
				}

				text.Draw(back, "GAME OVER", fontface, 56, 65-(10), color.White)
				text.Draw(back, "PRESS SPACE", fontface, 24, 65+(10+31), color.White)

				ops.GeoM.Reset()
				ops.GeoM.Translate(320-(389/2), 512-(130/2))
				if drawErr := view.DrawImage(back, ops); drawErr != nil {
					return nil, drawErr
				}
			}
		}

		return view, nil
	}, nil
}
