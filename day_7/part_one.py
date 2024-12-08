file = open("./input.txt", "r")
lines = file.read().splitlines()
file.close()


def rec(testVal, values, i) -> bool:
    val = values[i]
    if i == 0:
        return testVal / val == 1 or testVal - val == 0
    if testVal % val == 0:
        return rec(testVal / val, values, i - 1) or rec(testVal - val, values, i - 1)
    else:
        return rec(testVal - val, values, i - 1)


total = 0
for line in lines:
    [testVal, values] = line.split(": ")
    values = list(map(lambda num: int(num), values.split(" ")))
    testVal = int(testVal)
    if rec(testVal, values, len(values) - 1):
        total += testVal
print("Total is: ", total)
