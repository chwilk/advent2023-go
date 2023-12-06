package day05

import (
	"advent2023-go/helpers"
	"reflect"
	"strings"
	"testing"
)

const example string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := "35"
		got := Day05{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.CheckString(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := "0"
		got := Day05{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.CheckString(t, expect, got)
	})
}

func TestReadAlmanac(t *testing.T) {
	almanac := readAlmanac(strings.NewReader(example))

	t.Run("Check seeds list", func(t *testing.T) {
		expect := []int{79, 14, 55, 13}
		got := almanac.seeds
		if ! reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})
	t.Run("Check seed-to-soil map", func(t *testing.T) {
		expect := [3]int{98, 100, -48}
		got := almanac.maps[0][0]
		if ! reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})
	t.Run("Check location map", func(t *testing.T) {
		expect := [3]int{93, 97, -37}
		got := almanac.maps[6][1]
		if ! reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v, got %v", expect, got)
		}
	})
}

func TestFollowMap(t *testing.T) {
	almanac := readAlmanac(strings.NewReader(example))
	expect := 82
	got := followMap(almanac, 79)
	helpers.Check(t, expect, got)
}