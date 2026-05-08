def power(base: int, pow: int) -> int:
    return 1 if pow <= 0 else base * power(base, pow - 1)