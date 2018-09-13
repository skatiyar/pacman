package pacman

import (
	"math"
	"math/rand"
)

// MagicNumber
const MagicNumber = 0.7
const Columns = 10

type Maze struct {
	maze [][Columns][4]rune
	rand *rand.Rand
	rows int
}

func NewMaze(rows int, src *rand.Rand) *Maze {
	return &Maze{
		rand: src,
		maze: make([][Columns][4]rune, rows, rows),
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

func (m *Maze) Get(from, upto int) [][Columns][4]rune {
	if from > upto {
		from, upto = upto, from
	}
	if upto > m.rows {
		m.GrowBy(upto - m.rows)
	}
	return m.maze[from:upto]
}

func (m *Maze) Populate() {
	for row := 0; row < len(m.maze); row++ {
		m.populateRow(row)
	}
}

func (m *Maze) GrowBy(n int) {
	m.maze = append(m.maze, make([][Columns][4]rune, n, n)...)
	for i := m.rows; i < m.rows+n; i++ {
		m.populateRow(i)
	}
	m.rows += n
}

func (m *Maze) Compact(n int) {
	if n >= m.rows {
		m.rows = 0
		m.maze = m.maze[:0]
	} else {
		m.maze = m.maze[n:]
		for i := 0; i < Columns; i++ {
			m.maze[0][i][2] = 'S'
		}
		m.rows -= n
	}
}

func (m *Maze) populateRow(row int) {
	for i := 0; i < Columns; i++ {
		m.maze[row][i] = [4]rune{'N', 'E', 'S', 'W'}
	}

	switch row {
	case 0: // Assume empty maze.
	default:
		current := make([]int, 0)
		for i := 0; i < Columns; i++ {
			if i+1 == Columns || m.maze[row-1][i][1] != m.maze[row-1][i+1][3] {
				current = append(current, i)
				m.rand.Shuffle(len(current), func(i, j int) {
					current[i], current[j] = current[j], current[i]
				})
				offset := 1
				if len(current) > 2 {
					offset = 2
				}
				gates := int(math.Floor(m.rand.Float64()*(float64(len(current))/2))) + offset
				for j := 0; j < gates; j++ {
					m.maze[row-1][current[j]][0] = '_'
					m.maze[row][current[j]][2] = '_'
				}
				current = current[:0]
			} else {
				current = append(current, i)
			}
		}
	}
	m.mergeColumns(row)
}

func (m *Maze) mergeColumns(row int) {
	for i := 0; i < Columns-1; i++ {
		if m.rand.Float32() < MagicNumber {
			m.maze[row][i][1] = '_'
			m.maze[row][i+1][3] = '_'
		}
	}
}
