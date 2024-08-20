import (
    "strings"
)

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    stringsArray := make([]string, numRows)
    i := 0
    incremental := 1
    for _, char := range s {
        stringsArray[i] += string(char)
        if i == 0 {
            incremental = 1
        } else if i == numRows - 1 {
            incremental = -1
        }
        i += incremental
    }
    return strings.Join(stringsArray, "")
}
