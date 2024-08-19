import "math"

func secondGreatestDivisor(n int) int {
    if n % 2 == 0 {
        return n / 2
    }

    for i := 3; i < int(math.Sqrt(float64(n))) + 1; i +=2 {
        if n % i == 0 {
            return n / i
        }
    }

    return 1
}
 
func minSteps(n int) int {
    steps := 0
    if n == 1 {
        return steps
    }
    sgtd := secondGreatestDivisor(n)
    return steps + n / sgtd + minSteps(sgtd)
}
