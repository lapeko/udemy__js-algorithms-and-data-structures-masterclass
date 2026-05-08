import math

def findRotatedIndex(array: list[int], target: int) -> int:
    left, right = 0, len(array) - 1
    
    while left <= right:
        middle = math.floor((left + right) / 2)
        
        if target == array[middle]:
            return middle
        
        if array[left] <= array[middle]:
            if target > array[left] and target < array[middle]:
                right = middle - 1 
            else:
                left = middle + 1
        else:
            if target > array[middle] and target < array[right]:
                left = middle + 1
            else:
                right = middle - 1

    return -1
