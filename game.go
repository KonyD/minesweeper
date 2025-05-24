package main

import (
	"math/rand"
	"time"
)

type state struct {
	menu       bool
	gameOver   bool
	gameWon    bool
	rows       int32
	cols       int32
	mines      int32
	field      [][]point
	startedAt  time.Time
	finishedAt time.Time
}

func NewGameState() *state {
	return &state{rows: 9, cols: 9, mines: 10}
}

func (s *state) reset() {
	s.gameWon = false
	s.gameOver = false
	s.menu = true
}

func (s *state) getWidth() int {
	return TILE_SIZE * int(s.cols)
}

func (s *state) getHeight() int {
	return TILE_SIZE*int(s.rows) + TILE_SIZE
}

func (s *state) start() {
	// Build grid
	s.field = make([][]point, s.rows)
	for x := range s.rows {
		s.field[x] = make([]point, s.cols)
		for y := range s.cols {
			s.field[x][y] = point{}
		}
	}

	// Plant mines
	m := s.mines
	for m > 0 {
		x, y := rand.Intn(int(s.rows)), rand.Intn(int(s.cols))

		// make sure placements are unique
		if s.field[x][y].hasMine {
			continue
		}

		s.field[x][y].hasMine = true
		// mark neighbours
		s.doForNeighbours(x, y, func(x, y int) {
			s.field[x][y].minesAround++
		})
		m--
	}

	s.menu = false
	s.startedAt = time.Now()
}
