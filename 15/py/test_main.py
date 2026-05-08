from main import findRotatedIndex


def test():
    assert findRotatedIndex([3, 4, 1, 2], 4) == 1
    assert findRotatedIndex([6, 7, 8, 9, 1, 2, 3, 4], 8) == 2
    assert findRotatedIndex([6, 7, 8, 9, 1, 2, 3, 4], 3) == 6
    assert findRotatedIndex([37, 44, 66, 102, 10, 22], 14) == -1
    assert findRotatedIndex([6, 7, 8, 9, 1, 2, 3, 4], 12) == -1
    assert findRotatedIndex([11, 12, 13, 14, 15, 16, 3, 5, 7, 9], 16) == 5



if __name__ == "__main__":
    test()
    print("OK")