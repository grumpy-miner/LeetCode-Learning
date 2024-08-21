import (
	"strconv"
	"math"
	"regexp"
    "strings"
)

func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    pattern := `^([-+]?\d+)`
    regex := *regexp.MustCompile(pattern)
    match := regex.Find([]byte(s))
    matchAtoi, _ := strconv.Atoi(string(match[:]))

    return min(max(matchAtoi, math.MinInt32), math.MaxInt32)
}
