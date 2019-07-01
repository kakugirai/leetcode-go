from typing import List


class Solution:
    @staticmethod
    def searchMatrix(matrix: List[List[int]], target: int) -> bool:
        """
        :param matrix: a 2D integers list
        :param target: an integer
        :return: a boolean, indicate whether matrix contains target
        """
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return False
        n, m = len(matrix), len(matrix[0])
        start, end = 0, n * m - 1
        while start + 1 < end:
            mid = (start + end) // 2
            x, y = mid // m, mid % m
            if matrix[x][y] < target:
                start = mid
            else:
                end = mid
        x, y = start // m, start % m
        if matrix[x][y] == target:
            return True

        x, y = end // m, end % m
        if matrix[x][y] == target:
            return True
        return False


if __name__ == '__main__':
    mySolution = Solution()
    # my_matrix = [[]]
    my_matrix = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 50]]
    my_target = 3
    my_result = mySolution.searchMatrix(my_matrix, my_target)
    print(my_result)
