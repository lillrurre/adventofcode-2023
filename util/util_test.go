package util

import (
	"reflect"
	"testing"
)

func TestFlipGrid(t *testing.T) {
	testCases := []struct {
		in    [][]int
		out   [][]int
		flips int
	}{
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			out: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
			flips: 0,
		},
		{
			in: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
			out: [][]int{
				{16, 11, 6, 1},
				{17, 12, 7, 2},
				{18, 13, 8, 3},
				{19, 14, 9, 4},
				{20, 15, 10, 5},
			},
			flips: 1,
		},
		{
			in: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			out: [][]int{
				{12, 11, 10, 9},
				{8, 7, 6, 5},
				{4, 3, 2, 1},
			},
			flips: 2,
		},
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			out: [][]int{
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			},
			flips: 3,
		},
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			out: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			flips: 4,
		},
	}

	for _, testCase := range testCases {
		out := FlipGrid(testCase.in, testCase.flips)
		if !reflect.DeepEqual(out, testCase.out) {
			t.Fatal("lol")
		}
	}
}
