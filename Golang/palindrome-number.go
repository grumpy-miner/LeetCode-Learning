func isPalindrome(x int) bool {
    reversed := 0;
    for i := x; i > 0; i /= 10 {
        reversed *= 10
        reversed += i % 10
    }

    return reversed == x
}
