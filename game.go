package pacman

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type gameState int

type Game struct {
	state    gameState
	rand     *rand.Rand
	maze     *Maze
	sindex   int
	data     *Data
	skinView func(gameState, *Data) (*ebiten.Image, error)
	gridView func(gameState, *Data) (*ebiten.Image, error)
	keys     struct {
		up, left, down, right bool
	}
	direction direction
}

const (
	GameLoading gameState = iota
	GameStart
	GamePause
	GameOver
)

func NewGame() (*Game, error) {
	skin, skinErr := LoadSkin()
	if skinErr != nil {
		return nil, skinErr
	}

	font, fontErr := LoadArcadeFont()
	if fontErr != nil {
		return nil, fontErr
	}

	characters, charactersErr := LoadCharacters()
	if charactersErr != nil {
		return nil, charactersErr
	}

	powers, powersErr := LoadPowers()
	if powersErr != nil {
		return nil, powersErr
	}

	walls, wallsErr := LoadWalls()
	if wallsErr != nil {
		return nil, wallsErr
	}

	mazeView, mazeViewErr := MazeView(walls)
	if mazeViewErr != nil {
		return nil, mazeViewErr
	}

	gridView, gridViewErr := GridView(characters, powers, font, mazeView)
	if gridViewErr != nil {
		return nil, gridViewErr
	}

	skinView, skinViewErr := SkinView(skin, powers, font)
	if skinViewErr != nil {
		return nil, skinViewErr
	}

	return &Game{
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		state:     GameLoading,
		skinView:  skinView,
		gridView:  gridView,
		direction: None,
	}, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.keybord()

	switch g.state {
	case GameLoading:
		if spaceReleased() {
			xcol := g.rand.Intn(Columns)
			g.data = NewData()
			g.maze = NewPopulatedMaze(32, g.rand)
			g.data.grid = g.maze.Get(g.sindex, MazeViewSize/32)
			g.data.pacman = Pacman{
				Position{
					cellX:     xcol,
					cellY:     0,
					posX:      float64((xcol * 32) + 16),
					posY:      16,
					direction: getDirection(g.data.grid[0][xcol].Walls()),
				},
			}
			g.data.grid[0][xcol].active = true
			ghosts := make([]Ghost, 0)
			for i := 1; i < 8; i++ {
				cellX := g.rand.Intn(Columns)
				cellY := g.rand.Intn(2) + (2 * i)
				kind := Ghost1
				if i%4 == 0 {
					kind = Ghost4
				} else if i%3 == 0 {
					kind = Ghost3
				} else if i%2 == 0 {
					kind = Ghost2
				}
				ghosts = append(ghosts, NewGhost(cellX, cellY, kind, North))
			}
			g.data.ghosts = ghosts
			g.state = GameStart
		} else {
			g.data = nil
			g.maze = nil
		}
	case GameStart:
		if spaceReleased() {
			g.state = GamePause
			return nil
		}
		if g.data.pacman.cellY == len(g.data.grid)-8 {
			g.sindex += 4
			if g.sindex > (g.maze.Rows() - MazeViewSize/32) {
				g.maze.GrowBy(16)
			}
			for i := 0; i < len(g.data.ghosts); i++ {
				if g.data.ghosts[i].cellY < g.sindex {
					g.data.ghosts = append(g.data.ghosts[:i], g.data.ghosts[i+1:]...)
				} else {
					g.data.ghosts[i].cellY -= 4
				}
			}
			g.data.grid = g.maze.Get(g.sindex, g.sindex+(MazeViewSize/32))
			g.data.pacman.cellY -= 4
			g.data.gridOffsetY -= 128
		}

		speed := 1.5
		xcell := g.data.pacman.cellX
		ycell := g.data.pacman.cellY

		if !g.data.grid[ycell][xcell].active {
			if math.Abs(float64(
				(g.data.pacman.cellX*32)+16,
			)-(g.data.pacman.posX)) < 10 &&
				math.Abs(float64(
					(g.data.pacman.cellY*32)+16,
				)-(g.data.pacman.posY+g.data.gridOffsetY)) < 10 {
				g.data.grid[ycell][xcell].active = true
				g.data.score += 1
			}
		}

		switch g.direction {
		case North:
			if g.data.pacman.direction != North {
				g.data.pacman.direction = North
			} else {
				if canPacmanMove(
					g.data.pacman.posX,
					g.data.pacman.posY+g.data.gridOffsetY+speed,
					g.data.pacman.cellX,
					g.data.pacman.cellY,
					g.data.grid[ycell][xcell].walls,
				) {
					if g.data.pacman.posY > 320 {
						g.data.gridOffsetY += speed
						for i := 0; i < len(g.data.ghosts); i++ {
							g.data.ghosts[i].posY -= speed
						}
					} else {
						g.data.pacman.posY += speed
					}
				}
				if g.data.pacman.posY+g.data.gridOffsetY+8 > float64((ycell*32)+32) {
					g.data.pacman.cellY += 1
				}
			}
		case South:
			if g.data.pacman.direction != South {
				g.data.pacman.direction = South
			} else {
				if canPacmanMove(
					g.data.pacman.posX,
					g.data.pacman.posY+g.data.gridOffsetY-speed,
					g.data.pacman.cellX,
					g.data.pacman.cellY,
					g.data.grid[ycell][xcell].walls,
				) {
					if g.data.pacman.posY > 320 && g.data.gridOffsetY > 0 {
						g.data.gridOffsetY -= speed
						for i := 0; i < len(g.data.ghosts); i++ {
							g.data.ghosts[i].posY += speed
						}
					} else {
						g.data.pacman.posY -= speed
					}
				}
				if g.data.pacman.posY+g.data.gridOffsetY-8 < float64(ycell*32) {
					g.data.pacman.cellY -= 1
				}
			}
		case East:
			if g.data.pacman.direction != East {
				g.data.pacman.direction = East
			} else {
				if canPacmanMove(
					g.data.pacman.posX+speed,
					g.data.pacman.posY+g.data.gridOffsetY,
					g.data.pacman.cellX,
					g.data.pacman.cellY,
					g.data.grid[ycell][xcell].walls,
				) {
					g.data.pacman.posX += speed
				}
				if g.data.pacman.posX+8 > float64((xcell*32)+32) {
					g.data.pacman.cellX += 1
				}
			}
		case West:
			if g.data.pacman.direction != West {
				g.data.pacman.direction = West
			} else {
				if canPacmanMove(
					g.data.pacman.posX-speed,
					g.data.pacman.posY+g.data.gridOffsetY,
					g.data.pacman.cellX,
					g.data.pacman.cellY,
					g.data.grid[ycell][xcell].walls,
				) {
					g.data.pacman.posX -= speed
				}
				if g.data.pacman.posX-8 < float64(xcell*32) {
					g.data.pacman.cellX -= 1
				}
			}
		}
	case GamePause:
		if spaceReleased() {
			g.state = GameStart
		}
	case GameOver:
		if spaceReleased() {
			g.state = GameLoading
		}
	default:
		// reset state to GameLoading
		// dont return error
		g.state = GameLoading
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	sview, sviewErr := g.skinView(g.state, g.data)
	if sviewErr != nil {
		return sviewErr
	}

	gview, gviewErr := g.gridView(g.state, g.data)
	if gviewErr != nil {
		return gviewErr
	}

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	if drawErr := screen.DrawImage(sview, ops); drawErr != nil {
		return drawErr
	}

	ops.GeoM.Reset()
	ops.GeoM.Translate(18, 80)
	if drawErr := screen.DrawImage(gview, ops); drawErr != nil {
		return drawErr
	}

	return nil
}

func (g *Game) Run() error {
	return ebiten.Run(func(screen *ebiten.Image) error {
		return g.Update(screen)
	}, 356, 610, 1, "PACMAN")
}

func (g *Game) keybord() {
	if upKeyPressed() {
		g.direction = North
	}
	if downKeyPressed() {
		g.direction = South
	}
	if leftKeyPressed() {
		g.direction = West
	}
	if rightKeyPressed() {
		g.direction = East
	}
}

func canPacmanMove(posX, posY float64, x, y int, walls [4]rune) bool {
	psx := posX - 8
	psy := posY - 8
	pex := posX + 8
	pey := posY + 8

	sx := x * 32
	sy := y * 32
	ex := sx + 32
	ey := sy + 32

	if walls[0] == 'N' {
		if pey > float64(ey-6) {
			return false
		}
	}
	if walls[1] == 'E' {
		if pex > float64(ex-6) {
			return false
		}
	}
	if walls[2] == 'S' {
		if psy < float64(sy+6) {
			return false
		}
	}
	if walls[3] == 'W' {
		if psx < float64(sx+6) {
			return false
		}
	}

	// NW corner
	if pey > float64(ey-4) && psx < float64(sx+4) {
		return false
	}
	// NE
	if pey > float64(ey-4) && pex > float64(ex-4) {
		return false
	}
	// SW
	if psy < float64(sy+4) && psx < float64(sx+4) {
		return false
	}
	// SE
	if psy < float64(sy+4) && pex > float64(ex-4) {
		return false
	}

	return true
}

func getDirection(walls [4]rune) direction {
	if walls[0] != 'N' {
		return North
	}
	if walls[1] != 'E' {
		return East
	}
	if walls[2] != 'S' {
		return South
	}
	if walls[3] != 'W' {
		return West
	}
	return North
}
