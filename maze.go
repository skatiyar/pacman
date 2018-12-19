package pacman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/skatiyar/pacman/assets"
	"github.com/skatiyar/pacman/spritetools"
)

const MazeViewSize = 1536
const CellSize = 64

func MazeView(
	walls *assets.Walls,
) (func(state gameState, data *Data) (*ebiten.Image, error), error) {
	icWallSide, icWallSideErr := spritetools.ScaleSprite(walls.InActiveSide, 1.0, 1.0)
	if icWallSideErr != nil {
		return nil, icWallSideErr
	}

	icWallCorner, icWallCornerErr := spritetools.ScaleSprite(walls.InActiveCorner, 1.0, 1.0)
	if icWallCornerErr != nil {
		return nil, icWallCornerErr
	}

	mazeView, mazeViewErr := ebiten.NewImage(CellSize*Columns, MazeViewSize, ebiten.FilterDefault)
	if mazeViewErr != nil {
		return nil, mazeViewErr
	}

	var lastGrid [][Columns][4]rune

	return func(state gameState, data *Data) (*ebiten.Image, error) {
		if equal, copy := deepEqual(lastGrid, data.grid); equal {
			return mazeView, nil
		} else {
			lastGrid = copy
		}

		if clearErr := mazeView.Clear(); clearErr != nil {
			return nil, clearErr
		}

		ops := &ebiten.DrawImageOptions{}

		for i := 0; i < len(data.grid); i++ {
			for j := 0; j < len(data.grid[i]); j++ {
				side := icWallSide
				corner := icWallCorner

				cellWalls := data.grid[i][j]
				if cellWalls[0] == 'N' {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(j*CellSize)+12,
						float64(MazeViewSize-((i*CellSize)+CellSize)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[1] == 'E' {
					ops.GeoM.Reset()
					ops.GeoM.Rotate(1.5708)
					ops.GeoM.Translate(float64(j*CellSize)+CellSize,
						float64(MazeViewSize-((i*CellSize)+52)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[2] == 'S' {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(j*CellSize)+12,
						float64(MazeViewSize-((i*CellSize)+12)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[3] == 'W' {
					ops.GeoM.Reset()
					ops.GeoM.Rotate(1.5708)
					ops.GeoM.Translate(float64(j*CellSize)+12,
						float64(MazeViewSize-((i*CellSize)+52)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}

				// Corners NE
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*CellSize)+52,
					float64(MazeViewSize-((i*CellSize)+CellSize)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}

				// NW
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*CellSize),
					float64(MazeViewSize-((i*CellSize)+CellSize)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}

				// SE
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*CellSize)+52,
					float64(MazeViewSize-((i*CellSize)+12)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}

				// SW
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*CellSize),
					float64(MazeViewSize-((i*CellSize)+12)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}
			}
		}

		return mazeView, nil
	}, nil
}

func deepEqual(previous, next [][Columns][4]rune) (bool, [][Columns][4]rune) {
	deepCopy := func(src [][Columns][4]rune) [][Columns][4]rune {
		copy := make([][Columns][4]rune, 0)
		for i := 0; i < len(next); i++ {
			row := [Columns][4]rune{}
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
