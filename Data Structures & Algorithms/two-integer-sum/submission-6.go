func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if index, ok := maps[diff]; ok {
			return []int{index, i}
		}
		maps[v] = i
	}
	return nil
}