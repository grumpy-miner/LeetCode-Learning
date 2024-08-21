import "slices"

func minCostClimbingStairs(cost []int) int {
    // clone `cost`` so we don't modify `cost``
    costsClone := slices.Clone(cost)
    for i := len(costsClone) - 3; i >= 0 ; i-- {
        costsClone[i] = costsClone[i] + min(costsClone[i + 1], costsClone[i + 2])
    }
    return min(costsClone[0], costsClone[1])
}
