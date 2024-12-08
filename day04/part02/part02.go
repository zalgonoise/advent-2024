package part02

import (
	"github.com/zalgonoise/advent-2024/grid"
	"strings"
)

func Parse(input string) grid.Grid[byte] {
	lines := strings.Split(input, "\n")

	g := grid.Grid[byte]{
		Height: len(lines),
		Width:  len(lines[0]),
		Values: make([][]byte, 0, len(lines)),
	}

	for i := range lines {
		g.Values = append(g.Values, []byte(lines[i]))
	}

	return g
}

func FindAll(g grid.Grid[byte]) []grid.Coord {
	centers := grid.FindAll(g, 'A')
	valid := make([]grid.Coord, 0, len(centers))

	for i := range centers {
		if isCenterValid(g, centers[i]) {
			valid = append(valid, centers[i])
		}
	}

	return valid
}

func isCenterValid(g grid.Grid[byte], coord grid.Coord) bool {
	topLeft := coord.Add(-1, -1)
	topRight := coord.Add(1, -1)
	bottomLeft := coord.Add(-1, 1)
	bottomRight := coord.Add(1, 1)

	tl, ok := g.Value(topLeft)
	if !ok {
		return false
	}

	tr, ok := g.Value(topRight)
	if !ok {
		return false
	}

	bl, ok := g.Value(bottomLeft)
	if !ok {
		return false
	}

	br, ok := g.Value(bottomRight)
	if !ok {
		return false
	}

	if ((tl == 'M' && br == 'S') || (tl == 'S' && br == 'M')) &&
		((tr == 'M' && bl == 'S') || (tr == 'S' && bl == 'M')) {
		return true
	}

	return false
}
