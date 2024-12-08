import math

file = open("./input.txt", "r")
lines = file.read().splitlines()
file.close()


def concat(a, b):
    return int(str(a) + str(b))

def recCat(testVal, current, values, i) -> bool:
    if i > len(values):
        return False
    if i == len(values):
        return testVal == current
    ret = False
    ret = ret or recCat(testVal, current * values[i], values, i + 1)
    ret = ret or recCat(testVal, current + values[i], values, i + 1)
    ret = ret or recCat(testVal, concat(current, values[i]), values, i + 1)
    return ret

total = 0
for line in lines:
    [testVal, values] = line.split(": ")
    values = list(map(lambda num: int(num), values.split(" ")))
    testVal = int(testVal)
    if recCat(testVal, values[0], values, 1):
        total += testVal
print("Total is: ", total)
