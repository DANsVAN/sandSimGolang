package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var BaseCellWidth int32 = 10
var BaseCellHeight int32 = 10
var CurrentMonitor int
var MonitorWidth int32
var MonitorHeight int32
var NumberOfCellsInWorld int32
var NumberOfCellsInRow int32
var NumberOfCellsInCol int32

type Cell struct {
	x, y, numberOfCols, numberOfRows, width, height, index int32
	color                                                  rl.Color
}

func (C Cell) DrawCell() {
	rl.DrawRectangle(C.x, C.y, C.width, C.height, C.color)
}
func (C *Cell) UpdateCell() {
	// if C.y < (NumberOfCellsInCol-1)*BaseCellHeight {
	// 	C.y += BaseCellHeight
	// }
	if C.numberOfRows < (NumberOfCellsInCol - 1) {
		C.index += NumberOfCellsInRow
		C.numberOfRows += 1
		C.y = C.numberOfRows * 10
	}
}

func main() {
	rl.InitWindow(MonitorWidth, MonitorHeight, "raylib-go [core] example - basic window")
	CurrentMonitor = rl.GetCurrentMonitor()
	MonitorWidth = int32(rl.GetMonitorWidth(CurrentMonitor))
	MonitorHeight = int32(rl.GetMonitorHeight(CurrentMonitor))
	NumberOfCellsInRow = MonitorWidth / BaseCellWidth
	NumberOfCellsInCol = MonitorHeight / BaseCellHeight
	NumberOfCellsInWorld = NumberOfCellsInRow * NumberOfCellsInCol
	WorldMap := make([]Cell, NumberOfCellsInWorld)
	ProcessingWorldMap := make([]Cell, NumberOfCellsInWorld)
	fmt.Println("---------------------------")
	fmt.Println(MonitorWidth)
	fmt.Println(MonitorHeight)
	fmt.Println("---------------------------")
	rl.SetTargetFPS(60)
	AddCellToProcessingWorldMap(ProcessingWorldMap, 0, 0, color.RGBA{200, 0, 0, 255})
	AddCellToProcessingWorldMap(ProcessingWorldMap, 50, 50, color.RGBA{200, 0, 0, 255})
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		UpdateProcessedWorldMap(WorldMap, ProcessingWorldMap)
		ProcessTheProcessingWorldMap(WorldMap, ProcessingWorldMap)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func UpdateProcessedWorldMap(worldMap []Cell, processingWorldMap []Cell) {
	copy(worldMap, processingWorldMap)
}
func DrawWorldMap(worldMap []Cell) {
	for _, cell := range worldMap {
		cell.DrawCell()
	}
}
func AddCellToProcessingWorldMap(processingWorldMap []Cell, numberOfCols int32, numberOfrows int32, color rl.Color) {
	var index int32 = (numberOfrows * NumberOfCellsInRow) + numberOfCols
	var x = numberOfCols * BaseCellWidth
	var y = numberOfrows * BaseCellHeight
	processingWorldMap[index] = Cell{x, y, numberOfCols, numberOfrows, BaseCellWidth, BaseCellHeight, index, color}
}
func ProcessTheProcessingWorldMap(worldMap []Cell, processingWorldMap []Cell) {
	// loop backwords bottom to top
	for i := len(processingWorldMap) - 1; i >= 0; i-- {
		processingWorldMap[i].UpdateCell()
	}
	DrawWorldMap(worldMap)
}
