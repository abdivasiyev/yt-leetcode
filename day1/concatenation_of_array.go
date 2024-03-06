package day1

func getConcatenation(nums []int) []int {
	var ans = make([]int, len(nums)*2)

	for i := 0; i < len(nums); i++ {
		ans[i], ans[i+len(nums)] = nums[i], nums[i]
	}

	return ans
}
