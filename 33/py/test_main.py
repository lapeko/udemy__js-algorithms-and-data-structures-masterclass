from main import selectionSort as sortFn
from typing import TypedDict

class KittyData(TypedDict):
    name: str
    age: int

moar_kitty_data: list[KittyData] = [{
    "name": "LilBub",
    "age":  7,
}, {
    "name": "Garfield",
    "age":  40,
}, {
    "name": "Heathcliff",
    "age":  45,
}, {
    "name": "Blue",
    "age":  1,
}, {
    "name": "Grumpy",
    "age":  6,
}]

expected_moar_kittyData: list[KittyData] = [{
    "name": "Heathcliff",
    "age":  45,
}, {
    "name": "Garfield",
    "age":  40,
}, {
    "name": "LilBub",
    "age":  7,
}, {
    "name": "Grumpy",
    "age":  6,
}, {
    "name": "Blue",
    "age":  1,
}]

def str_comp(a: str, b: str) -> int:
    if a < b:
        return -1
    elif a > b:
        return 1
    return 0
	

def oldest_to_youngest(a: KittyData, b: KittyData) -> int:
    return b["age"] - a["age"]


def test():
    assert sortFn([4, 20, 12, 10, 7, 9]) == [4, 7, 9, 10, 12, 20]
    assert sortFn([0, -10, 7, 4]) == [-10, 0, 4, 7]
    assert sortFn([1, 2, 3]) == [1, 2, 3]
    assert sortFn([]) == []
    assert sortFn([4, 3, 5, 3, 43, 232, 4, 34, 232, 32, 4, 35, 34, 23, 2, 453, 546, 75, 67, 4342, 32]) == [2, 3, 3, 4, 4, 4, 5, 23, 32, 32, 34, 34, 35, 43, 67, 75, 232, 232, 453, 546, 4342]
    assert sortFn(["LilBub", "Garfield", "Heathcliff", "Blue", "Grumpy"], str_comp) == ["Blue", "Garfield", "Grumpy", "Heathcliff", "LilBub"]
    assert sortFn(moar_kitty_data, oldest_to_youngest) == expected_moar_kittyData


if __name__ == "__main__":
    test()
    print("OK")