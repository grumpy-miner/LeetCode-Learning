class Solution:
    def isPalindrome(self, x: int) -> bool:
        reversed_x = 0
        x_copy = x
        while x_copy > 0:
            reversed_x *= 10
            reversed_x += x_copy % 10
            x_copy //= 10

        return x == reversed_x
