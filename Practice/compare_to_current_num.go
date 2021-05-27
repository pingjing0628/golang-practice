package main
func smallerNumbersThanCurrent(nums []int) []int {
	output := make([]int, len(nums))
	count := 0
	for i := 0; i < len(nums); i++{
		count = compare(nums, i)
		output[i] = count
	}
	return output
}

func compare(nums []int, index int) int {
	count := 0
	for i := 0; i < len(nums); i++{
		if nums[i] < nums[index] {
			count += 1
		}
	}
	return count
}
