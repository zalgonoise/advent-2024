package grid

// Grid is a fourth-quadrant grid
//
//	     |
//	   Q1|Q2
//	-----+-----
//	   Q3|Q4 --> y increments (negative pane); x increments (positive pane)
//	     |
type Grid[T comparable] struct {
	Height, Width int
	Values        [][]T
}

type Coord struct {
	X, Y int
}

func (c Coord) Add(x, y int) Coord {
	return Coord{c.X + x, c.Y + y}
}

func (c Coord) Up() Coord {
	return Coord{c.X, c.Y - 1}
}

func (c Coord) Down() Coord {
	return Coord{c.X, c.Y + 1}
}

func (c Coord) Left() Coord {
	return Coord{c.X - 1, c.Y}
}

func (c Coord) Right() Coord {
	return Coord{c.X + 1, c.Y}
}

func (g Grid[T]) Value(coord Coord) (T, bool) {
	if coord.Y < 0 || coord.Y >= len(g.Values) {
		return *new(T), false
	}

	if coord.X < 0 || coord.X >= len(g.Values[coord.Y]) {
		return *new(T), false
	}

	return g.Values[coord.Y][coord.X], true
}

func FindAll[T comparable](g Grid[T], value T) []Coord {
	if len(g.Values) == 0 {
		return []Coord{}
	}

	coords := make([]Coord, 0, len(g.Values)*len(g.Values[0]))

	for y := range g.Values {
		for x := range g.Values[y] {
			if g.Values[y][x] == value {
				coords = append(coords, Coord{x, y})
			}
		}
	}

	return coords
}

// grid in which quadrant?
//      |
//    Q1|Q2
// -----+-----
//    Q3|Q4 --> y increments (negative pane); x increments (positive pane)
//      |

// which is the zero point?
// 0?   |
//      |
// -----0?----
//      |
//      |
