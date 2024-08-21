import re

class Solution:
    def myAtoi(self, s: str) -> int:
        pattern = r'^\s*([-+]?\d+)'
        match = re.search(pattern, s)
        if match:
            return min(
                max(
                    int(match.group(1)),
                    -2147483648
                ),
                2147483647
            )
        else:
            return 0
