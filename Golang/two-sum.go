func twoSum(nums []int, target int) []int {
    tmp_map := make(map[int]int)

    for i, num := range nums {
        cord, ok := tmp_map[target - num]
        if ok {
            return []int{cord, i}
        } else {
            tmp_map[num] = i
        }
    }
    return []int{}
}