def insertion_sort[T](array: list[T], comparator: callable[[T, T], bool] | None = None) -> list[T]:
    comparator = comparator or (lambda a, b: a - b)
    for i in range(1, len(array)):
        insert_val = array[i]
        j = i - 1
        while j >= 0 and comparator(array[j], insert_val) > 0:
            array[j + 1] = array[j]
            j -= 1
        array[j + 1] = insert_val
    return array