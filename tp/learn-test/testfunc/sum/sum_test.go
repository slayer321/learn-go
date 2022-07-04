package main

import "testing"

func TestInts(t *testing.T) {

	tt := []struct {
		numbers []int
		sum     int
	}{
		{[]int{1, 2, 3}, 6},
		{nil, 0},
		{[]int{1, -1}, 0},
	}

	for _, tc := range tt {
		s := Ints(tc.numbers...)
		if s != tc.sum {
			t.Errorf("got %v , want %v", s, tc.sum)
		}
	}

}
