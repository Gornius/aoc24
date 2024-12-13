package main

import "fmt"

func main() {
	filePath := "./challenge_data.txt"

	fmt.Println("==== PART 1 ====")
	part1(filePath)
	// fmt.Println("==== PART 2 ====")
	// part2(filePath)
}

func part1(filePath string) {
	garden, err := parseFile(filePath)
	if err != nil {
		panic(err)
	}

	totalPrice := 0

	regions := garden.FindRegions()
	for _, region := range regions {
		totalPrice += region.GetPrice()
	}

	fmt.Printf("totalPrice: %v\n", totalPrice)

}

// func part2(filePath string)

type Garden struct {
	Plots [][]Plot
}

func (g *Garden) FindRegions() []Region {
	regions := []Region{}
	processedPlots := []Plot{}

	isProcessedAlready := func(coords Coords) bool {
		for _, processedPlot := range processedPlots {
			if processedPlot.Coords.X == coords.X && processedPlot.Coords.Y == coords.Y {
				return true
			}
		}
		return false
	}

	var exploreRegion func(region *Region, currentPlot *Plot)
	exploreRegion = func(region *Region, currentPlot *Plot) {
		region.Plots = append(region.Plots, *currentPlot)
		processedPlots = append(processedPlots, *currentPlot)

		nextCoords := []Coords{
			{X: currentPlot.Coords.X, Y: currentPlot.Coords.Y + 1},
			{X: currentPlot.Coords.X, Y: currentPlot.Coords.Y - 1},
			{X: currentPlot.Coords.X + 1, Y: currentPlot.Coords.Y},
			{X: currentPlot.Coords.X - 1, Y: currentPlot.Coords.Y},
		}

		for _, nextCoord := range nextCoords {
			if nextCoord.X < 0 || nextCoord.Y < 0 {
				continue
			}

			if nextCoord.X >= g.GetWidth() || nextCoord.Y >= g.GetHeight() {
				continue
			}

			nextPlot := g.Plots[nextCoord.Y][nextCoord.X]
			if nextPlot.Plant != currentPlot.Plant {
				continue
			}

			if isProcessedAlready(nextPlot.Coords) {
				continue
			}

			exploreRegion(region, &g.Plots[nextCoord.Y][nextCoord.X])
		}
	}

	for _, row := range g.Plots {
		for _, plot := range row {
			if isProcessedAlready(plot.Coords) {
				continue
			}
			region := &Region{Garden: g}
			exploreRegion(region, &plot)
			regions = append(regions, *region)
		}
	}

	return regions
}

func (g *Garden) ToString() string {
	buf := ""
	for _, row := range g.Plots {
		for _, plot := range row {
			buf += string(plot.Plant)
		}
		buf += "\n"
	}
	return buf
}

func (g *Garden) GetHeight() int {
	return len(g.Plots)
}

func (g *Garden) GetWidth() int {
	return len(g.Plots[0])

}

type Plot struct {
	Garden *Garden
	Coords Coords
	Plant  rune
}

type Coords struct {
	X, Y int
}

type Region struct {
	Plots  []Plot
	Garden *Garden
}

func (r *Region) GetArea() int {
	return len(r.Plots)
}
func (r *Region) GetPerimeter() int {
	fences := 0
	for _, plot := range r.Plots {
		nextCoords := []Coords{
			{X: plot.Coords.X, Y: plot.Coords.Y + 1},
			{X: plot.Coords.X, Y: plot.Coords.Y - 1},
			{X: plot.Coords.X + 1, Y: plot.Coords.Y},
			{X: plot.Coords.X + 1, Y: plot.Coords.Y},
		}

		for _, nextCoord := range nextCoords {
			if nextCoord.X < 0 || nextCoord.Y < 0 {
				fences++
				continue
			}

			if nextCoord.X >= r.Garden.GetWidth() || nextCoord.Y >= r.Garden.GetHeight() {
				fences++
				continue
			}

			nextPlot := r.Garden.Plots[nextCoord.Y][nextCoord.X]
			if nextPlot.Plant != plot.Plant {
				fences++
			}
		}
	}

	return fences
}
func (r *Region) GetPrice() int {
	return r.GetPerimeter() * r.GetArea()
}
