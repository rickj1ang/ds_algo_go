package dynamicprogramming

func maxCakePrice(priceList []int, weight int) int {
	if weight <= 1 {
		return priceList[weight]
	}

	res := 0
	for i := 0; i < weight; i++ {
		res = max(res, maxCakePrice(priceList, weight)+priceList[weight-i])
	}
	return res
}
