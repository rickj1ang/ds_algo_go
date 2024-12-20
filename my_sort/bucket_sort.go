package mysort

import "sort"

func bucketSort(nums []int) {
	// init bucket
	bucketNum := len(nums) / 1
	buckets := make([][]int, bucketNum)
	for i := 0; i < bucketNum; i++ {
		buckets[i] = make([]int, 0)
	}

	//put nums into bucket
	for _, num := range nums {
		i := num / 10
		buckets[i] = append(buckets[i], num)
	}

	//sort every bucket
	for i := 0; i < bucketNum; i++ {
		sort.Ints(buckets[i])
	}

	//merge the buckets
	i := 0
	for _, bucket := range buckets {
		for _, num := range bucket {
			nums[i] = num
			i++
		}
	}
}
