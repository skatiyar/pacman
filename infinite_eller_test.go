package pacman

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var randSrc = rand.New(rand.NewSource(1))

func TestNewMaze(t *testing.T) {
	assert.NotNil(t, NewMaze(10, randSrc), "Should not be nil.")
}

func TestRows(t *testing.T) {
	maze := NewMaze(10, randSrc)
	assert.NotNil(t, maze, "Should not be nil.")
	assert.Equal(t, 10, maze.Rows(), "Rows should be same")
}

func TestPopulate(t *testing.T) {
	maze := NewMaze(10, 7, 1)
	maze.Populate()
	rows, columns := maze.Dimensions()
	assert.Equal(t, 10, rows, "Rows should be same")
	assert.Equal(t, 7, columns, "Columns should be same")
}

func TestGrowBy(t *testing.T) {
	maze := NewMaze(2, 5, 1)
	maze.Populate()
	result := [][]Cell{{
		{1, [4]rune{'_', '_', 'S', 'W'}},
		{1, [4]rune{'_', 'E', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', 'W'}},
		{3, [4]rune{'_', '_', 'S', '_'}},
		{3, [4]rune{'N', 'E', 'S', '_'}},
	}, {
		{1, [4]rune{'N', 'E', '_', 'W'}},
		{1, [4]rune{'N', '_', '_', 'W'}},
		{1, [4]rune{'N', 'E', 'S', '_'}},
		{3, [4]rune{'N', '_', '_', 'W'}},
		{3, [4]rune{'N', 'E', 'S', '_'}},
	}}
	rows, columns := maze.Dimensions()
	assert.Equal(t, 2, rows, "Rows should be same")
	assert.Equal(t, 5, columns, "Columns should be same")
	assert.Equal(t, result, maze.Get(0, 2), "Should be equal")
	maze.GrowBy(2)
	result = [][]Cell{{
		{1, [4]rune{'_', '_', 'S', 'W'}},
		{1, [4]rune{'_', 'E', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', 'W'}},
		{3, [4]rune{'_', '_', 'S', '_'}},
		{3, [4]rune{'N', 'E', 'S', '_'}},
	}, {
		{1, [4]rune{'_', 'E', '_', 'W'}},
		{1, [4]rune{'N', '_', '_', 'W'}},
		{1, [4]rune{'_', 'E', 'S', '_'}},
		{3, [4]rune{'N', '_', '_', 'W'}},
		{3, [4]rune{'_', 'E', 'S', '_'}},
	}, {
		{1, [4]rune{'_', '_', '_', 'W'}},
		{1, [4]rune{'_', 'E', 'S', '_'}},
		{1, [4]rune{'N', '_', '_', 'W'}},
		{1, [4]rune{'_', '_', 'S', '_'}},
		{1, [4]rune{'_', 'E', '_', '_'}},
	}, {
		{1, [4]rune{'N', 'E', '_', 'W'}},
		{1, [4]rune{'N', 'E', '_', 'W'}},
		{2, [4]rune{'N', '_', 'S', 'W'}},
		{2, [4]rune{'N', '_', '_', '_'}},
		{2, [4]rune{'N', 'E', '_', '_'}},
	}}
	rows, columns = maze.Dimensions()
	assert.Equal(t, 4, rows, "Rows should be same")
	assert.Equal(t, 5, columns, "Columns should be same")
	assert.Equal(t, result, maze.Get(0, 4), "Should be equal")
}

func TestMergeColumns_FourRows(t *testing.T) {
	maze := Maze{
		maze: [][]Cell{{
			{1, [4]rune{'N', 'E', 'S', 'W'}},
			{2, [4]rune{'N', 'E', 'S', 'W'}},
			{3, [4]rune{'N', 'E', 'S', 'W'}},
			{4, [4]rune{'N', 'E', 'S', 'W'}},
		}},
		columns: 4,
		rows:    1,
		rand:    rand.New(rand.NewSource(1)),
	}
	maze.mergeColumns(0)
	result := []Cell{
		{1, [4]rune{'N', '_', 'S', 'W'}},
		{1, [4]rune{'N', 'E', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', 'W'}},
		{3, [4]rune{'N', 'E', 'S', '_'}},
	}
	assert.Equal(t, maze.maze[0], result, "Arrays should be same after merge")
}

func TestMergeColumns_TenRows(t *testing.T) {
	maze := Maze{
		maze: [][]Cell{{
			{1, [4]rune{'N', 'E', 'S', 'W'}},
			{2, [4]rune{'N', 'E', 'S', 'W'}},
			{3, [4]rune{'N', 'E', 'S', 'W'}},
			{4, [4]rune{'N', 'E', 'S', 'W'}},
			{5, [4]rune{'N', 'E', 'S', 'W'}},
			{6, [4]rune{'N', 'E', 'S', 'W'}},
			{7, [4]rune{'N', 'E', 'S', 'W'}},
			{8, [4]rune{'N', 'E', 'S', 'W'}},
			{9, [4]rune{'N', 'E', 'S', 'W'}},
			{10, [4]rune{'N', 'E', 'S', 'W'}},
		}},
		columns: 10,
		rows:    1,
		rand:    rand.New(rand.NewSource(1)),
	}
	maze.mergeColumns(0)
	result := []Cell{
		{1, [4]rune{'N', '_', 'S', 'W'}},
		{1, [4]rune{'N', 'E', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', 'W'}},
		{3, [4]rune{'N', '_', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', '_'}},
		{3, [4]rune{'N', '_', 'S', '_'}},
		{3, [4]rune{'N', 'E', 'S', '_'}},
	}
	assert.Equal(t, maze.maze[0], result, "Arrays should be same after merge")
}
