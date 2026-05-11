def selectionSort[T](array: list[T], comparator: callable[[T, T], bool] | None = None) -> list[T]:
    comparator = comparator or (lambda a, b: a - b)
    for i in range(len(array) - 1):
        minIdx = i
        for j in range(i + 1, len(array)):
            if comparator(array[j], array[minIdx]) < 0:
                minIdx = j
        if minIdx != i:
            array[i], array[minIdx] = array[minIdx], array[i]
    return array
