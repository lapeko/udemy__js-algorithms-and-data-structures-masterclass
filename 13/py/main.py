def countZeroes(array):
    left = 0
    right = len(array) - 1
    middle = 0

    while left <= right:
        middle = (left + right) // 2

        if array[middle] == 0:
            right = middle - 1
        else:
            left = middle + 1

    if array[middle] == 1:
        middle += 1

    return len(array) - middle
