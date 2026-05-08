def productOfArray(array: list[int]) -> int:
    return array[0] * productOfArray(array[1:]) if len(array) != 1 else array[0]