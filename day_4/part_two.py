file = open("input.txt", "r")
lines = file.read().splitlines()
file.close()

width = len(lines[0])
height = len(lines)

def inBox(x, y):
    return x >= 0 and y >= 0 and x < width and y < height

def patternFits(x, y):
    for i in range(-1, 2, 2):
        for j in range(-1, 2, 2):
            if not (inBox(x + i, y + j) and lines[y + j][x + i] in ["M", "S"]):
                return False
    return True

def checkIfPattern(x, y):
    if not patternFits(x, y):
        return False
    topLeft = lines[y + 1][x - 1]
    topRight = lines[y + 1][x + 1]
    bottomLeft = lines[y - 1][x - 1]
    bottomRight = lines[y - 1][x + 1]
    if topLeft == bottomRight:
        return False
    if topRight == bottomLeft:
        return False
    return True

total = 0 
for y, line in enumerate(lines):
    for x, c in enumerate(line):
        if c == "A":
            if checkIfPattern(x, y):
                total += 1

print("Total: ", total)
