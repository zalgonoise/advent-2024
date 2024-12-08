package part01

import (
	"github.com/zalgonoise/advent-2024/day04"
	"github.com/zalgonoise/advent-2024/is"
	"testing"
)

const target = "XMAS"

func TestPart01(t *testing.T) {
	for _, testcase := range []struct {
		name  string
		input string
		wants int
	}{
		{
			name: "Example",
			input: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			wants: 18,
		},
		{
			name:  "Challenge",
			input: day04.Input,
			wants: 2591,
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			grid := Parse(testcase.input)

			is.Equal(t, testcase.wants, Count(grid, target))
		})
	}
}

func TestBreakdownGrid(t *testing.T) {
	t.Run("Horizontal", func(t *testing.T) {
		for _, testcase := range []struct {
			name  string
			input string
			wants [][]byte
		}{
			{
				name: "OK/Simple",
				input: `XY
AB`,
				wants: [][]byte{[]byte("XY"), []byte("AB"), []byte("YX"), []byte("BA")},
			},
			{
				name: "OK/Simple",
				input: `XYZ
ABC
123`,
				wants: [][]byte{
					[]byte("XYZ"), []byte("ABC"), []byte("123"),
					[]byte("ZYX"), []byte("CBA"), []byte("321"),
				},
			},
		} {
			t.Run(testcase.name, func(t *testing.T) {
				grid := Parse(testcase.input)
				horizontal := breakdownHorizontal(grid)

				is.Equal(t, len(testcase.wants), len(horizontal))
				for i := range horizontal {
					is.Equal(t, string(testcase.wants[i]), string(horizontal[i]))
				}
			})
		}
	})

	t.Run("Vertical", func(t *testing.T) {
		for _, testcase := range []struct {
			name  string
			input string
			wants [][]byte
		}{
			{
				name: "OK/Simple",
				input: `XY
AB`,
				wants: [][]byte{[]byte("XA"), []byte("YB"), []byte("AX"), []byte("BY")},
			},
			{
				name: "OK/Simple",
				input: `XYZ
ABC
123`,
				wants: [][]byte{
					[]byte("XA1"), []byte("YB2"), []byte("ZC3"),
					[]byte("1AX"), []byte("2BY"), []byte("3CZ"),
				},
			},
		} {
			t.Run(testcase.name, func(t *testing.T) {
				grid := Parse(testcase.input)
				horizontal := breakdownVertical(grid)

				is.Equal(t, len(testcase.wants), len(horizontal))
				for i := range horizontal {
					is.Equal(t, string(testcase.wants[i]), string(horizontal[i]))
				}
			})
		}
	})

	t.Run("DiagonalTL", func(t *testing.T) {
		for _, testcase := range []struct {
			name  string
			input string
			wants [][]byte
		}{
			{
				name: "OK/Simple",
				input: `XYZ
ABC
123`,
				wants: [][]byte{
					[]byte("1"), []byte("A2"), []byte("XB3"), []byte("YC"), []byte("Z"),
					[]byte("1"), []byte("2A"), []byte("3BX"), []byte("CY"), []byte("Z"),
				},
			},
			{
				name: "OK/Simple",
				input: `XYZ0
ABCD
1234
5678`,
				wants: [][]byte{
					[]byte("5"), []byte("16"), []byte("A27"), []byte("XB38"), []byte("YC4"), []byte("ZD"), []byte("0"),
					[]byte("5"), []byte("61"), []byte("72A"), []byte("83BX"), []byte("4CY"), []byte("DZ"), []byte("0"),
				},
			},
		} {
			t.Run(testcase.name, func(t *testing.T) {
				grid := Parse(testcase.input)
				horizontal := breakdownDiagonalTopLeft(grid)

				is.Equal(t, len(testcase.wants), len(horizontal))
				for i := range horizontal {
					is.Equal(t, string(testcase.wants[i]), string(horizontal[i]))
				}
			})
		}
	})

	t.Run("DiagonalTR", func(t *testing.T) {
		for _, testcase := range []struct {
			name  string
			input string
			wants [][]byte
		}{
			{
				name: "OK/Simple",
				input: `XYZ
ABC
123`,
				wants: [][]byte{
					[]byte("X"), []byte("YA"), []byte("ZB1"), []byte("C2"), []byte("3"),
					[]byte("X"), []byte("AY"), []byte("1BZ"), []byte("2C"), []byte("3"),
				},
			},
			{
				name: "OK/Simple",
				input: `XYZ0
ABCD
1234
5678`,
				wants: [][]byte{
					[]byte("X"), []byte("YA"), []byte("ZB1"), []byte("0C25"), []byte("D36"), []byte("47"), []byte("8"),
					[]byte("X"), []byte("AY"), []byte("1BZ"), []byte("52C0"), []byte("63D"), []byte("74"), []byte("8"),
				},
			},
		} {
			t.Run(testcase.name, func(t *testing.T) {
				grid := Parse(testcase.input)
				horizontal := breakdownDiagonalTopRight(grid)

				is.Equal(t, len(testcase.wants), len(horizontal))
				for i := range horizontal {
					is.Equal(t, string(testcase.wants[i]), string(horizontal[i]))
				}
			})
		}
	})
}
