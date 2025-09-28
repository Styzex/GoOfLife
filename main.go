package main

import (
	"github.com/Styzex/GoOfLife/config"
	"github.com/Styzex/GoOfLife/simulation"
	"github.com/Styzex/GoOfLife/utils"

	"github.com/charmbracelet/log"
)

type cell struct {
	alive bool
	x     int
	y     int
}

func main() {

	seed, _ := utils.GenSeed()
	//seed, _ := new(big.Int).SetString("99999999999999999999", 10)
	log.Debug("The seed was generated", "seed", seed)
	log.Debug("The seeds lenght is", "len", len(seed.String()))

	grid := utils.GenGrid(seed)

	tick := config.Tick
	for {
		log.Info("Time step", "tick", tick)

		cellMap := new([]cell)
		grid = Update(grid)
		UpdateCellMap(cellMap, grid)
		utils.PrintCellMap(grid)

		log.Debug("Cell map", "map", cellMap)

		liveCells := utils.CountLiveCells(grid)
		log.Debug("Live cells remaining", "count", liveCells)

		if liveCells == 0 {
			log.Info("All cells died - simulation ended", "final_tick", tick)
			break
		}

		if tick >= 10000 {
			break
		}
		tick += 1
	}
}

func UpdateCellMap(cellMap *[]cell, grid [][]bool) {
	for x := 0; x < config.GridWidth; x++ {
		for y := 0; y < config.GridHeight; y++ {
			value, err := simulation.CheckLife(x, y, grid)
			if err != nil {
				log.Fatal(err)
			}
			currentCell := cell{
				alive: value,
				x:     x,
				y:     y,
			}
			*cellMap = append(*cellMap, currentCell)
		}
	}
}

func Update(grid [][]bool) [][]bool {
	if len(grid) == 0 {
		return grid
	}

	newGrid := make([][]bool, config.GridWidth)
	for i := range newGrid {
		newGrid[i] = make([]bool, config.GridHeight)
	}

	for x := 0; x < config.GridWidth; x++ {
		for y := 0; y < config.GridHeight; y++ {
			isAlive, err := simulation.CheckLife(x, y, grid)
			if err != nil {
				log.Error("Check life failed", "err", err)
				continue
			}
			newGrid[x][y] = isAlive
		}
	}

	return newGrid
}
