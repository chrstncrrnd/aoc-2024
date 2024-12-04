file = open("input.txt", "r")
lines = file.read().splitlines()
file.close()

order = ["X", "M", "A", "S"]

width = len(lines[0])
height = len(lines)

def inBox(x, y):
    return x >= 0 and y >= 0 and x < width and y < height

def explore(x, y, dx, dy):
    currentChar = lines[y][x]
    if currentChar == "S":
        return True
    if not inBox(x + dx, y + dy):
        return False
    nextChar = lines[y + dy][x + dx] 
    cPos = order.index(currentChar)
    nPos = order.index(nextChar)
    if cPos + 1 != nPos:
        return False
    return explore(x + dx, y + dy, dx, dy)


total = 0 
for y, line in enumerate(lines):
    for x, c in enumerate(line):
        if c == "X":
            for i in range(-1, 2):
                for j in range(-1, 2):
                    newY = y + j
                    newX = x + i
                    if inBox(newX, newY) and lines[newY][newX] == "M" and explore(newX, newY, i, j):
                        total += 1
print("Total matches: ", total)
