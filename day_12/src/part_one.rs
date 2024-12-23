use std::{collections::HashMap, fs};

#[derive(Debug)]
struct Plant {
    neighbouring_distinct: i32,
    character: char,
}

type Pos = (i32, i32);

fn load_file(filename: &str) -> Vec<Vec<char>> {
    let contents = fs::read_to_string(filename).unwrap();
    contents
        .split_terminator("\n")
        .map(|st: &str| st.chars().collect())
        .collect()
}



fn populate_plants(input: &[Vec<char>]) -> HashMap<Pos, Plant> {
    let mut out: HashMap<Pos, Plant> = HashMap::new();
    let width = input.first().unwrap().len();
    let height = input.len();
    for (y, row) in input.iter().enumerate() {
        for (x, ch) in row.iter().enumerate() {
            let mut n = 4;
            if y + 1 < height && input.get(y + 1).unwrap().get(x).unwrap() == ch {
                n -= 1
            }
            if y > 0 && input.get(y - 1).unwrap().get(x).unwrap() == ch {
                n -= 1
            }

            if x > 0 && row.get(x - 1).unwrap() == ch {
                n -= 1
            }

            if x + 1 < width && row.get(x + 1).unwrap() == ch {
                n -= 1
            }

            out.insert((x as i32, y as i32), Plant {
                neighbouring_distinct: n,
                character: *ch,
            });
        }
    }

    out
}

fn explore_section(size: (i32, i32), searching_plant_type: char, plants: &HashMap<Pos, Plant>, visited: &mut HashMap<Pos, bool>, position: (i32, i32)){
    if *visited.get(&position).unwrap_or(&false){
        return;
    }
    if position.0 < 0 || position.1 < 0 || position.0 >= size.0 || position.1 >= size.1 {
        return;
    } 

    if searching_plant_type != plants.get(&position).unwrap().character{
        return;
    }
    visited.insert(position, true);
    explore_section(size, searching_plant_type, plants, visited, (position.0 + 1, position.1));
    explore_section(size, searching_plant_type, plants, visited, (position.0, position.1 + 1));
    explore_section(size, searching_plant_type, plants, visited, (position.0 - 1, position.1));
    explore_section(size, searching_plant_type, plants, visited, (position.0, position.1 - 1));
}

fn get_total(plants: &HashMap<Pos, Plant>, size: (i32, i32)) -> i32{
    let mut global_visited: HashMap<Pos, bool> = HashMap::new();
    let mut total = 0;
    for x in 0..size.0{
        for y in 0..size.1{
            if *global_visited.get(&(x, y)).unwrap_or(&false){
                continue;
            }
            let mut visited = HashMap::new();
            let ch = plants.get(&(x,y)).expect("Plant does not exist!").character;
            explore_section(size, ch, plants, &mut visited, (x, y));
            let area = visited.len() as i32;
            global_visited.extend(&visited);
            let perimiter: i32 = visited.clone().into_keys().map(|pos| plants.get(&pos).unwrap().neighbouring_distinct).sum();
            let cost = area * perimiter;
            visited.clear();
            total += cost;
        }
    }
    total
}


pub fn main() {
    let input = load_file("input.txt");
    let plants = populate_plants(&input);
    let width = input.first().unwrap().len() as i32;
    let height = input.len() as i32;
    let size = (width, height);
    let total = get_total(&plants, size);
    println!("Total: {}", total);
}
