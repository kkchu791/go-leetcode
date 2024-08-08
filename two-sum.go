func twoSum(nums []int, target int) []int {
    differenceMap := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        difference := target - nums[i]
        if val, ok := differenceMap[nums[i]]; ok {
            return []int{val, i}
        }

        differenceMap[difference] = i
    }

    return nil
}
