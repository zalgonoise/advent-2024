package part01

import (
	"slices"
	"strings"
)

type Grid[T comparable] struct {
	height, width int
	values        [][]T
}

func Parse(input string) Grid[byte] {
	lines := strings.Split(input, "\n")

	grid := Grid[byte]{
		height: len(lines),
		width:  len(lines[0]),
		values: make([][]byte, 0, len(lines)),
	}

	for i := range lines {
		grid.values = append(grid.values, []byte(lines[i]))
	}

	return grid
}

// XYZ --> 3 + 3 horizontal (front and back)
// ABC --> 3 + 3 vertical (front and back)
// 123 --> 3 + 3 diagonal top-left to bottom-right (front and back)
//     --> 3 + 3 diagonal top-right to bottom-left (front and back)

func BreakdownGrid(grid Grid[byte]) [][]byte {
	set := make([][]byte, 0, (len(grid.values[0])*2)*4)

	h := breakdownHorizontal(grid)
	v := breakdownVertical(grid)
	tl := breakdownDiagonalTopLeft(grid)
	tr := breakdownDiagonalTopRight(grid)

	set = append(set, h...)
	set = append(set, v...)
	set = append(set, tl...)
	set = append(set, tr...)

	return set
}

func Count(grid Grid[byte], sub string) int {
	sets := BreakdownGrid(grid)

	var n int

	for i := range sets {
		n += strings.Count(string(sets[i]), sub)
	}

	rev := []byte(sub)
	slices.Reverse(rev)

	for i := range sets {
		n += strings.Count(string(sets[i]), string(rev))
	}

	return n
}

func breakdownHorizontal(grid Grid[byte]) [][]byte {
	//set := make([][]byte, 0, len(grid.values[0])*2)
	horizontal := make([][]byte, 0, len(grid.values[0]))

	for i := 0; i < len(grid.values); i++ {
		horizontal = append(horizontal, grid.values[i])
	}

	return horizontal

	//set = append(set, horizontal...)

	//revHorizontal := make([][]byte, len(horizontal))
	//for i := range horizontal {
	//	entryCopy := make([]byte, len(horizontal[i]))
	//	copy(entryCopy, horizontal[i])
	//
	//	revHorizontal[i] = entryCopy
	//}
	//
	//for i := range revHorizontal {
	//	slices.Reverse(revHorizontal[i])
	//}
	//
	//set = append(set, revHorizontal...)

	//return set
}

func breakdownVertical(grid Grid[byte]) [][]byte {
	//set := make([][]byte, 0, len(grid.values[0])*2)
	vertical := make([][]byte, 0, len(grid.values[0]))
	lim := len(grid.values)

	for x := 0; x < lim; x++ {
		verticalSet := make([]byte, 0, len(grid.values[0]))

		for y := 0; y < lim; y++ {
			verticalSet = append(verticalSet, grid.values[y][x])
		}

		vertical = append(vertical, verticalSet)
	}

	return vertical
	//set = append(set, vertical...)
	//
	//revVertical := make([][]byte, len(vertical))
	//for i := range vertical {
	//	entryCopy := make([]byte, len(vertical[i]))
	//	copy(entryCopy, vertical[i])
	//
	//	revVertical[i] = entryCopy
	//}
	//
	//for i := range revVertical {
	//	slices.Reverse(revVertical[i])
	//}
	//
	//set = append(set, revVertical...)
	//
	//return set
}

func breakdownDiagonalTopLeft(grid Grid[byte]) [][]byte {
	//set := make([][]byte, 0, len(grid.values[0])*2)

	// XYZ
	// ABC
	// 123
	//
	// idx (Y, X):
	//   -> [2,0]
	//   -> [1,0],[2,1]
	//   -> [0,0],[1,1],[2,2]
	//   -> [0,1],[1,2]
	//   -> [0,2]

	diagonalTopLeft := make([][]byte, 0, len(grid.values[0])*2)
	for dtl := len(grid.values); dtl >= 0; dtl-- {
		dtlSet := make([]byte, 0, len(grid.values[0]))
		x, y := 0, dtl

		for ; y < len(grid.values); y, x = y+1, x+1 {
			dtlSet = append(dtlSet, grid.values[y][x])
		}

		if len(dtlSet) > 0 {
			diagonalTopLeft = append(diagonalTopLeft, dtlSet)
		}
	}

	for dtl := 1; dtl < len(grid.values); dtl++ {
		dtlSet := make([]byte, 0, len(grid.values[0]))
		x, y := dtl, 0

		for ; x < len(grid.values); y, x = y+1, x+1 {
			dtlSet = append(dtlSet, grid.values[y][x])
		}

		if len(dtlSet) > 0 {
			diagonalTopLeft = append(diagonalTopLeft, dtlSet)
		}
	}

	return diagonalTopLeft

	//set = append(set, diagonalTopLeft...)
	//
	//rev := make([][]byte, len(diagonalTopLeft))
	//for i := range diagonalTopLeft {
	//	entryCopy := make([]byte, len(diagonalTopLeft[i]))
	//	copy(entryCopy, diagonalTopLeft[i])
	//
	//	rev[i] = entryCopy
	//}
	//
	//for i := range rev {
	//	slices.Reverse(rev[i])
	//}
	//
	//set = append(set, rev...)
	//
	//return set
}

func breakdownDiagonalTopRight(grid Grid[byte]) [][]byte {
	//set := make([][]byte, 0, len(grid.values[0])*2)

	// XYZ
	// ABC
	// 123
	//
	// idx (Y, X):
	//   -> [0,0]
	//   -> [0,1],[1,0]
	//   -> [0,2],[1,1],[2,0]
	//   -> [1,2],[2,1]
	//   -> [2,2]
	diagonalTopRight := make([][]byte, 0, len(grid.values[0])*2)
	for dtr := 0; dtr < len(grid.values); dtr++ {
		dtrSet := make([]byte, 0, len(grid.values[0]))
		y, x := 0, dtr

		for ; x >= 0; y, x = y+1, x-1 {
			dtrSet = append(dtrSet, grid.values[y][x])
		}

		if len(dtrSet) > 0 {
			diagonalTopRight = append(diagonalTopRight, dtrSet)
		}
	}

	for dtr := 1; dtr < len(grid.values); dtr++ {
		dtrSet := make([]byte, 0, len(grid.values[dtr]))
		x, y := len(grid.values[dtr])-1, dtr

		for ; y < len(grid.values); y, x = y+1, x-1 {
			dtrSet = append(dtrSet, grid.values[y][x])
		}

		if len(dtrSet) > 0 {
			diagonalTopRight = append(diagonalTopRight, dtrSet)
		}
	}

	return diagonalTopRight
	//set = append(set, diagonalTopRight...)
	//
	//rev := make([][]byte, len(diagonalTopRight))
	//for i := range diagonalTopRight {
	//	entryCopy := make([]byte, len(diagonalTopRight[i]))
	//	copy(entryCopy, diagonalTopRight[i])
	//
	//	rev[i] = entryCopy
	//}
	//
	//for i := range rev {
	//	slices.Reverse(rev[i])
	//}
	//
	//set = append(set, rev...)
	//
	//return set
}
