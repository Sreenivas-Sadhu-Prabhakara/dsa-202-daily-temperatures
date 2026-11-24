package main

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
		{[]int{90, 60, 30}, []int{0, 0, 0}},
		{[]int{70, 70, 70}, []int{0, 0, 0}}, // strictly warmer: equal doesn't count
		{[]int{}, []int{}},
		{[]int{50}, []int{0}},
	}
	for _, c := range cases {
		if got := DailyTemperatures(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("DailyTemperatures(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
