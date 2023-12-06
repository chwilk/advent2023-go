package day05

import (
	"advent2023-go/helpers"
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Day05 struct {
}

type Almanac struct {
	seeds []int
	maps  [7][][3]int
}

func (f Day05) Run(input io.Reader, part int) (result string) {
	switch part {
	case helpers.PartA:
		result = fmt.Sprint(PartA(input))
	default:
		result = fmt.Sprint(PartB(input))
	}
	return
}

func PartA(lines io.Reader) int {
	almanac := readAlmanac(lines)
	locations := make([]int, len(almanac.seeds))
	for i := range almanac.seeds {
		locations[i] = followMap(almanac, almanac.seeds[i])
	}
	sort.Ints(locations)
	return locations[0]
}

func PartB(lines io.Reader) (answer int) {
	almanac := readAlmanac(lines)
	answer += almanac.seeds[0]
	return
}

func readAlmanac(input io.Reader) (almanac Almanac) {

	scanner := bufio.NewScanner(input)

	mapIndex := -1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "seeds:":
			for _, num := range tokens[1:] {
				tmp, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				almanac.seeds = append(almanac.seeds, tmp)
			}
		case "seed-to-soil":
			fallthrough
		case "soil-to-fertilizer":
			fallthrough
		case "fertilizer-to-water":
			fallthrough
		case "water-to-light":
			fallthrough
		case "light-to-temperature":
			fallthrough
		case "temperature-to-humidity":
			fallthrough
		case "humidity-to-location":
			mapIndex++
		default:
			var tmp [3]int
			var err error
			for i := range tokens {
				tmp[i], err = strconv.Atoi(tokens[i])
				if err != nil {
					panic(err)
				}
			}
			tmp[0], tmp[1], tmp[2] = tmp[1], tmp[1]+tmp[2], tmp[0]-tmp[1]
			almanac.maps[mapIndex] = append(almanac.maps[mapIndex], tmp)
		}
	}
	return
}

func followMap(almanac Almanac, seed int) (answer int) {
	answer = seed
	for _, mappings := range almanac.maps {
		for _, m := range mappings {
			if answer >= m[0] && answer <= m[1] {
				answer += m[2]
				break
			}
		}
	}
	return
}
