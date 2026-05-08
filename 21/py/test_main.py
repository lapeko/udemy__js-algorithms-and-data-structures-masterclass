from main import reverse


def test():
    assert reverse("awesome") == "emosewa"
    assert reverse("rithmschool") == "loohcsmhtir"


if __name__ == "__main__":
    test()
    print("OK")