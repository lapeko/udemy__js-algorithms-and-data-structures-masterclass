from typing import Callable

def quick_sort[T](a: list[T], comparator: Callable[[T, T], int] | None = None) -> list[T]:
    comparator = comparator or (lambda a, b: a - b)

    def _quick_sort(low: int, high: int):
        if low < high:
            i = partition(a, low, high, comparator)
            _quick_sort(low, i - 1)
            _quick_sort(i + 1, high)

    _quick_sort(0, len(a) - 1)

    return a


# TODO Required
def partition[T](a: list[T], low: int, high: int, comparator: Callable[[T, T], int]) -> int:
    pivot = a[high]

    for j in range(low, high):
        if comparator(a[j], pivot) < 0:
            a[low], a[j] = a[j], a[low]
            low += 1

    a[low], a[high] = a[high], a[low]

    return low



    

# def comparator(a: int, b: int) -> int:
#     return a - b

# print(quick_sort([5, 5, 5, 2, 1]))
# print(quick_sort([2, 4, 1, 6, 7, 3]))
# print(quick_sort([5, 4, 1, 6, 7, 3]))
# print(quick_sort([5, 4, 6, 1, 3, 7]))
# print(quick_sort([5, 4, 6, 33, 1, 3, 7, 33, -1, 34, -1, -1]))
# print(quick_sort([3, 2, 1]))
# print(quick_sort([1, 2, 3]))
# print(quick_sort([2, 1]))
# print(quick_sort([2, 3, 1]))
# print(quick_sort([3, 3, 3]))
# print(quick_sort([4, 1, 4, 1]))
# print(quick_sort([10, 11, 12, 13, 14, 15, 1, 2, 3, 4, 5]))


# def partition(a):
#     pivot = a[0]

#     l, r = 1, len(a) - 1

#     while True:
#         while True:
#             if a[l] >= pivot:
#                 break
#             if l + 1 >= r:
#                 break
#             l += 1

#         while True:
#             if a[r] < pivot:
#                 break
#             if l + 1 >= r:
#                 break
#             r -= 1

#         if a[l] >= pivot > a[r]:
#             a[l], a[r] = a[r], a[l]
            
#         if l + 2 >= r:
#             break
#         else:
#             l += 1
#             r -= 1

#     idx = 0
#     for i in range(l, r + 1):
#         if pivot > a[i]:
#             idx = i
    
#     a[0], a[idx] = a[idx], a[0]
#     return idx