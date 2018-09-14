package pacman

import (
	"math"
	"math/rand"
)

// MagicNumber for deciding whether or not
// to tear down wall between two columns.
const MagicNumber = 0.7

// Columns defines the number of cells in
// a row in maze.
const Columns = 10

/*
Maze represents a maze of size rows x 10.

Maze creation is based on modified version of Eller's algorithm
http://weblog.jamisbuck.org/2010/12/29/maze-generation-eller-s-algorithm.
Eller's algorithm creates a perfect maze, a perfect maze has only one path
between any two cells. Secondly to create next row, it requires knowledge
of current row only. Giving us ability to create maze with infinite rows.

Current implementation has been modified to give a non-perfect maze i.e.
it can have more than one path between any two cells.
*/
type Maze struct {
	maze [][Columns][4]rune
	rand *rand.Rand
	rows int
}

// NewMaze returns an unintialized maze with
// given number of rows. Rand source is used for
// all random operations, to give deterministic
// results for a given seed.
func NewMaze(rows int, src *rand.Rand) *Maze {
	return &Maze{
		rand: src,
		maze: make([][Columns][4]rune, rows, rows),
		rows: rows,
	}
}

// NewPopulatedMaze returns a valid maze with given
// number of rows. It calls Populate after calling NewMaze.
func NewPopulatedMaze(rows int, src *rand.Rand) *Maze {
	maze := NewMaze(rows, src)
	maze.Populate()
	return maze
}

// Rows returns number of rows in maze.
func (m *Maze) Rows() int {
	return m.rows
}

// Get returns section of maze specified by indexes.
// It reorders from & upto if from is greater than
// upto. In case upto is bigger than number of rows,
// all rows till last are returned.
func (m *Maze) Get(from, upto int) [][Columns][4]rune {
	if from > upto {
		from, upto = upto, from
	}
	if upto > m.rows {
		m.GrowBy(upto - m.rows)
	}
	return m.maze[from:upto]
}

// Populate creates a valid maze for given grid.
func (m *Maze) Populate() {
	for row := 0; row < len(m.maze); row++ {
		m.populateRow(row)
	}
}

// GrowBy extends the grid by given number &
// creates a valid maze out of new rows.
func (m *Maze) GrowBy(n int) {
	m.maze = append(m.maze, make([][Columns][4]rune, n, n)...)
	for i := m.rows; i < m.rows+n; i++ {
		m.populateRow(i)
	}
	m.rows += n
}

// Compact removes given number of rows from head of grid
// and also creates walls in South direction to prevent player
// from acessing previous rows.
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

// populateRow initializes the given row,
// connencts it to previous row and merges
// columns to create passages.
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

// mergeColumns decides whether to remove
// walls between two columns or not, to create
// horizontal passages in the row.
func (m *Maze) mergeColumns(row int) {
	for i := 0; i < Columns-1; i++ {
		if m.rand.Float32() < MagicNumber {
			m.maze[row][i][1] = '_'
			m.maze[row][i+1][3] = '_'
		}
	}
}
