package main

import "time"

type point struct {
	hasMine     bool
	open        bool
	marked      bool
	minesAround int32
}

func (s *state) revealTile(x, y int) {
	if s.field[x][y].open {
		return
	}

	s.field[x][y].open = true

	if s.field[x][y].hasMine {
		s.gameOver = true
		s.finishedAt = time.Now()
		return
	}

	s.gameWon = s.checkWin()

	// No neighbors, reveal all adjacent tiles recursively
	if s.field[x][y].minesAround == 0 {
		s.doForNeighbours(x, y, func(nx, ny int) {
			s.revealTile(nx, ny)
		})
	}
}

func (s *state) doForNeighbours(x, y int, do func(int, int)) {
	// with diagnonals
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := range len(dx) {
		nx, ny := x+dx[i], y+dy[i]

		if nx >= 0 && nx < int(s.rows) && ny >= 0 && ny < int(s.cols) {
			do(nx, ny)
		}
	}
}

func (s *state) checkWin() bool {
	open := 0
	total := int(s.rows * s.cols)

	for x := range s.rows {
		for y := range s.cols {
			if s.field[x][y].open {
				open++
			}
		}
	}

	return open == total-int(s.mines)
}
