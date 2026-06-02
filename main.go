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
var map_whith int32
var map_hight int32
var mousePositionXY rl.Vector2
var mousePosition rl.Vector2
var curentmaterialType int32 = 1
var processingWorldMap []Cell
var WorldMap []Cell

type Cell struct {
	color rl.Color
}

func (C Cell) DrawCell(i int32) {
	xy := indexToXY(i)
	rl.DrawRectangle(int32(xy.X)*BaseCellWidth, int32(xy.Y)*BaseCellHeight, BaseCellWidth, BaseCellHeight, C.color)
}

func (C *Cell) UpdateCell(xy rl.Vector2) {
	if getCellByRowCol(xy).color.A == 0 {
		return
	}
	var tempCell rl.Vector2 = xy
	tempCell.Y = tempCell.Y + 1
	swapCellByXY(xy, tempCell)

}

func main() {
	rl.InitWindow(MonitorWidth, MonitorHeight, "raylib-go [core] example - basic window")
	rl.HideCursor()
	CurrentMonitor = rl.GetCurrentMonitor()
	MonitorWidth = int32(rl.GetMonitorWidth(CurrentMonitor))
	MonitorHeight = int32(rl.GetMonitorHeight(CurrentMonitor))
	NumberOfCellsInRow = MonitorWidth / BaseCellWidth
	NumberOfCellsInCol = MonitorHeight / BaseCellHeight

	map_whith = MonitorWidth / BaseCellWidth
	map_hight = MonitorHeight / BaseCellHeight
	NumberOfCellsInWorld = map_whith * map_hight
	processingWorldMap = make([]Cell, NumberOfCellsInWorld)
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		updateMousePostion()
		updateDrawing()
	}
	rl.CloseWindow()
}

func drawBrush() {
	rl.DrawRectangleLines(int32(mousePositionXY.X), int32(mousePositionXY.Y), BaseCellWidth, BaseCellHeight, color.RGBA{200, 0, 0, 255})
}
func updateMousePostion() {
	mousePosition = rl.GetMousePosition()
	mousePosition.X = float32((int32(mousePosition.X) / BaseCellWidth))
	mousePosition.Y = float32((int32(mousePosition.Y) / BaseCellHeight))
	mousePositionXY.X = mousePosition.X * float32(BaseCellWidth)
	mousePositionXY.Y = mousePosition.Y * float32(BaseCellHeight)

}
func addCellAtMouse() {
	AddCellToProcessingWorldMap(mousePosition, color.RGBA{200, 0, 0, 255})

}

func updateDrawing() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	drawBrush()

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		addCellAtMouse()
	}
	ProcessTheProcessingWorldMap(WorldMap, processingWorldMap)
	UpdateProcessedWorldMap(WorldMap, processingWorldMap)
	rl.EndDrawing()
}

func UpdateProcessedWorldMap(worldMap []Cell, processingWorldMap []Cell) {
	copy(worldMap, processingWorldMap)
}
func DrawWorldMap(worldMap []Cell) {
	for i, cell := range worldMap {
		cell.DrawCell(int32(i))
	}
}
func AddCellToProcessingWorldMap(xy rl.Vector2, color rl.Color) {
	cell := getCellByRowCol(xy)

	var newcell = Cell{color}

	*cell = newcell

}
func ProcessTheProcessingWorldMap(worldMap []Cell, processingWorldMap []Cell) {
	for i := len(processingWorldMap) - 1; i >= 0; i-- {
		// fmt.Print(i, len(processingWorldMap), "\n")
		processingWorldMap[i].UpdateCell(indexToXY(int32(i)))
	}
	DrawWorldMap(processingWorldMap)
}
func indexToXY(index int32) rl.Vector2 {
	var x int32 = index % map_whith
	var y int32 = index / map_whith
	var xy rl.Vector2
	xy.X = float32(x)
	xy.Y = float32(y)
	return xy
}
func XYToIndex(xy rl.Vector2) int32 {
	var index int32 = int32(xy.Y)*map_whith + int32(xy.X)

	return index
}
func getCellByIndex(i int32) *Cell {
	return &processingWorldMap[i]
}
func getCellByRowCol(xy rl.Vector2) *Cell {

	return &processingWorldMap[XYToIndex(xy)]
}
func swapCellByXY(xyCell1 rl.Vector2, xyCell2 rl.Vector2) {
	if !iscellValid(xyCell1) {
		return
	}
	if !iscellValid(xyCell2) {
		return
	}
	cell1 := getCellByRowCol(xyCell1)
	cell2 := getCellByRowCol(xyCell2)
	tempCell1 := *cell1
	*cell1 = *cell2
	*cell2 = tempCell1
}
func iscellValid(xy rl.Vector2) bool {
	fmt.Print(xy, "\n")
	if xy.X < 0 {
		return false
	}
	if xy.Y < 0 {
		return false
	}
	if int32(xy.X) > map_whith-1 {
		return false
	}
	if int32(xy.Y) > map_hight-1 {
		return false
	}
	return true
}
