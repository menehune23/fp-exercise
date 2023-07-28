package fp

import (
	"strconv"
)

// Can `ComputeTotal` be rewritten using functional programming concepts?

func ComputeTotal(list []string) int {
	var (
		ints  []int
		total = 0
	)

	for _, str := range list {
		val, err := strconv.Atoi(str)
		if err == nil {
			ints = append(ints, val)
		}
	}

	for _, val := range ints {
		total += val
	}

	return total
}
