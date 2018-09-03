package pacman

type direction int
type ghostType int
type powerType int

type Data struct {
	grid        [][Columns]Cell
	lifes       int
	score       int
	pacman      Pacman
	ghosts      []Ghost
	gridOffsetY float64
}

const (
	North direction = iota
	East
	South
	West

	Ghost1 ghostType = iota
	Ghost2
	Ghost3
	Ghost4

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

type Power struct {
	Position
	kind powerType
}
