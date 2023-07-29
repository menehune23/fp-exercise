package fp

import (
	"strconv"

	"github.com/samber/lo"
)

func ComputeTotal(list []string) int {
	ints := toInts(list)
	return sum(ints)

	// Or even shorter, without the helpers:
	//return lo.Sum(
	//	lo.Map(list, func(str string, _ int) int {
	//		val, _ := lo.TryOr(func() (int, error) { return strconv.Atoi(str) }, 0)
	//		return val
	//	}),
	//)
}

func toInts(list []string) []int {
	return lo.Map(list, func(str string, _ int) int {
		val, _ := lo.TryOr(func() (int, error) { return strconv.Atoi(str) }, 0)
		return val
	})
}

func sum(ints []int) int {
	return lo.Reduce(ints, func(acc int, val int, _ int) int {
		return acc + val
	}, 0)
}
