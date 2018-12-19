package pacman

type direction int
type ghostType int
type powerType int

type Data struct {
	grid        [][Columns][4]rune
	active      [][Columns]bool
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
	kind ghostType
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
