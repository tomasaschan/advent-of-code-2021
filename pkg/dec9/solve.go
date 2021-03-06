package dec9

import (
	"container/list"
	"sort"

	"github.com/tomasaschan/advent-of-code-2021/pkg/utils/twod"
)

func A(input string) int {
	m := twod.IntMapFromString(input, twod.Ints())

	totalRisk := 0
	for _, sink := range sinks(m) {
		totalRisk += sink.x + 1
	}

	return totalRisk
}

func B(input string) int {
	m := twod.IntMapFromString(input, twod.Ints())
	basins := []data{}

	for _, sink := range sinks(m) {
		basins = append(basins, data{sink.p, sizeOfBasin(m, sink.p)})
	}

	sort.Slice(basins, func(i, j int) bool { return basins[i].x > basins[j].x })

	threeLargest := 1
	for i := 0; i < 3; i++ {
		threeLargest *= basins[i].x
	}
	return threeLargest
}

type data struct {
	p twod.Vector
	x int
}

func sinks(m twod.IntMap) []data {
	upperLeft, bottomRight := m.Corners()
	results := make([]data, 0)
	for x := upperLeft.X; x <= bottomRight.X; x++ {
		for y := upperLeft.Y; y <= bottomRight.Y; y++ {
			here := twod.Vector{X: x, Y: y}
			height := *m.At(here)
			isSink := true
			for _, p := range m.NeighborsOf(here) {
				there := *m.At(p)
				if height >= there {
					isSink = false
					break
				}
			}
			if isSink {
				results = append(results, data{here, height})
			}
		}
	}
	return results
}

func sizeOfBasin(m twod.IntMap, p twod.Vector) int {
	q := list.List{}
	q.PushBack(p)
	size := 0
	seen := map[twod.Vector]bool{}

	for {
		front := q.Front()
		q.Remove(front)
		here, ok := front.Value.(twod.Vector)
		if !ok {
			break
		}
		for _, x := range m.NeighborsOf(here) {
			if _, ok := seen[x]; ok {
				continue
			}
			h := *m.At(x)
			if h < 9 {
				q.PushBack(x)
				size += 1
				seen[x] = true
			}
		}
		if q.Len() == 0 {
			break
		}
	}
	return size
}
