def bubbleSort[T](array: list[T], comparator: callable[[T, T], int] | None = None):
    comparator = comparator or (lambda a, b: a - b)
    for i in range(1, len(array)):
        swap = False
        for j in range(len(array) - i):
            if comparator(array[j], array[j + 1]) > 0:
                swap = True
                array[j], array[j + 1] = array[j + 1], array[j]
        if not swap:
            break
    return array
