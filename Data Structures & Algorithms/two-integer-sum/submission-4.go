func twoSum(nums []int, target int) []int {
	count := make(map[int]int, len(nums))
	for i, v := range nums {
		need := target - v 
		if k, ok := count[need]; ok {
			return []int{k, i}
		}
		count[v] = i
	}
	return nil
}