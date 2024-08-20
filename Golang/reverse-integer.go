import "math"
func reverse(x int) int {

    reversed_x := 0
    for x != 0 {
        reversed_x = reversed_x * 10 + x % 10
        x /= 10
    }

    if math.MinInt32 > reversed_x || reversed_x > math.MaxInt32 {
        return 0
    } else {
        return reversed_x
    }
}
