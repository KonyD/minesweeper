package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game := NewGameState()
	game.reset()

	rl.InitWindow(800, 450, "Minesweeper")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		if game.gameWon || game.menu {
			centerWindow(WIDTH, HEIGHT)
		} else {
			centerWindow(game.getWidth(), game.getHeight())
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		if game.gameWon {
			game.drawCongrats()
		} else if game.menu {
			game.drawMenu()
		} else {
			game.drawField()
		}

		rl.EndDrawing()
	}
}
