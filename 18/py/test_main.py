from main import productOfArray


def test():
    assert productOfArray([1,2,3]) == 6
    assert productOfArray([1,2,3,10]) == 60


if __name__ == "__main__":
    test()
    print("OK")