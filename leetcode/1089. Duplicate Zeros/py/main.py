def duplicate_zeros(arr: list[int]) -> list[int]:
    r, l, size = 0, 0, len(arr)
    for i in range(size):
        r += 2 if arr[i] == 0 else 1
        if r >= size:
            l = i
            break
    if r > size:
        arr[-1] = 0
        r -= 2
        l -= 1
    r -= 1
    while l < r:
        arr[r] = arr[l]
        if arr[l] == 0:
            r -= 1
            arr[r] = arr[l]
        r -= 1
        l -= 1
    return arr


assert duplicate_zeros([1, 0, 2, 3, 5, 4]) == [1, 0, 0, 2, 3, 5]
assert duplicate_zeros([0, 1, 0, 2, 3, 4]) == [0, 0, 1, 0, 0, 2]
assert duplicate_zeros([1, 0, 2, 3, 0, 4]) == [1, 0, 0, 2, 3, 0]
assert duplicate_zeros([1, 0, 0, 2, 3, 0, 4, 5]) == [1, 0, 0, 0, 0, 2, 3, 0]
assert duplicate_zeros([]) == []
assert duplicate_zeros([0]) == [0]
assert duplicate_zeros([1]) == [1]
assert duplicate_zeros([1, 2, 3]) == [1, 2, 3]
print("Ok")