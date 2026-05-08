def reverse(word: str) -> str:
    if len(word) == 1:
        return word

    return word[-1] + reverse(word[:-1])
