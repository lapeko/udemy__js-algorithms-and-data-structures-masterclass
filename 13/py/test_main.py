from main import countZeroes


def test():
    assert countZeroes([1, 1, 1, 1, 0, 0]) == 2
    assert countZeroes([1, 0, 0, 0, 0]) == 4
    assert countZeroes([0, 0, 0]) == 3
    assert countZeroes([1, 1, 1, 1]) == 0


if __name__ == "__main__":
    test()
    print("OK")
