from main import power

def test():
    assert power(2, 0) == 1
    assert power(2, 2) == 4
    assert power(2, 4) == 16


if __name__ == "__main__":
    test()
    print("OK")