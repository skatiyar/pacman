package pacman

import (
	"math"
	"math/rand"
)

const MagicNumber = 0.7
const Columns = 10

type Cell struct {
	set    int
	walls  [4]rune
	active bool
}

func (c Cell) Walls() [4]rune {
	return c.walls
}

func (c Cell) Active() bool {
	return c.active
}

type Maze struct {
	maze [][Columns]Cell
	rand *rand.Rand
	rows int
}

func NewMaze(rows int, src *rand.Rand) *Maze {
	return &Maze{
		rand: src,
		maze: make([][Columns]Cell, rows, rows),
		rows: rows,
	}
}

func NewPopulatedMaze(rows int, src *rand.Rand) *Maze {
	maze := NewMaze(rows, src)
	maze.Populate()
	return maze
}

func (m *Maze) Rows() int {
	return m.rows
}

func (m *Maze) Get(from, upto int) [][Columns]Cell {
	return m.maze[from:upto]
}

func (m *Maze) Populate() {
	for row := 0; row < len(m.maze); row++ {
		m.populateRow(row)
	}
}

func (m *Maze) GrowBy(n int) {
	m.maze = append(m.maze, make([][Columns]Cell, n, n)...)
	for i := m.rows; i < m.rows+n; i++ {
		m.populateRow(i)
	}
	m.rows += n
}

func (m *Maze) Compact(n int) {
	m.maze = m.maze[n:]
	for i := 0; i < Columns; i++ {
		m.maze[0][i].walls[2] = 'S'
	}
	m.rows -= n
}

func (m *Maze) populateRow(row int) {
	switch row {
	case 0: // Assume empty maze and assign, unique set to all cells
		for i := 0; i < Columns; i++ {
			m.maze[row][i].set = i + 1
			m.maze[row][i].walls = [4]rune{'N', 'E', 'S', 'W'}
		}
	default:
		current := make([]int, 0)
		for i := 0; i < Columns; i++ {
			if i+1 == Columns || m.maze[row-1][i].walls[1] != m.maze[row-1][i+1].walls[3] {
				current = append(current, i)
				m.rand.Shuffle(len(current), func(i, j int) {
					current[i], current[j] = current[j], current[i]
				})
				gates := int(math.Floor(m.rand.Float64()*(float64(len(current))/2))) + 1
				for j := 0; j < gates; j++ {
					m.maze[row-1][current[j]].walls[0] = '_'
					m.maze[row][current[j]].set = m.maze[row-1][current[j]].set
					m.maze[row][current[j]].walls = [4]rune{'N', 'E', '_', 'W'}
				}
				current = current[:0]
			} else {
				current = append(current, i)
			}
		}
		for i := 0; i < Columns; i++ {
			if m.maze[row][i].set == 0 {
				m.maze[row][i].set = m.rand.Intn(Columns-1) + 1
				m.maze[row][i].walls = [4]rune{'N', 'E', 'S', 'W'}
			}
		}
	}
	m.mergeColumns(row)
}

func except(set map[int]int, max int) int {
	for i := 1; i < max; i++ {
		if set[i] == 0 {
			return i
		}
	}
	return 0
}

func (m *Maze) mergeColumns(row int) {
	for i := 0; i < Columns-1; i++ {
		//  && m.maze[row][i].set != m.maze[row][i+1].set
		if m.rand.Float32() < MagicNumber {
			m.maze[row][i].walls[1] = '_'
			m.maze[row][i+1].set = m.maze[row][i].set
			m.maze[row][i+1].walls[3] = '_'
		}
	}
}
