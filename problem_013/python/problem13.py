def sum_digits_of(number):
    return reduce(lambda x, y: int(x) + int(y), str(number))
