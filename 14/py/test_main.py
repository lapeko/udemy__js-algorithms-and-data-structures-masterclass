from main import sorted_frequency


def test():
    assert sorted_frequency([1, 1, 2, 2, 2, 2, 3], 2) == 4
    assert sorted_frequency([1, 1, 2, 2, 2, 2, 3], 3) == 1
    assert sorted_frequency([1, 1, 2, 2, 2, 2, 3], 1) == 2
    assert sorted_frequency([1, 1, 2, 2, 2, 2, 3], 4 == -1)


if __name__ == "__main__":
    test()
    print("OK")
