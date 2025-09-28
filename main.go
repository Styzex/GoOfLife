package main

import (
	"github.com/Styzex/GoOfLife/config"
	"github.com/Styzex/GoOfLife/simulation"
	"github.com/Styzex/GoOfLife/utils"

	"github.com/charmbracelet/log"
)

func main() {
	log.SetLevel(log.DebugLevel)

	seed, _ := utils.GenSeed()
	log.Debug("The seed was generated", "seed", seed)
	log.Debug("The seeds lenght is", "len", len(seed.String()))

	grid := utils.GenGrid(seed)

	type cell struct {
		alive bool
		x     int
		y     int
	}

	cellMap := new([]cell)

	for i := 0; i < config.GridWidth; i++ {
		for j := 0; j < config.GridHeight; j++ {
			value, err := simulation.CheckLife(i, j, grid)
			if err != nil {
				log.Fatal(err)
			}
			currentCell := cell{
				alive: value,
				x:     i,
				y:     j,
			}
			*cellMap = append(*cellMap, currentCell)
		}
	}

	tick := config.Tick
	for {
		log.Debug("Time step", "tick", tick)
		log.Debug("Cell map", "map", cellMap)

		grid = Update(grid)

		if tick >= 20 {
			break
		}
		tick += 1
	}
}

func Update(grid [][]bool) [][]bool {
	if len(grid) == 0 {
		return grid
	}

	newGrid := make([][]bool, config.GridHeight)
	for i := range newGrid {
		newGrid[i] = make([]bool, config.GridWidth)
	}

	for y := range config.GridHeight {
		for x := range config.GridWidth {
			isAlive, err := simulation.CheckLife(x, y, grid)
			if err != nil {
				log.Error("Check life failed", "err", err)
				continue
			}
			newGrid[y][x] = isAlive
		}
	}

	return newGrid
}
