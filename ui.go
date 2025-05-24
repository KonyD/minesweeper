package main

import (
	"fmt"
	"time"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *state) drawMenu() {
	var (
		baseY   float32 = 50
		spacing float32 = 50
	)

	rl.DrawText("Minesweeper", WIDTH/2-50, int32(baseY), 20, rl.Black)
	baseY += spacing * 1.5

	s.createButton(rl.Vector2{X: 0, Y: baseY}, rl.Vector2{X: WIDTH, Y: TILE_SIZE}, "BEGINNER", func() {
		s.rows = 9
		s.cols = 9
		s.mines = 10
	})
	baseY += spacing

	s.createButton(rl.Vector2{X: 0, Y: baseY}, rl.Vector2{X: WIDTH, Y: TILE_SIZE}, "INTERMEDIATE", func() {
		s.rows = 16
		s.cols = 16
		s.mines = 40
	})
	baseY += spacing

	s.createButton(rl.Vector2{X: 0, Y: baseY}, rl.Vector2{X: WIDTH, Y: TILE_SIZE}, "EXPERT", func() {
		s.rows = 30
		s.cols = 30
		s.mines = 99
	})
	baseY += spacing * 2

	s.createButton(rl.Vector2{X: 0, Y: baseY}, rl.Vector2{X: WIDTH, Y: TILE_SIZE}, "START GAME", func() {
		s.start()
	})
}

func (s *state) createButton(pos rl.Vector2, size rl.Vector2, text string, callback func()) {
	button := gui.Button(rl.NewRectangle(pos.X, pos.Y, size.X, size.Y), text)
	if button {
		callback()
	}
}

func (s *state) drawField() {
	w := float32(s.getWidth())
	h := float32(s.getHeight())

	gui.StatusBar(rl.NewRectangle(0, h-TILE_SIZE, w, TILE_SIZE), s.getStatus())
	s.createButton(rl.Vector2{X: w - 65, Y: h - TILE_SIZE + 5}, rl.Vector2{X: 60, Y: TILE_SIZE - 10}, "Restart", func() {
		s.reset()
	})

	for x := range s.field {
		for y := range s.field[x] {
			if s.gameOver {
				var (
					text  string
					color rl.Color
				)

				if s.field[x][y].hasMine {
					text = "*"
					color = rl.Red
				} else if s.field[x][y].minesAround > 0 {
					text = fmt.Sprintf("%d", s.field[x][y].minesAround)
					color = getTextColor(int(s.field[x][y].minesAround))
				}

				rl.DrawText(text, 5+int32(x)*TILE_SIZE, 5+int32(y)*TILE_SIZE, 20, color)
				continue
			}

			rect := rl.NewRectangle(float32(x*TILE_SIZE), float32(y*TILE_SIZE), TILE_SIZE, TILE_SIZE)

			// Mark on right mouse button
			if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
				if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) {
					if !s.field[x][y].open {
						s.field[x][y].marked = !s.field[x][y].marked
					}
				}
			}

			if s.field[x][y].marked {
				rl.DrawText("M", 5+int32(x)*TILE_SIZE, 5+int32(y)*TILE_SIZE, 20, rl.Violet)
			} else if s.field[x][y].open {
				text := ""
				if s.field[x][y].minesAround > 0 {
					text = fmt.Sprintf("%d", s.field[x][y].minesAround)
				}

				rl.DrawText(text, 5+int32(x)*TILE_SIZE, 5+int32(y)*TILE_SIZE, 20, getTextColor(int(s.field[x][y].minesAround)))
			} else {
				if open := gui.Button(rect, ""); open {
					s.revealTile(x, y)
				}
			}

		}
	}
}

func (s *state) drawCongrats() {
	w := WIDTH
	var lineHeight int32 = 50

	if s.gameWon {
		text := "WELL DONE !"
		fontSize := TILE_SIZE
		textWidth := rl.MeasureText(text, int32(fontSize))
		x := (w - int(textWidth)) / 2

		rl.DrawText(text, int32(x), lineHeight, int32(fontSize), rl.White)
	}

	clicked := gui.Button(rl.NewRectangle(0, float32(2*lineHeight), float32(w), TILE_SIZE), "PLAY AGAIN")
	if clicked {
		s.reset()
	}
}

func (s *state) getStatus() string {
	fps := rl.GetFPS()

	var elapsed time.Duration
	if s.gameWon || s.gameOver {
		elapsed = s.finishedAt.Sub(s.startedAt)
	} else {
		elapsed = time.Since(s.startedAt)
	}

	return fmt.Sprintf("FPS: %d | Time: %.2f", fps, elapsed.Seconds())
}
