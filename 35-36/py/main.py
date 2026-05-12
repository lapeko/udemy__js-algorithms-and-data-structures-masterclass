from typing import Callable

def merge_sort[T](array: list[T], comparator: Callable[[T, T], int] | None = None) -> list[T]:
    comparator = comparator or (lambda a, b: a - b)

    def fn(array):
        if len(array) <= 1:
            return array
        
        middle = len(array) // 2

        return merge(
            fn(array[:middle]),
            fn(array[middle:]),
            comparator,
        )
     
    return fn(array)


def merge[T](a: list[T], b: list[T], comparator: Callable[[T, T], int]) -> list[T]:
    l, r, lr, size = 0, 0, 0, len(a) + len(b)
    c = [None] * size
    
    while lr < size:
        if l >= len(a):
            c[lr] = b[r]
            r += 1
        elif r >= len(b):
            c[lr] = a[l]
            l += 1
        elif comparator(a[l], b[r]) < 0:
            c[lr] = a[l]
            l += 1
        else:
            c[lr] = b[r]
            r += 1
        lr += 1
    
    return c
