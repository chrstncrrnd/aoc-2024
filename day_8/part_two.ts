const text = (await Deno.readTextFile("./input.txt")).split("\n")
text.pop()

type Coord = {
  x: number,
  y: number
}


const positions = new Map<string, boolean>
const height = text.length
const width = text[0].length

const inArea = (coordinate: Coord): boolean => coordinate.x >= 0 && coordinate.y >= 0 && coordinate.x < width && coordinate.y < height

const initSearch = (coordinate: Coord, char: string) => {
  let dy = 0
  let dx = 0
  for (let y = coordinate.y; y < height; y++){
    for (let x = 0; x < width; x++){
      if (y === coordinate.y && x <= coordinate.x) {
        continue
      }
      if (text[y][x] === char){
        dy = coordinate.y - y
        dx = coordinate.x - x
        calcPositions(coordinate, dy, dx)
      }
    }
  }

}

const c2s = (c: Coord): string => c.x.toString() + "," + c.y.toString()

const calcPositions = (coord: Coord, dy: number, dx: number) => {
  // forwardMode
  const forwardCoord = {...coord};
  while (inArea(forwardCoord)){
    positions.set(c2s(forwardCoord), true)
    forwardCoord.x += dx
    forwardCoord.y += dy
  }
  const backwardCoord = {...coord};
  while (inArea(backwardCoord)){
    positions.set(c2s(backwardCoord), true)
    backwardCoord.x -= dx
    backwardCoord.y -= dy
  }
}


for (const [y, line] of text.entries()) {
  for (const [x, char] of line.split("").entries()) {
    if (char !== "."){
      initSearch({x: x, y: y}, char)
    }
  }
}

console.log("Total: ", positions.size)
