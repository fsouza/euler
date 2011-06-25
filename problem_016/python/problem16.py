number = 2 ** 1000
digits = str(number)
print reduce(lambda x, y: int(x) + int(y), digits)
