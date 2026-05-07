def sorted_frequency(array: list[int], search: int) -> int:
    min = find_min_idx(array, search)

    if min == -1:
        return -1

    return find_max_idx(array, search) - min + 1


def find_min_idx(array: list[int], search: int) -> int:
    if array[0] == search:
        return 0

    left, right, res = 0, len(array) - 1, -1
    while left <= right:
        middle = (left + right) // 2
        if array[middle] == search:
            res = middle
        if array[middle] >= search:
            right = middle - 1
        else:
            left = middle + 1

    return res


def find_max_idx(array: list[int], search: int) -> int:
    if array[-1] == search:
        return len(array) - 1

    left, right, res = 0, len(array) - 1, -1
    while left <= right:
        middle = (left + right) // 2
        if array[middle] == search:
            res = middle
        if array[middle] <= search:
            left = middle + 1
        else:
            right = middle - 1

    return res
