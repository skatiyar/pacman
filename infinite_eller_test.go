package pacman

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randSrc() *rand.Rand {
	return rand.New(rand.NewSource(1))
}

func TestNewMaze(t *testing.T) {
	assert.NotNil(t, NewMaze(10, randSrc()), "Should not be nil.")
}

func TestRows(t *testing.T) {
	maze := NewMaze(10, randSrc())
	assert.NotNil(t, maze, "Should not be nil.")
	assert.Equal(t, 10, maze.Rows(), "Rows should be same")
}

func TestPopulate(t *testing.T) {
	result := [][Columns][4]rune{{
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'_', 'E', 'S', '_'},
		[4]rune{'_', '_', 'S', 'W'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', 'E', 'S', '_'},
	}, {
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', 'E', 'S', '_'},
		[4]rune{'N', '_', '_', 'W'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'N', 'E', '_', '_'},
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'N', 'E', 'S', '_'},
	}}
	maze := NewMaze(2, randSrc())
	maze.Populate()
	assert.Equal(t, 2, maze.Rows(), "Rows should be same")
	assert.Equal(t, result, maze.Get(0, 2), "Should be equal")
}

func TestGrowBy(t *testing.T) {
	maze := NewMaze(2, randSrc())
	maze.Populate()
	result := [][Columns][4]rune{{
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'_', 'E', 'S', '_'},
		[4]rune{'_', '_', 'S', 'W'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', 'E', 'S', '_'},
	}, {
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', 'E', 'S', '_'},
		[4]rune{'N', '_', '_', 'W'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'N', 'E', '_', '_'},
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'N', 'E', 'S', '_'},
	}}
	assert.Equal(t, 2, maze.Rows(), "Rows should be same")
	assert.Equal(t, result, maze.Get(0, 2), "Should be equal")
	maze.GrowBy(2)
	result = [][Columns][4]rune{{
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'_', 'E', 'S', '_'},
		[4]rune{'_', '_', 'S', 'W'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', 'E', 'S', '_'},
	}, {
		[4]rune{'_', '_', 'S', 'W'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'_', '_', 'S', '_'},
		[4]rune{'N', 'E', 'S', '_'},
		[4]rune{'_', '_', '_', 'W'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'_', 'E', '_', '_'},
		[4]rune{'_', '_', 'S', 'W'},
		[4]rune{'N', 'E', 'S', '_'},
	}, {
		[4]rune{'N', '_', '_', 'W'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'_', '_', '_', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'_', '_', '_', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'_', '_', '_', '_'},
		[4]rune{'N', 'E', 'S', '_'},
	}, {
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', '_', 'S', '_'},
		[4]rune{'N', 'E', '_', '_'},
		[4]rune{'N', 'E', 'S', 'W'},
		[4]rune{'N', 'E', '_', 'W'},
		[4]rune{'N', 'E', 'S', 'W'},
		[4]rune{'N', '_', 'S', 'W'},
		[4]rune{'N', '_', '_', '_'},
		[4]rune{'N', 'E', 'S', '_'},
	}}
	assert.Equal(t, 4, maze.Rows(), "Rows should be same")
	assert.Equal(t, result, maze.Get(0, 4), "Should be equal")
}
