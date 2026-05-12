def quick_sort[T](array: list[T], comparator: callable[[T, T], T] | None = None) -> list[T]:
    if len(array) <= 1:
        return array
    
    # comparator = comparator or (lambda a, b: a - b)
    
    pivot_idx = partition(array)

    return quick_sort(array[:pivot_idx]) + [array[pivot_idx]] + quick_sort(array[pivot_idx + 1:])

    

def partition(a):
    pivot = a[0]

    l, r = 1, len(a) - 1

    while True:
        while True:
            if a[l] >= pivot:
                break
            if l + 1 >= r:
                break
            l += 1

        while True:
            if a[r] < pivot:
                break
            if l + 1 >= r:
                break
            r -= 1

        if a[l] >= pivot > a[r]:
            a[l], a[r] = a[r], a[l]
            
        if l + 2 >= r:
            break
        else:
            l += 1
            r -= 1

    idx = 0
    for i in range(l, r + 1):
        if pivot > a[i]:
            idx = i
    
    a[0], a[idx] = a[idx], a[0]
    return idx
            

    

def comparator(a: int, b: int) -> int:
    return a - b

print(quick_sort([5, 5, 5, 2, 1]))
print(quick_sort([2, 4, 1, 6, 7, 3]))
print(quick_sort([5, 4, 1, 6, 7, 3]))
print(quick_sort([5, 4, 6, 1, 3, 7]))
print(quick_sort([5, 4, 6, 33, 1, 3, 7, 33, -1, 34, -1, -1]))
print(quick_sort([3,2,1]))
print(quick_sort([1,2,3]))
print(quick_sort([2,1]))
print(quick_sort([2,3,1]))
print(quick_sort([3,3,3]))
print(quick_sort([4,1,4,1]))
print(quick_sort([10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5]))
