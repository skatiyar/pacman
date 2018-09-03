package pacman

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const MazeViewSize = 768

func MazeView(
	walls *Walls,
) (func(state gameState, data *Data) (*ebiten.Image, error), error) {
	limeAlpha := color.RGBA{250, 233, 8, 200}

	dot, dotErr := ebiten.NewImage(4, 4, ebiten.FilterDefault)
	if dotErr != nil {
		return nil, dotErr
	}
	if fillErr := dot.Fill(limeAlpha); fillErr != nil {
		return nil, fillErr
	}

	icWallSide, icWallSideErr := ScaleSprite(walls.InActiveSide, 0.5, 0.5)
	if icWallSideErr != nil {
		return nil, icWallSideErr
	}

	icWallCorner, icWallCornerErr := ScaleSprite(walls.InActiveCorner, 0.5, 0.5)
	if icWallCornerErr != nil {
		return nil, icWallCornerErr
	}

	mazeView, mazeViewErr := ebiten.NewImage(32*Columns, MazeViewSize, ebiten.FilterDefault)
	if mazeViewErr != nil {
		return nil, mazeViewErr
	}

	var lastGrid [][Columns]Cell

	return func(state gameState, data *Data) (*ebiten.Image, error) {
		if equal, copy := deepEqual(lastGrid, data.grid); equal {
			return mazeView, nil
		} else {
			lastGrid = copy
		}

		if clearErr := mazeView.Clear(); clearErr != nil {
			return nil, clearErr
		}

		gridLength := MazeViewSize
		ops := &ebiten.DrawImageOptions{}

		for i := 0; i < len(data.grid); i++ {
			for j := 0; j < len(data.grid[i]); j++ {
				side := icWallSide
				corner := icWallCorner

				if !data.grid[i][j].Active() {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(j*32)+14, float64(gridLength-((i*32)+18)))
					if drawErr := mazeView.DrawImage(dot, ops); drawErr != nil {
						return nil, drawErr
					}
				}

				cellWalls := data.grid[i][j].Walls()
				if cellWalls[0] == 'N' {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(j*32)+6, float64(gridLength-((i*32)+32)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[1] == 'E' {
					ops.GeoM.Reset()
					ops.GeoM.Rotate(1.5708)
					ops.GeoM.Translate(float64(j*32)+32, float64(gridLength-((i*32)+26)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[2] == 'S' {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(j*32)+6, float64(gridLength-((i*32)+6)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[3] == 'W' {
					ops.GeoM.Reset()
					ops.GeoM.Rotate(1.5708)
					ops.GeoM.Translate(float64(j*32)+6, float64(gridLength-((i*32)+26)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}

				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*32)+26, float64(gridLength-((i*32)+32)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}

				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*32), float64(gridLength-((i*32)+32)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}

				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*32)+26, float64(gridLength-((i*32)+6)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}

				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*32), float64(gridLength-((i*32)+6)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}
			}
		}

		return mazeView, nil
	}, nil
}

func deepEqual(previous, next [][Columns]Cell) (bool, [][Columns]Cell) {
	deepCopy := func(src [][Columns]Cell) [][Columns]Cell {
		copy := make([][Columns]Cell, 0)
		for i := 0; i < len(next); i++ {
			row := [Columns]Cell{}
			for j := 0; j < Columns; j++ {
				row[j] = next[i][j]
			}
			copy = append(copy, row)
		}
		return copy
	}
	if len(previous) != len(next) {
		return false, deepCopy(next)
	}
	for i := 0; i < len(previous); i++ {
		if previous[i] != next[i] {
			return false, deepCopy(next)
		}
		for j := 0; j < Columns; j++ {
			if previous[i][j] != next[i][j] {
				return false, deepCopy(next)
			}
		}
	}

	return true, next
}
