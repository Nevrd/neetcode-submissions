func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left+(right-left)/2
		guess := nums[mid]
		if guess == target {
			return mid
		} else if guess < target {
			left = mid + 1
		} else {
			right = mid - 1 
		}
	}
	return -1
}
