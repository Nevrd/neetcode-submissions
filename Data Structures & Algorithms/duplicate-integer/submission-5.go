func hasDuplicate(nums []int) bool {
	countNums := make(map[int]int)
	for _, v := range nums {
		countNums[v]++
	}

	for _, v := range countNums{
		if v > 1 {
			return true
		}
	}
	return false
}
