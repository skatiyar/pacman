package pacman

type direction int
type ghostType int
type powerType int

type Cell struct {
	walls  [4]rune
	active bool
}

type Data struct {
	grid        [][Columns]Cell
	lifes       int
	score       int
	pacman      Pacman
	ghosts      []Ghost
	powers      []Power
	gridOffsetY float64
	invincible  bool
}

const (
	North direction = iota
	East
	South
	West
)

const (
	Ghost1 ghostType = iota
	Ghost2
	Ghost3
	Ghost4
)

const (
	Life powerType = iota
	Invincibility
)

func NewData() *Data {
	return &Data{
		lifes: 5,
		score: 1,
	}
}

func rowsToCells(rows [][Columns][4]rune) [][Columns]Cell {
	cells := make([][Columns]Cell, len(rows), len(rows))
	for i := 0; i < len(rows); i++ {
		for j := 0; j < Columns; j++ {
			cells[i][j].walls = rows[i][j]
		}
	}
	return cells
}

type Position struct {
	cellX, cellY int
	posX, posY   float64
	direction    direction
}

type Pacman struct {
	Position
}

type Ghost struct {
	Position
	visited []Position
	kind    ghostType
}

func NewGhost(x, y int, kind ghostType, dir direction) Ghost {
	return Ghost{
		Position{
			cellX:     x,
			cellY:     y,
			posX:      float64((x * CellSize) + CellSize/2),
			posY:      float64((y * CellSize) + CellSize/2),
			direction: dir,
		},
		[]Position{
			Position{
				cellX:     x,
				cellY:     y,
				direction: dir,
			},
		},
		kind,
	}
}

type Power struct {
	Position
	kind powerType
}

func NewPower(x, y int, kind powerType) Power {
	return Power{
		Position{
			cellX: x,
			cellY: y,
		},
		kind,
	}
}
