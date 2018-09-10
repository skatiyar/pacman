package pacman

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type gameState int

type Game struct {
	state       gameState
	rand        *rand.Rand
	maze        *Maze
	data        *Data
	skinView    func(gameState, *Data) (*ebiten.Image, error)
	gridView    func(gameState, *Data) (*ebiten.Image, error)
	direction   direction
	powerTicker *time.Ticker
}

const (
	GameLoading gameState = iota
	GameStart
	GamePause
	GameOver

	OffsetY = CellSize * 10
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
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
		state:    GameLoading,
		skinView: skinView,
		gridView: gridView,
	}, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	switch g.state {
	case GameLoading:
		if spaceReleased() {
			xcol := g.rand.Intn(Columns)
			g.data = NewData()
			g.maze = NewPopulatedMaze(32, g.rand)
			g.data.grid = g.maze.Get(0, MazeViewSize/CellSize)
			g.data.pacman = Pacman{
				Position{
					cellX:     xcol,
					cellY:     0,
					posX:      float64((xcol * CellSize) + (CellSize / 2)),
					posY:      CellSize / 2,
					direction: North,
				},
			}
			g.direction = g.data.pacman.direction
			g.data.grid[0][xcol].active = true

			powers := make([]Power, 0)
			for i := 0; i < MazeViewSize/CellSize; i += 4 {
				cellX := g.rand.Intn(Columns)
				cellY := g.rand.Intn(4) + i
				kind := Invincibility
				if (cellY-i)%2 == 0 {
					kind = Life
				}
				powers = append(powers, NewPower(cellX, cellY, kind))
			}
			g.data.powers = powers

			ghosts := make([]Ghost, 0)
			for i := 0; i < MazeViewSize/CellSize; i += 2 {
				cellX := g.rand.Intn(Columns)
				cellY := g.rand.Intn(4) + i
				kind := Ghost1
				if (cellY-i)%4 == 0 {
					kind = Ghost4
				} else if (cellY-i)%3 == 0 {
					kind = Ghost3
				} else if (cellY-i)%2 == 0 {
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
		} else if g.data.lifes < 1 {
			g.state = GameOver
		} else {
			if g.data.pacman.cellY == len(g.data.grid)-8 {
				g.maze.Compact(4)
				if (g.maze.Rows() - MazeViewSize/CellSize) < 4 {
					g.maze.GrowBy(16)
				}
				g.data.grid = g.maze.Get(0, MazeViewSize/CellSize)
				g.data.pacman.cellY -= 4
				g.data.gridOffsetY -= CellSize * 4

				for i := 0; i < len(g.data.powers); i++ {
					g.data.powers[i].cellY -= 4
					if g.data.powers[i].cellY < 0 {
						cellX := g.rand.Intn(Columns)
						cellY := g.rand.Intn(4) + ((MazeViewSize / CellSize) - 4)
						g.data.powers[i] = NewPower(cellX, cellY, g.data.powers[i].kind)
					}
				}
				for i := 0; i < len(g.data.ghosts); i++ {
					g.data.ghosts[i].cellY -= 4
					g.data.ghosts[i].posY -= CellSize * 4
					if g.data.ghosts[i].cellY < 0 {
						cellX := g.rand.Intn(Columns)
						cellY := g.rand.Intn(4) + ((MazeViewSize / CellSize) - 4)
						g.data.ghosts[i] = NewGhost(
							cellX, cellY, g.data.ghosts[i].kind, North)
					}
				}
			}

			g.keybord()
			g.movePacman()

			// check powers
			for i := 0; i < len(g.data.powers); i++ {
				cellX := g.rand.Intn(Columns)
				cellY := g.rand.Intn(4) +
					(((g.data.powers[i].cellY / 4) * 4) + (MazeViewSize / CellSize))
				if g.pacmanTouchesPower(i) {
					switch g.data.powers[i].kind {
					case Life:
						if g.data.lifes < MaxLifes {
							g.data.lifes += 1
						}
					case Invincibility:
						if !g.data.invincible {
							g.data.invincible = true
						}
						g.startCountdown(15)
					}
					g.data.powers[i] = NewPower(cellX, cellY, g.data.powers[i].kind)
				}
			}
			// check ghosts
			for i := 0; i < len(g.data.ghosts); i++ {
				if g.pacmanTouchesGhost(i) {
					if !g.data.invincible {
						g.data.lifes -= 1
					} else {
						g.data.score += 200
					}
					cellX := g.rand.Intn(Columns)
					cellY := g.rand.Intn(4) +
						(((g.data.ghosts[i].cellY / 4) * 4) + (MazeViewSize / CellSize))
					g.data.ghosts[i] = NewGhost(
						cellX, cellY, g.data.ghosts[i].kind, North)
				}
				g.moveGhost(i)
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
	ops.GeoM.Translate(38, 162)
	if drawErr := screen.DrawImage(gview, ops); drawErr != nil {
		return drawErr
	}

	return nil
}

func (g *Game) Run() error {
	return ebiten.Run(func(screen *ebiten.Image) error {
		return g.Update(screen)
	}, 712, 1220, 1/ebiten.DeviceScaleFactor(), "PACMAN")
}

func (g *Game) keybord() {
	if g.data != nil {
		walls := g.data.grid[g.data.pacman.cellY][g.data.pacman.cellX].Walls()
		if upKeyPressed() {
			if walls[0] == '_' {
				g.direction = North
			}
		}
		if downKeyPressed() {
			if walls[2] == '_' {
				g.direction = South
			}
		}
		if leftKeyPressed() {
			if walls[3] == '_' {
				g.direction = West
			}
		}
		if rightKeyPressed() {
			if walls[1] == '_' {
				g.direction = East
			}
		}
	}
}

func (g *Game) startCountdown(duration int) {
	if g.powerTicker != nil {
		g.powerTicker.Stop()
	}

	g.powerTicker = time.NewTicker(time.Duration(duration) * time.Second)
	go func() {
		select {
		case <-g.powerTicker.C:
			g.data.invincible = false
		}
	}()
}

func canMove(size float64, posX, posY float64, x, y int, walls [4]rune) bool {
	psx := posX - size
	psy := posY - size
	pex := posX + size
	pey := posY + size

	sx := x * CellSize
	sy := y * CellSize
	ex := sx + CellSize
	ey := sy + CellSize

	if walls[0] == 'N' {
		if pey > float64(ey-12) {
			return false
		}
	}
	if walls[1] == 'E' {
		if pex > float64(ex-12) {
			return false
		}
	}
	if walls[2] == 'S' {
		if psy < float64(sy+12) {
			return false
		}
	}
	if walls[3] == 'W' {
		if psx < float64(sx+12) {
			return false
		}
	}

	// NW corner
	if pey > float64(ey-12) && psx < float64(sx+12) {
		return false
	}
	// NE
	if pey > float64(ey-12) && pex > float64(ex-12) {
		return false
	}
	// SW
	if psy < float64(sy+12) && psx < float64(sx+12) {
		return false
	}
	// SE
	if psy < float64(sy+12) && pex > float64(ex-12) {
		return false
	}

	return true
}

func (g *Game) movePacman() {
	speed := 2.0
	xcell := g.data.pacman.cellX
	ycell := g.data.pacman.cellY

	if !g.data.grid[ycell][xcell].active {
		if math.Abs(float64(
			(g.data.pacman.cellX*CellSize)+(CellSize/2),
		)-(g.data.pacman.posX)) < 20 &&
			math.Abs(float64(
				(g.data.pacman.cellY*CellSize)+(CellSize/2),
			)-(g.data.pacman.posY+g.data.gridOffsetY)) < 20 {
			g.data.grid[ycell][xcell].active = true
			g.data.score += 1
		}
	}

	switch g.direction {
	case North, South:
		if g.data.pacman.posX == float64((CellSize*xcell)+(CellSize/2)) {
			g.data.pacman.direction = g.direction
		}
	case East, West:
		if g.data.pacman.posY+g.data.gridOffsetY == float64((CellSize*ycell)+(CellSize/2)) {
			g.data.pacman.direction = g.direction
		}
	}

	switch g.data.pacman.direction {
	case North:
		if canMove(
			20.0,
			g.data.pacman.posX,
			g.data.pacman.posY+g.data.gridOffsetY+speed,
			g.data.pacman.cellX,
			g.data.pacman.cellY,
			g.data.grid[ycell][xcell].walls,
		) {
			if g.data.pacman.posY > OffsetY {
				g.data.gridOffsetY += speed
			} else {
				g.data.pacman.posY += speed
			}
			if g.data.pacman.posY+g.data.gridOffsetY+20 > float64((ycell*CellSize)+CellSize) {
				g.data.pacman.cellY += 1
			}
		}
	case South:
		if canMove(
			20.0,
			g.data.pacman.posX,
			g.data.pacman.posY+g.data.gridOffsetY-speed,
			g.data.pacman.cellX,
			g.data.pacman.cellY,
			g.data.grid[ycell][xcell].walls,
		) {
			if g.data.pacman.posY > OffsetY && g.data.gridOffsetY > 0 {
				g.data.gridOffsetY -= speed
			} else {
				g.data.pacman.posY -= speed
			}
			if g.data.pacman.posY+g.data.gridOffsetY-20 < float64(ycell*CellSize) {
				g.data.pacman.cellY -= 1
			}
		}
	case East:
		if canMove(
			20.0,
			g.data.pacman.posX+speed,
			g.data.pacman.posY+g.data.gridOffsetY,
			g.data.pacman.cellX,
			g.data.pacman.cellY,
			g.data.grid[ycell][xcell].walls,
		) {
			g.data.pacman.posX += speed
			if g.data.pacman.posX+20 > float64((xcell*CellSize)+CellSize) {
				g.data.pacman.cellX += 1
			}
		}
	case West:
		if canMove(
			20.0,
			g.data.pacman.posX-speed,
			g.data.pacman.posY+g.data.gridOffsetY,
			g.data.pacman.cellX,
			g.data.pacman.cellY,
			g.data.grid[ycell][xcell].walls,
		) {
			g.data.pacman.posX -= speed
			if g.data.pacman.posX-20 < float64(xcell*CellSize) {
				g.data.pacman.cellX -= 1
			}
		}
	}
}

func (g *Game) getGhostDirection(i int) direction {
	pacX := g.data.pacman.cellX
	pacY := g.data.pacman.cellY
	ghost := g.data.ghosts[i]

	if pacX < ghost.cellX &&
		g.data.grid[ghost.cellY][ghost.cellX].walls[3] == '_' {
		return West
	}
	if pacX > ghost.cellX &&
		g.data.grid[ghost.cellY][ghost.cellX].walls[1] == '_' {
		return East
	}
	if pacY < ghost.cellY &&
		g.data.grid[ghost.cellY][ghost.cellX].walls[2] == '_' {
		return South
	}
	if pacY > ghost.cellY &&
		g.data.grid[ghost.cellY][ghost.cellX].walls[0] == '_' {
		return North
	}

	switch ghost.direction {
	case North, South:
		if g.data.grid[ghost.cellY][ghost.cellX].walls[1] == '_' {
			return East
		} else if g.data.grid[ghost.cellY][ghost.cellX].walls[3] == '_' {
			return West
		} else {
			if ghost.direction == North &&
				g.data.grid[ghost.cellY][ghost.cellX].walls[0] == 'N' {
				return South
			}
			if ghost.direction == South &&
				g.data.grid[ghost.cellY][ghost.cellX].walls[2] == 'S' {
				return North
			}
		}
	case East, West:
		if g.data.grid[ghost.cellY][ghost.cellX].walls[2] == '_' {
			return South
		} else if g.data.grid[ghost.cellY][ghost.cellX].walls[0] == '_' {
			return North
		} else {
			if ghost.direction == East &&
				g.data.grid[ghost.cellY][ghost.cellX].walls[1] == 'E' {
				return West
			}
			if ghost.direction == West &&
				g.data.grid[ghost.cellY][ghost.cellX].walls[3] == 'W' {
				return East
			}
		}
	}

	return ghost.direction
}

func (g *Game) moveGhost(i int) {
	speed := 1.0
	ghost := g.data.ghosts[i]
	if len(ghost.visited) > Columns {
		g.data.ghosts[i].visited = ghost.visited[1:]
	}
	if ghost.cellY >= MazeViewSize/CellSize {
		return
	}

	if ghost.posX == float64((CellSize*ghost.cellX)+(CellSize/2)) &&
		ghost.posY == float64((CellSize*ghost.cellY)+(CellSize/2)) {
		g.data.ghosts[i].direction = g.getGhostDirection(i)
	}

	switch g.data.ghosts[i].direction {
	case North:
		if canMove(
			20.0,
			g.data.ghosts[i].posX,
			g.data.ghosts[i].posY+speed,
			g.data.ghosts[i].cellX,
			g.data.ghosts[i].cellY,
			g.data.grid[g.data.ghosts[i].cellY][g.data.ghosts[i].cellX].walls,
		) {
			g.data.ghosts[i].posY += speed
			if g.data.ghosts[i].posY+20 > float64((g.data.ghosts[i].cellY*CellSize)+CellSize) {
				g.data.ghosts[i].cellY += 1
			}
		}
	case South:
		if canMove(
			20.0,
			g.data.ghosts[i].posX,
			g.data.ghosts[i].posY-speed,
			g.data.ghosts[i].cellX,
			g.data.ghosts[i].cellY,
			g.data.grid[g.data.ghosts[i].cellY][g.data.ghosts[i].cellX].walls,
		) {
			g.data.ghosts[i].posY -= speed
			if g.data.ghosts[i].posY-20 < float64((g.data.ghosts[i].cellY * CellSize)) {
				g.data.ghosts[i].cellY -= 1
			}
		}
	case East:
		if canMove(
			20.0,
			g.data.ghosts[i].posX+speed,
			g.data.ghosts[i].posY,
			g.data.ghosts[i].cellX,
			g.data.ghosts[i].cellY,
			g.data.grid[ghost.cellY][ghost.cellX].walls,
		) {
			g.data.ghosts[i].posX += speed
			if g.data.ghosts[i].posX+20 > float64((g.data.ghosts[i].cellX*CellSize)+CellSize) {
				g.data.ghosts[i].cellX += 1
			}
		}
	case West:
		if canMove(
			20.0,
			g.data.ghosts[i].posX-speed,
			g.data.ghosts[i].posY,
			g.data.ghosts[i].cellX,
			g.data.ghosts[i].cellY,
			g.data.grid[g.data.ghosts[i].cellY][g.data.ghosts[i].cellX].walls,
		) {
			g.data.ghosts[i].posX -= speed
			if g.data.ghosts[i].posX-20 < float64(g.data.ghosts[i].cellX*CellSize) {
				g.data.ghosts[i].cellX -= 1
			}
		}
	}
}

func (g *Game) pacmanTouchesPower(i int) bool {
	if g.data.powers[i].cellX == g.data.pacman.cellX &&
		g.data.powers[i].cellY == g.data.pacman.cellY {
		posX := float64((g.data.powers[i].cellX * CellSize) + CellSize/2)
		posY := float64((g.data.powers[i].cellY * CellSize) + CellSize/2)
		if math.Abs(posX-g.data.pacman.posX) < 20 &&
			math.Abs(posY-(g.data.pacman.posY+g.data.gridOffsetY)) < 20 {
			return true
		}
	}

	return false
}

func (g *Game) pacmanTouchesGhost(i int) bool {
	if g.data.ghosts[i].cellX == g.data.pacman.cellX &&
		g.data.ghosts[i].cellY == g.data.pacman.cellY {
		posX := g.data.ghosts[i].posX
		posY := g.data.ghosts[i].posY
		if math.Abs(posX-g.data.pacman.posX) < 30 &&
			math.Abs(posY-(g.data.pacman.posY+g.data.gridOffsetY)) < 30 {
			return true
		}
	}

	return false
}
