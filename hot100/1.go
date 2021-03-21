package leetcode

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		if p, ok := hashMap[target-num]; ok {
			return []int{p, i}
		}
		hashMap[num] = i
	}
	return nil
}
