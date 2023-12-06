package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	alm := parseInput(input)
	locations := make([]int, len(alm.Seeds))
	for idx, seed := range alm.Seeds {
		locations[idx] = alm.HumidityToLoc.Get(
			alm.TempToHumidity.Get(
				alm.LightToTemp.Get(
					alm.WaterToLight.Get(
						alm.FertilizerToWater.Get(
							alm.SoilToFertilizer.Get(alm.SeedToSoil.Get(seed)))))))
	}
	return slices.Min(locations)
}

func Part2(input string) (result int) {
	alm := parseInput(input)
	_, searchMax := alm.HumidityToLoc.RangeDest()
	for loc := 0; loc < searchMax; loc++ {
		seed := alm.SeedToSoil.ReverseGet(
			alm.SoilToFertilizer.ReverseGet(
				alm.FertilizerToWater.ReverseGet(
					alm.WaterToLight.ReverseGet(
						alm.LightToTemp.ReverseGet(
							alm.TempToHumidity.ReverseGet(
								alm.HumidityToLoc.ReverseGet(loc)))))))
		if alm.ContainsSeed(seed) {
			return loc
		}
	}
	return
}

func parseInput(input string) (res Almanac) {
	scan := bufio.NewScanner(strings.NewReader(input))
	if scan.Scan() {
		for _, seedStr := range strings.Split(strings.Split(scan.Text(), ":")[1], " ") {
			seed, err := strconv.Atoi(seedStr)
			if err == nil {
				res.Seeds = append(res.Seeds, seed)
			}
		}
	}
	scan.Scan()
	res.SeedToSoil = parseMap(scan)
	res.SoilToFertilizer = parseMap(scan)
	res.FertilizerToWater = parseMap(scan)
	res.WaterToLight = parseMap(scan)
	res.LightToTemp = parseMap(scan)
	res.TempToHumidity = parseMap(scan)
	res.HumidityToLoc = parseMap(scan)
	return res
}

func parseMap(scan *bufio.Scanner) (res RangeMap) {
	scan.Scan()
	for scan.Scan() {
		rangeParts := strings.Split(strings.TrimSpace(scan.Text()), " ")
		if len(rangeParts) != 3 {
			break
		}
		entry := RangeMapEntry{}
		entry.DestRangeStart, _ = strconv.Atoi(rangeParts[0])
		entry.SrcRangeStart, _ = strconv.Atoi(rangeParts[1])
		entry.RangeN, _ = strconv.Atoi(rangeParts[2])
		res = append(res, entry)
	}
	slices.SortFunc(res, func(j, k RangeMapEntry) int {
		return j.DestRangeStart - k.DestRangeStart
	})
	return res
}

type RangeMap []RangeMapEntry

type RangeMapEntry struct {
	DestRangeStart int
	SrcRangeStart  int
	RangeN         int
}

func (rm RangeMap) Get(id int) int {
	for _, r := range rm {
		if id >= r.SrcRangeStart && id < (r.SrcRangeStart+r.RangeN) {
			return r.DestRangeStart + (id - r.SrcRangeStart)
		}
	}
	return id
}

func (rm RangeMap) ReverseGet(destID int) int {
	for _, r := range rm {
		if destID >= r.DestRangeStart && destID < (r.DestRangeStart+r.RangeN) {
			return r.SrcRangeStart + (destID - r.DestRangeStart)
		}
	}
	return destID
}

func (rm RangeMap) RangeDest() (min, max int) {
	mins, maxes := make([]int, 0), make([]int, 0)
	for _, r := range rm {
		mins = append(mins, r.DestRangeStart)
		maxes = append(maxes, r.DestRangeStart+r.RangeN)
	}
	return slices.Min(mins), slices.Max(maxes)
}

type Almanac struct {
	Seeds             []int
	SeedToSoil        RangeMap
	SoilToFertilizer  RangeMap
	FertilizerToWater RangeMap
	WaterToLight      RangeMap
	LightToTemp       RangeMap
	TempToHumidity    RangeMap
	HumidityToLoc     RangeMap
}

func (alm Almanac) ContainsSeed(seed int) bool {
	for idx := 0; idx < len(alm.Seeds); idx += 2 {
		start, num := alm.Seeds[idx], alm.Seeds[idx+1]
		if seed >= start && seed <= start+num {
			return true
		}
	}
	return false
}

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
