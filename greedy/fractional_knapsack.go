package greedy

import "sort"

type Item struct {
	weight int
	value  int
}

func fractionalKnapsack(wgt, val []int, cap int) float64 {
	items := make([]Item, len(wgt))
	for i := range items {
		items[i] = Item{wgt[i], val[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return float64(items[i].value)/float64(items[i].weight) > float64(items[j].value)/float64(items[j].weight)
	})

	res := 0.0
	for _, item := range items {
		if item.weight <= cap {
			res += float64(item.value)
			cap -= item.weight
		} else {
			res += float64(item.value) / float64(item.weight) * float64(cap)
			break
		}
	}
	return res
}
