package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib-go [core] example - basic window")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window on Omarchy!", 190, 200, 20, rl.LightGray)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
