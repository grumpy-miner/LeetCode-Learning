class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s

        strings = [""] * numRows
        i = 0
        incremental = 1
        for character in s:
            if i == 0:
                incremental = 1
            elif i == numRows - 1:
                incremental = -1
            strings[i] += character
            i += incremental

        return "".join(strings)
