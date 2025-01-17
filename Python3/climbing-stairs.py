from functools import lru_cache


class Solution:
    @lru_cache(maxsize=None)
    def climbStairs(self, n: int) -> int:
        if n <= 2:
            return n
        return self.climbStairs(n - 1) + self.climbStairs(n - 2)
