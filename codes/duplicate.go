package interview

func containsDuplicate(nums []int) bool {
	cache := make(map[int]int)
	for i, n := range nums {
		if _, ok := cache[n]; ok {
			return false
		}
		cache[n] = i
	}
	return true
}
