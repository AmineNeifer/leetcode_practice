func majorityElement(nums []int) int {
    dict := make(map[int]int)
	for _, v := range nums {
		dict[v]++
	}   
	for k, v := range dict {
		if v > len(nums)/2 {
			return k
		}
	}
    return nums[0]
}
