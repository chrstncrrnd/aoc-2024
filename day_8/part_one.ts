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
        calcPosition(coordinate, {x: x, y: y}, dy, dx)
      }
    }
  }

}

const c2s = (c: Coord): string => c.x.toString() + "," + c.y.toString()

const calcPosition = (coord1: Coord, coord2: Coord, dy: number, dx: number) => {
  const p1 = {x: coord1.x + dx, y: coord1.y + dy}
  const p2 = {x: coord2.x - dx, y: coord2.y - dy}
  if (inArea(p1)){
    positions.set(c2s(p1), true)
  }
  if (inArea(p2)){
    positions.set(c2s(p2), true)
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
