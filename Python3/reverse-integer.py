class Solution:
    def reverse(self, x: int) -> int:
        multiplier = -1 if x < 0 else 1

        x *= multiplier
        reversed_x = 0
        while x != 0:
            reversed_x = reversed_x * 10 + x % 10
            x //= 10

            if reversed_x >= 2**31:
                return 0

        return reversed_x * multiplier
