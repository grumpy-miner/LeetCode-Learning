from math import sqrt

class Solution:

    def second_greatest_divisor(self, n: int) -> int:
        # if number is even then greatest is n // 2
        if n % 2 == 0:
            return n // 2

        for i in range(3, int(sqrt(n)) + 1, 2):
            if n % i == 0:
                return n // i

        return 1

    def minSteps(self, n: int) -> int:
        steps = 0
        # If n == 1 we don't need to do any operation
        if n == 1:
            return steps
        sgtd = self.second_greatest_divisor(n)
        # f(n) = f(gtd) + n / gtd (1 copy operations + (n / gtd - 1) pastes)
        return steps + n // sgtd + self.minSteps(sgtd)
